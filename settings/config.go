package settings

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

var DebugMode bool
var ResetUser bool
var UsePG bool
var PGSSL bool
var PGUser string
var PGPass string
var PGName string
var PGHost string
var PGPort string
var DBPath string
var JWTSecret string
var JWTExpDays int
var MigrationsPath string
var OmdbApiKey string
var ServerHost string
var CookieDomain string
var BuildMode string
var Version string

func init() {
	if os.Getenv("CONDUCTORR_DEBUG") != "" {
		DebugMode = true
	}

	if len(os.Getenv("JWT_SECRET")) >= 64 {
		JWTSecret = os.Getenv("JWT_SECRET")
	} else if DebugMode {
		// In debug mode just use a stupid JWT secret
		JWTSecret = "abcdefghijklmnopqrstuvwxyz0123456789"
	} else {
		log.Fatal(errors.New("required environment variable JWT_SECRET not provided, or was not at least 64 characters. Set JWT_SECRET to a random string of 64+ characters"))
	}

	if os.Getenv("JWT_EXP_DAYS") != "" {
		jwtExpDaysStr := os.Getenv("JWT_EXP_DAYS")
		jwtExpDaysInt, err := strconv.Atoi(jwtExpDaysStr)
		if err != nil {
			log.Fatalf("error parsing JWT_EXP_DAYS (%s) as integer", jwtExpDaysStr)
		}
		JWTExpDays = jwtExpDaysInt
	} else {
		JWTExpDays = 7
	}

	if os.Getenv("RESET_USER") != "" {
		ResetUser = true
		log.Println("Allowing user reset since RESET_USER environment variable is set")
	}

	if os.Getenv("PG_USER") != "" {
		PGUser = os.Getenv("PG_USER")
		PGPass = os.Getenv("PG_PASS")
		PGName = os.Getenv("PG_NAME")
		PGHost = os.Getenv("PG_HOST")
		PGPort = os.Getenv("PG_PORT")
		PGSSL = os.Getenv("PG_SSL") != ""
		UsePG = true
		log.Println("Using postgres for database")
	} else {
		log.Println("Defaulting to sqlite for database")
	}

	if os.Getenv("DB_PATH") != "" {
		DBPath = os.Getenv("DB_PATH") + "?_foreign_keys=on&cache=shared&mode=rwc"
	} else {
		DBPath = "./conductorr.db?_foreign_keys=on&cache=shared&mode=rwc"
	}

	if os.Getenv("OMDB_API_KEY") != "" {
		OmdbApiKey = os.Getenv("OMDB_API_KEY")
	} else if DebugMode {
		/*
			This is a free API key that has a limit of 1,000 requests per day. It is intended
			only for development purposes. Please do not abuse it. If you would like an API
			key for yourself, you can get one for free at https://www.omdbapi.com
		*/
		OmdbApiKey = "2fe27f52"
	}

	// Set cookie domain from SERVER_HOST. Can parse IP addresses or hostnames
	if os.Getenv("SERVER_HOST") != "" {
		ServerHost = os.Getenv("SERVER_HOST")
	} else {
		ServerHost = "localhost"
	}
	addr := net.ParseIP(ServerHost)
	if addr == nil && ServerHost != "localhost" {
		CookieDomain = "." + ServerHost
	} else {
		CookieDomain = ServerHost
	}

	if os.Getenv("MIGRATIONS_PATH") != "" {
		MigrationsPath = os.Getenv("MIGRATIONS_PATH")
	} else {
		MigrationsPath = "./migrations"
	}
}
