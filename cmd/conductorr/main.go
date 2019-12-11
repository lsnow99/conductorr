package main

import (
	"net/http"

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

	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Could not load .env file: %s", err.Error())
	}

	db = pg.Connect(&pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
		Addr: os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
	})
	defer db.Close()

	runMigrations()

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

	r.HandleFunc("/auth/firstRun", FirstRunHandler).Methods("GET")
	r.HandleFunc("/auth/signup", SignupHandler).Methods("POST")
	r.HandleFunc("/auth/login", LoginHandler).Methods("POST")
	r.HandleFunc("/_link", LinkHandler).Methods("POST")

	ar.HandleFunc("/api/refreshToken", JWTRefreshHandler).Methods("GET")
	ar.HandleFunc("/api/config/{service}", GetConfigHandler).Methods("GET")
	ar.HandleFunc("/api/config/{service}", SetConfigHandler).Methods("POST")

	an := negroni.New(negroni.HandlerFunc(mw.HandlerWithNext), negroni.Wrap(ar))
	r.PathPrefix("/api").Handler(an)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

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
	out, err := exec.Command(path+"/migrations", "migrate").CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to run migration, output:\n%s", out)
		panic(err)
	}
	log.Printf(string(out))
	log.Printf("Migrations completed")
}
