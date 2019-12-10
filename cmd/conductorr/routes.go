package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/go-pg/pg/v9"
	"github.com/lsnow2017/conductorr/internal/constants"
	"github.com/lsnow2017/conductorr/internal/schema"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// HomeHandler api welcome message
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Conductorr Server API!\n")
}

// isFirstRun return true if the reset_account variable is set or db empty
func isFirstRun() bool {
	sysConfig := &schema.SystemConfiguration{}
	sysConfig.SystemConfigurationID = true

	// Check if there isn't already a user entry in the database
	err := db.Select(sysConfig)
	firstRun := (err == pg.ErrNoRows)
	if err != nil && err != pg.ErrNoRows {
		panic(err)
	}

	resetting := os.Getenv("RESET_ACCOUNT") == "1"

	return resetting || firstRun
}

// FirstRunHandler allow frontend to ask if it is the first run
func FirstRunHandler(w http.ResponseWriter, r *http.Request) {

	responseObj := map[string]bool{"first_run": isFirstRun()}
	responseJSON, err := json.Marshal(responseObj)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(responseJSON))
}

// SignupHandler write a new user to the database if none exists, or RESET_ACCOUNT = 1
func SignupHandler(w http.ResponseWriter, r *http.Request) {

	if !isFirstRun() {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Account already exists. Set RESET_ACCOUNT env to reset")
		return
	}

	sysConfig := &schema.SystemConfiguration{}
	creds := &schema.Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)
	sysConfig.Password = hashedPassword
	sysConfig.Username = creds.Username

	err = db.Insert(sysConfig)
	if err != nil {
		panic(err)
	}

	ss := GenerateToken()

	fmt.Fprint(w, ss)
}

// GenerateToken generate a new JWT token
func GenerateToken() string {
	mySigningKey := []byte(os.Getenv("JWT_SIGNING_KEY"))

	// Create the Claims
	expirationTime := time.Now().Add(48 * time.Hour)
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Issuer:    "Conductorr server",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		panic(err)
	}
	return ss
}

// LoginHandler handle a login request
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	sysConfig := &schema.SystemConfiguration{}
	creds := &schema.Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		panic(err)
	}

	sysConfig.Username = creds.Username
	err = db.Model(sysConfig).Where("username = ?", creds.Username).Select()
	if err != nil {
		panic(err)
	}

	err = bcrypt.CompareHashAndPassword(sysConfig.Password, []byte(creds.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Incorrect password")
		return
	}

	ss := GenerateToken()

	fmt.Fprint(w, ss)
	// if err := json.NewEncoder(w).Encode(response); err != nil {
	// 	panic(err)
	// }
}

// JWTRefreshHandler refresh JWT if authenticated
func JWTRefreshHandler(w http.ResponseWriter, r *http.Request) {
	ss := GenerateToken()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(ss))
	// if err := json.NewEncoder(w).Encode(response); err != nil {
	// 	panic(err)
	// }
}

func ConfigurationHandler(w http.ResponseWriter, r *http.Request) {
	configuration := schema.FilebotConfiguration{}
	configuration.FilebotConfigurationID = true

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

// PostConfigHandler update config in database
func PostConfigHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	switch vars["service"] {
	case "filebot":
		config := &schema.FilebotConfiguration{}
		config.FilebotConfigurationID = true

	case "sonarr":

	case "radarr":

	case "plex":

	default:

	}
}

// GetConfigHandler return config form data
func GetConfigHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var configObj interface{}

	switch vars["service"] {
	case "filebot":
		configObj = constants.FilebotConfig
	case "sonarr":
		configObj = constants.SonarrConfig
	case "radarr":
		configObj = constants.RadarrConfig
	case "plex":
		configObj = constants.PlexConfig
	default:
		w.Header().Set("Content-Type", "text")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %s", "Invalid service")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(configObj); err != nil {
		panic(err)
	}
}
