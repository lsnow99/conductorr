package main

import (
	"fmt"
	"github.com/go-pg/pg/v9"
	"github.com/joho/godotenv"
	// "github.com/go-pg/pg/v9/orm"
	"github.com/gorilla/mux"
	"log"
	"os"
	"github.com/urfave/negroni"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
)

var db *pg.DB

func main() {
	fmt.Printf("Running Conductorr v1\n")

	err := godotenv.Load(".env")
	if err != nil{
		log.Printf("Could not load .env file: %s", err.Error())
	}

	n := initRoutes()
	n.Run(":" + os.Getenv("PORT"))
}

func initRoutes() *negroni.Negroni {
	r := mux.NewRouter()
	ar := mux.NewRouter()

	mw := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SIGNING_KEY")), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/signup", SignupHandler).Methods("POST")
	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/backend/config/{service}", GetConfigHandler).Methods("GET")

	ar.HandleFunc("/api/refreshToken", JWTRefreshHandler).Methods("GET")
	ar.HandleFunc("/api/settings", ConfigurationHandler).Methods("GET")

	an := negroni.New(negroni.HandlerFunc(mw.HandlerWithNext), negroni.Wrap(ar))
	r.PathPrefix("/api").Handler(an)

	n := negroni.Classic()
	n.UseHandler(r)

	return n
}
