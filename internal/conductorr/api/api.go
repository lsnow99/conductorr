// Package api for Conductorr.
//
// This package provides an API for communicating with the Conductorr backend.
//
//     Schemes: http, https
//     BasePath: /v1
//     Version: 1.0.0
//     License: GPL-3.0 https://www.gnu.org/licenses/gpl-3.0.html
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - api_key:
//	   - auth_token:
//
//     SecurityDefinitions:
//     api_key:
//          type: apiKey
//          name: x-token
//          in: header
//	   auth_token:
//			type:
//     oauth2:
//         type: oauth2
//         authorizationUrl: /oauth2/auth
//         tokenUrl: /oauth2/token
//         in: header
//         scopes:
//           bar: foo
//         flow: accessCode
//
//     Extensions:
//     x-meta-value: value
//     x-meta-array:
//       - value1
//       - value2
//     x-meta-array-obj:
//       - name: obj
//         value: field
//
// swagger:meta
package api

import (
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/lsnow99/conductorr/internal/conductorr/dbstore"
	"github.com/rs/zerolog/log"
	"github.com/lsnow99/conductorr/internal/conductorr/settings"
)

var whitelistPaths = []string{
	"/api/v1/signin",
	"/api/v1/firstTime",
	"/api/v1/checkAuth",
	"/api/csl.wasm",
	"/api/v1/logout",
	// Also whitelisted are any paths not beginning with /api
}

var errAuth = errors.New("token failed validation")

const UserAuthKey = "id_token"

type response struct {
	Success    bool        `json:"success"`
	Data       interface{} `json:"data"`
	Msg        string      `json:"msg"`
	FailedAuth bool        `json:"failed_auth"`
}

func Respond(w http.ResponseWriter, req *http.Request, err error, data interface{}, authorize bool) {
	r := response{}
	if err != nil {
		r.Success = false
		r.Msg = err.Error()
		log.Warn().
			Err(err).
			Msg("api error")
	} else {
		r.Success = true
	}
	r.Data = data
	r.FailedAuth = err == errAuth

	if authorize {
		tok, err := GenerateIDToken()
		if err != nil {
			r.Success = false
		} else {
			reqHost := req.Host
			if i := strings.Index(reqHost, ":"); i > 0 {
				reqHost = reqHost[:i]
			}
			var cookieDomain string
			addr := net.ParseIP(reqHost)
			if addr == nil && reqHost != "localhost" {
				cookieDomain = "." + reqHost
			} else {
				cookieDomain = reqHost
			}

			cookie := http.Cookie{
				Name:     UserAuthKey,
				Value:    tok,
				Path:     "/",
				Domain:   cookieDomain,
				Expires:  time.Now().Add(time.Hour * 24 * time.Duration(settings.JWTExpDays)),
				Secure:   !settings.DebugMode,
				HttpOnly: true,
			}

			if settings.DebugMode {
				cookie.SameSite = http.SameSiteLaxMode
			} else {
				cookie.SameSite = http.SameSiteStrictMode
			}

			http.SetCookie(w, &cookie)
		}
	}

	if err == errAuth {
		w.WriteHeader(http.StatusUnauthorized)
	}
	if err := json.NewEncoder(w).Encode(r); err != nil {
		log.Error().
			Err(err).
			Bool("internal", true).
			Msg("error encoding json api response")
	}
}

func GenerateIDToken() (string, error) {
	// Create token claims
	claims := make(jwt.MapClaims)
	claims["sub"] = true
	claims["exp"] = time.Now().Add(time.Hour * 24 * time.Duration(settings.JWTExpDays)).Unix()

	tok := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return tok.SignedString([]byte(settings.JWTSecret))
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var shouldAuth = true
		for _, path := range whitelistPaths {
			if path == r.URL.Path {
				shouldAuth = false
			}
		}

		if shouldAuth && strings.HasPrefix(r.URL.Path, "/api") {
			// Check cookie
			tok, err := r.Cookie(UserAuthKey)
			if (err == nil) && checkToken(tok.Value) {
				goto DoNextHandler
			}

			// Check user
			username, password, ok := r.BasicAuth()
			err = dbstore.CheckUser(r.Context(), username, password)
			if (err == nil) && ok {
				goto DoNextHandler
			}

			// TODO: Check auth token
			// authToken := r.Header.Get("X-Auth-Token")

			// Return an auth error if all three methods failed.
			Respond(w, r, errAuth, nil, false)
			return
		}

	DoNextHandler:

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		if next != nil {
			next.ServeHTTP(w, r)
		}
	})
}

func InvalidateAuthCookie(w http.ResponseWriter, r *http.Request) {
	reqHost := r.Host
	if i := strings.Index(reqHost, ":"); i > 0 {
		reqHost = reqHost[:i]
	}
	var cookieDomain string
	addr := net.ParseIP(reqHost)
	if addr == nil && reqHost != "localhost" {
		cookieDomain = "." + reqHost
	} else {
		cookieDomain = reqHost
	}

	cookie := http.Cookie{
		Name:     UserAuthKey,
		Value:    "loggedout",
		Path:     "/",
		Domain:   cookieDomain,
		Expires:  time.Now().Add(time.Hour * 24),
		Secure:   !settings.DebugMode,
		HttpOnly: true,
	}

	if settings.DebugMode {
		cookie.SameSite = http.SameSiteLaxMode
	} else {
		cookie.SameSite = http.SameSiteStrictMode
	}

	http.SetCookie(w, &cookie)

	Respond(w, r, nil, nil, false)
}

func checkToken(tok string) bool {
	_, err := jwt.Parse(tok, func(t *jwt.Token) (interface{}, error) {
		return []byte(settings.JWTSecret), nil
	})
	if err != nil {
		log.Warn().
			Err(err).
			Str("token", tok).
			Msg("invalid auth token")
	}
	return err == nil
}
