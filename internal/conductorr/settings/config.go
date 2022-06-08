package settings

import (
	"github.com/rs/zerolog/log"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

// Injected at build time
var DefaultOMDBAPIKey string
var DefaultTmdbAPIKey string

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
var TmdbAPIKey string
var ServerHost string
var CookieDomain string
var BuildMode string
var Version string
var TempDir string
var AllowInsecureRequests bool

func Init() {
	DebugMode = checkBool("CONDUCTORR_DEBUG")

	if jwtExp, exists := os.LookupEnv("JWT_EXP_DAYS"); exists {
		jwtExpDaysStr := jwtExp
		jwtExpDaysInt, err := strconv.Atoi(jwtExpDaysStr)
		if err != nil {
			log.Fatal().
				Msgf("error parsing JWT_EXP_DAYS (%s) as integer", jwtExpDaysStr)
		}
		JWTExpDays = jwtExpDaysInt
	} else {
		JWTExpDays = 7
	}

	if checkBool("RESET_USER") {
		ResetUser = true
		log.Info().
			Msg("allowing user reset since RESET_USER environment variable is set")
	}

	if _, exists := os.LookupEnv("PG_USER"); exists {
		PGUser = os.Getenv("PG_USER")
		PGPass = os.Getenv("PG_PASS")
		PGName = os.Getenv("PG_NAME")
		PGHost = os.Getenv("PG_HOST")
		PGPort = os.Getenv("PG_PORT")
		_, PGSSL = os.LookupEnv("PG_SSL")
		UsePG = true
		log.Info().
			Msg("using postgres for database")
	} else {
		log.Info().
			Msg("defaulting to sqlite for database")
	}

	pragmas := url.Values{}
	pragmas.Add("_pragma", "foreign_keys = on")
	pragmas.Add("_pragma", "cache = shared")
	pragmas.Add("_pragma", "mode = rwc")

	if dbPath, exists := os.LookupEnv("DB_PATH"); exists {
		DBPath = "file:" + dbPath + "?" + pragmas.Encode()
	} else {
		DBPath = "file:./conductorr.db?" + pragmas.Encode()
	}

	if omdbKey, exists := os.LookupEnv("OMDB_API_KEY"); exists {
		OmdbApiKey = omdbKey
	}
	if tmdbKey, exists := os.LookupEnv("TMDB_API_KEY"); exists {
		TmdbAPIKey = tmdbKey
	}

	// Set cookie domain from SERVER_HOST. Can parse IP addresses or hostnames
	if host, exists := os.LookupEnv("SERVER_HOST"); exists {
		ServerHost = host
	} else {
		ServerHost = "localhost"
	}
	addr := net.ParseIP(ServerHost)
	if addr == nil && ServerHost != "localhost" {
		CookieDomain = "." + ServerHost
	} else {
		CookieDomain = ServerHost
	}

	if mPath, exists := os.LookupEnv("MIGRATIONS_PATH"); exists {
		MigrationsPath = mPath
	} else {
		MigrationsPath = "./migrations"
	}

	AllowInsecureRequests = checkBool("ALLOW_INSECURE_REQUESTS")

	TempDir = os.Getenv("TEMP_DIR")
}

func checkBool(envName string) bool {
	val, exists := os.LookupEnv(envName)
	return exists && strings.ToLower(val) != "false"
}
