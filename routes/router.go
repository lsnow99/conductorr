package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lsnow99/conductorr/settings"
)

func GetRouter() http.Handler {
	r := mux.NewRouter()
	
	// Unauthenticated routes
	r.HandleFunc("/api/csl.wasm", LoadCSL).Methods("GET")
	r.HandleFunc("/api/v1/signin", SignIn).Methods("POST")
	r.HandleFunc("/api/v1/firstTime", FirstTime).Methods("GET")
	r.HandleFunc("/api/v1/checkAuth", CheckAuth).Methods("GET")

	r.HandleFunc("/api/v1/releaseProfileCfg", GetReleaseProfileCfg).Methods("GET")
	r.HandleFunc("/api/v1/profile", CreateProfile).Methods("POST")
	r.HandleFunc("/api/v1/profile/{id}", UpdateProfile).Methods("PUT")
	r.HandleFunc("/api/v1/profile/{id}", DeleteProfile).Methods("DELETE")
	r.HandleFunc("/api/v1/profile/{id}", GetProfile).Methods("GET")
	r.HandleFunc("/api/v1/profile", GetProfiles).Methods("GET")
	r.HandleFunc("/api/v1/library/search", SearchLibraryByTitle).Methods("GET")
	r.HandleFunc("/api/v1/new/search", SearchNewByTitle).Methods("GET")
	r.HandleFunc("/api/v1/add/{imdb_id}", AddMedia).Methods("POST")
	r.HandleFunc("/api/v1/poster/{id}", GetPoster).Methods("GET")
	r.HandleFunc("/api/v1/media/{id}", GetMedia).Methods("GET")
	r.HandleFunc("/api/v1/media/{id}/searchReleasesManual", SearchReleasesManual).Methods("POST")
	r.Use(AuthMiddleware)

	if settings.ResetUser {
		log.Println("Warning: allowing user registration either because no user exists in the database currenly, or the RESET_USER environment variable has been set. After the signup route successfully registers a user, the route will be disabled until the server exits.")
		http.HandleFunc("/api/v1/signup", SignUp)
	}

	return r
}
