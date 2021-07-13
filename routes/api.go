package routes

import (
	"encoding/json"
	"errors"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/lsnow99/conductorr/settings"
)

var whitelistPaths = []string{
	"/api/v1/signin",
	"/api/v1/firstTime",
	"/api/v1/checkAuth",
	"/api/csl.wasm",
	"/api/v1/logout",
	// Also whitelisted are any paths not beginning with /api/v1
}

var errAuth = errors.New("token failed validation")

const UserAuthKey = "id_token"

type response struct {
	Success    bool        `json:"success"`
	Data       interface{} `json:"data"`
	Msg        string      `json:"msg"`
	FailedAuth bool        `json:"failed_auth"`
}

func Respond(w http.ResponseWriter, reqHost string, err error, data interface{}, authorize bool) {
	r := response{}
	if err != nil {
		r.Success = false
		r.Msg = err.Error()
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

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(r); err != nil {
		log.Println(err)
	}
}

func GenerateIDToken() (string, error) {
	// Create token claims
	claims := make(jwt.MapClaims)
	claims["sub"] = true
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

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

		if shouldAuth && strings.HasPrefix(r.URL.Path, "/api/v1") {
			tok, err := r.Cookie(UserAuthKey)
			if err != nil || !checkToken(tok.Value) {
				Respond(w, "", errAuth, nil, false)
				return
			}
		}

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		if (next != nil ) {
			next.ServeHTTP(w, r)
		}
	})
}

func InvalidateAuthCookie(w http.ResponseWriter, r *http.Request) {
	reqHost := r.Header.Get("hostname")
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

	Respond(w, r.Header.Get("hostname"), nil, nil, false)
}

func checkToken(tok string) bool {
	_, err := jwt.Parse(tok, func(t *jwt.Token) (interface{}, error) {
		return []byte(settings.JWTSecret), nil
	})
	return err == nil
}
