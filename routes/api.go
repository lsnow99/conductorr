package routes

import (
	"encoding/json"
	"errors"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/lsnow99/conductorr/settings"
)

var whitelistPaths = []string{
	"/api/v1/signin",
	"/api/v1/firstTime",
	"/api/v1/checkAuth",
}

var authErr = errors.New("token failed validation")

const UserAuthKey = "id_token"
const MaxCookieAgeSecs = 60 * 60 * 24 * 14

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
	r.FailedAuth = err == authErr

	if authorize {
		tok, err := GenerateIDToken()
		if err != nil {
			r.Success = false
		} else {
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
				Expires:  time.Now().Add(time.Second * MaxCookieAgeSecs),
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

		if shouldAuth {
			tok, err := r.Cookie(UserAuthKey)
			if err != nil || !checkToken(tok.Value) {
				Respond(w, "", authErr, nil, false)
				return
			}
		}

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func checkToken(tok string) bool {
	_, err := jwt.Parse(tok, func(t *jwt.Token) (interface{}, error) {
		return []byte(settings.JWTSecret), nil
	})
	return err == nil
}
