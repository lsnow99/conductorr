package main

import (
	"github.com/go-pg/pg/v9"
	"github.com/joho/godotenv"

	// "github.com/go-pg/pg/v9/orm"
	"log"
	"os"
	"os/exec"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

var db *pg.DB

func main() {
	log.Printf("Starting Conductorr v1\n")

	runMigrations()

	err := godotenv.Load(".env")
	if err != nil {
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

func runMigrations() {
	log.Printf("Running any pending database migrations...")
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	out, err := exec.Command(path + "/migrations", "migrate").CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to run migration, output:\n%s", out)
		panic(err)
	}
	log.Printf(string(out))
	log.Printf("Migrations completed")
}