package main

import (
	"net/http"
	"strconv"

	"github.com/go-pg/pg/v9"
	"github.com/joho/godotenv"

	"log"
	"os"
	"os/exec"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

var db *pg.DB
var sonarr *Sonarr
var radarr *Radarr
var filebot *Filebot
var plex *Plex

func main() {
	log.Printf("Starting Conductorr v1\n")

	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Did not load .env file: %s. Using given environment variables", err.Error())
	}
	if len(os.Getenv("JWT_SIGNING_KEY")) < 32 {
		log.Fatal("Please set env var JWT_SIGNING_KEY to a random string at least 32 characters in length!")
		os.Exit(1)
	}

	db = pg.Connect(&pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
		Addr:     os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
	})
	defer db.Close()

	runMigrations()

	initConfigs()

	n := initRoutes()
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 80
	}
	n.Run(":" + strconv.Itoa(port))
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
	r.HandleFunc("/_import", ImportHandler).Methods("POST")

	ar.HandleFunc("/api/refreshToken", JWTRefreshHandler).Methods("GET")
	ar.HandleFunc("/api/config/{service}", GetConfigHandler).Methods("GET")
	ar.HandleFunc("/api/config/{service}", SetConfigHandler).Methods("POST")
	ar.HandleFunc("/api/testConfig/{service}", TestConfigHandler).Methods("POST")

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

func initConfigs() {
	filebot = new(Filebot)
	sonarr = new(Sonarr)
	plex = new(Plex)
	radarr = new(Radarr)
	sonarr.LoadConfiguration(true)
	radarr.LoadConfiguration(true)
	filebot.LoadConfiguration(true)
	plex.LoadConfiguration(true)
}
