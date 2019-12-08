package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lsnow2017/conductorr/internal/constants"
	"net/http"
	"github.com/lsnow2017/conductorr/internal/schema"
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the LUCPGH Server API!\n")
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	user := &schema.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	user.PasswordHash = hashedPassword

	err = db.Insert(user)
	if err != nil {
		panic(err)
	}

	ss := GenerateToken(user.UserId)

	response := &schema.APIResponse{}
	response.Status = 200
	response.Action = "AUTH"
	response.Response = ss

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func GenerateToken(id int64) string {
	mySigningKey := []byte(os.Getenv("JWT_SIGNING_KEY"))

	// Create the Claims
	expirationTime := time.Now().Add(48 * time.Hour)
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Issuer:    "lucpgh server",
		Subject:   strconv.FormatInt(id, 10),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		panic(err)
	}
	return ss
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	user := &schema.User{}
	creds := &schema.Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		panic(err)
	}

	err = db.Model(user).
		Where("email = ?", creds.Email).
		Limit(1).
		Select()
	if err != nil {
		panic(err)
	}

	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(creds.Password))
	if err != nil {
		fmt.Fprintf(w, "Incorrect password")
	}

	ss := GenerateToken(user.UserId)

	response := &schema.APIResponse{}
	response.Status = 200
	response.Action = "AUTH"
	response.Response = ss

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func JWTRefreshHandler(w http.ResponseWriter, r *http.Request) {
	tokenVals := r.Context().Value("user")
	id := GetUserIdFromToken(tokenVals)
	ss := GenerateToken(id)

	response := &schema.APIResponse{}
	response.Status = 200
	response.Action = "AUTH"
	response.Response = ss

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func GetUserIdFromToken(tokenVals interface{}) int64 {
	claims := tokenVals.(*jwt.Token).Claims.(jwt.MapClaims)
	str, ok := claims["sub"].(string)
	if !ok {
		panic("Error")
	}
	id, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return id
}

func ConfigurationHandler(w http.ResponseWriter, r *http.Request) {
	configuration := schema.Configuration{}
	configuration.ConfigurationId = 0

	err := db.Select(&configuration)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(configuration); err != nil {
		panic(err)
	}
}

func GetConfigHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars["service"] == "filebot" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(constants.FilebotConfig); err != nil {
			panic(err)
		}
	} else if vars["service"] == "sonarr" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(constants.SonarrConfig); err != nil {
			panic(err)
		}
	} else if vars["service"] == "radarr" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(constants.RadarrConfig); err != nil {
			panic(err)
		}
	} else if vars["service"] == "plex" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(constants.PlexConfig); err != nil {
			panic(err)
		}
	}
}