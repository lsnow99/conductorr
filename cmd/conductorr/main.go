package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/signin", routes.SignIn).Methods("POST")
	r.HandleFunc("/api/v1/first_time", routes.FirstTime).Methods("GET")
	r.HandleFunc("/api/v1/releaseProfileCfg", routes.GetReleaseProfileCfg).Methods("GET")
	r.HandleFunc("/api/v1/profile", routes.CreateProfile).Methods("POST")
	r.HandleFunc("/api/v1/profile/{id}", routes.UpdateProfile).Methods("PUT")
	r.HandleFunc("/api/v1/profile/{id}", routes.DeleteProfile).Methods("DELETE")
	r.HandleFunc("/api/v1/profile", routes.GetProfiles).Methods("GET")
	r.Use(routes.AuthMiddleware)

	if settings.ResetUser {
		log.Println("Warning: allowing user registration either because no user exists in the database currenly, or the RESET_USER environment variable has been set. After the signup route successfully registers a user, the route will be disabled until the server exits.")
		http.HandleFunc("/api/v1/signup", routes.SignUp)
	}

	http.Handle("/", r)

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
