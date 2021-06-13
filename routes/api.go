package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/lsnow99/conductorr/settings"
	"github.com/dgrijalva/jwt-go"
)

type response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
	Token   string		`json:"token"`
}

func Respond(w http.ResponseWriter, err error, data interface{}, authorize bool) {
	r := response{}
	if err != nil {
		r.Success = false
		r.Msg = err.Error()
	} else {
		r.Success = true
	}
	r.Data = data

	if authorize {
		tok, err := GenerateIDToken()
		if err != nil {
			r.Success = false
		} else {
			r.Token = tok
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
