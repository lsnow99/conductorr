package settings

import (
	"log"
	"os"
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
var MigrationsPath string

func init() {

	if os.Getenv("CONDUCTORR_DEBUG") != "" {
		DebugMode = true
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
		DBPath = os.Getenv("DB_PATH")
	} else {
		DBPath = "./conductorr.db"
	}

	if os.Getenv("MIGRATIONS_PATH") != "" {
		MigrationsPath = os.Getenv("MIGRATIONS_PATH")
	} else if DebugMode {
		MigrationsPath = "./migrations"
	} else {
		// TODO: use embed fs
		MigrationsPath = "./migrations"
	}
}
