package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/lsnow99/conductorr"
	"github.com/lsnow99/conductorr/dbstore"
	"github.com/lsnow99/conductorr/routes"
	"github.com/lsnow99/conductorr/settings"
)

func main() {
	if err := dbstore.Init(); err != nil {
		log.Fatal(err)
	}
	defer dbstore.Close()

	log.Fatal(serveRoutes(8282))
}

func serveRoutes(port int) error {
	http.HandleFunc("/api/v1/signin", routes.SignIn)
	http.HandleFunc("/api/v1/first_time", routes.FirstTime)
	http.HandleFunc("/api/v1/releaseProfileCfg", routes.GetReleaseProfileCfg)

	if settings.ResetUser {
		log.Println("Warning: allowing user registration either because no user exists in the database currenly, or the RESET_USER environment variable has been set. After the signup route successfully registers a user, the route will be disabled until the server exits.")
		http.HandleFunc("/api/v1/signup", routes.SignUp)
	}

	if !settings.DebugMode {
		var staticFS = http.FS(conductorr.WebDist)
		fs := http.FileServer(staticFS)
		http.Handle("/", fs)
	} else {
		log.Println("Warning: starting in debug mode")
	}

	log.Printf("Listening on :%d\n", port)
	return http.ListenAndServe(":"+strconv.Itoa(port), http.DefaultServeMux)
}
