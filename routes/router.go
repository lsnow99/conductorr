package routes

import (
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/lsnow99/conductorr"
	"github.com/lsnow99/conductorr/settings"
)

func GetRouter() http.Handler {
	r := mux.NewRouter()

	// Unauthenticated routes (whitelist is in api.go)
	var cslFS = http.FS(conductorr.CSLDist)
	r.HandleFunc("/api/csl.wasm", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedEncodings := strings.Split(r.Header.Get("Accept-Encoding"), ",")
		filename := "dist/csl.wasm"
		for _, encoding := range allowedEncodings {
			if strings.TrimSpace(encoding) == "br" {
				w.Header().Add("Content-Encoding", "br")
				filename = "dist/csl.wasm.br"
				break
			}
		}
		w.Header().Add("Content-Type", "application/wasm")
		if settings.BuildMode == "binary" {
			file, err := cslFS.Open(filename)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusOK)
			io.Copy(w, file)
		} else {
			file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusOK)
			io.Copy(w, file)
		}
	})).Methods("GET")
	r.HandleFunc("/api/v1/signin", SignIn).Methods("POST")
	r.HandleFunc("/api/v1/firstTime", FirstTime).Methods("GET")
	r.HandleFunc("/api/v1/checkAuth", CheckAuth).Methods("GET")
	r.HandleFunc("/api/v1/logout", InvalidateAuthCookie).Methods("GET")

	// Authenticated routes
	r.HandleFunc("/api/v1/releaseProfileCfg", GetReleaseProfileCfg).Methods("GET")
	r.HandleFunc("/api/v1/profile", CreateProfile).Methods("POST")
	r.HandleFunc("/api/v1/profile/{id}", UpdateProfile).Methods("PUT")
	r.HandleFunc("/api/v1/profile/{id}", GetProfile).Methods("GET")
	r.HandleFunc("/api/v1/profile/{id}", DeleteProfile).Methods("DELETE")
	r.HandleFunc("/api/v1/profile", GetProfiles).Methods("GET")
	r.HandleFunc("/api/v1/library/search", SearchLibraryByTitle).Methods("GET")
	r.HandleFunc("/api/v1/new/search", SearchNewByTitle).Methods("GET")
	r.HandleFunc("/api/v1/add/{imdb_id}", AddMedia).Methods("POST")
	r.HandleFunc("/api/v1/poster/{id}", GetPoster).Methods("GET")
	r.HandleFunc("/api/v1/media/{id}", GetMedia).Methods("GET")
	r.HandleFunc("/api/v1/media/{id}", UpdateMedia).Methods("PUT")
	r.HandleFunc("/api/v1/media/{id}", DeleteMedia).Methods("DELETE")
	r.HandleFunc("/api/v1/media/{id}/searchReleasesManual", SearchReleasesManual).Methods("GET")
	r.HandleFunc("/api/v1/media/{id}/download", DownloadMediaRelease).Methods("POST")
	r.HandleFunc("/api/v1/media/{id}/monitoring", SetMonitoringMedia).Methods("PUT")
	r.HandleFunc("/api/v1/testIndexer", TestIndexer).Methods("POST")
	r.HandleFunc("/api/v1/indexer", CreateIndexer).Methods("POST")
	r.HandleFunc("/api/v1/indexer", GetIndexers).Methods("GET")
	r.HandleFunc("/api/v1/indexer/{id}", UpdateIndexer).Methods("PUT")
	r.HandleFunc("/api/v1/indexer/{id}", DeleteIndexer).Methods("DELETE")
	r.HandleFunc("/api/v1/testPath", TestPath).Methods("POST")
	r.HandleFunc("/api/v1/path", GetPaths).Methods("GET")
	r.HandleFunc("/api/v1/path", NewPath).Methods("POST")
	r.HandleFunc("/api/v1/path/{id}", UpdatePath).Methods("PUT")
	r.HandleFunc("/api/v1/path/{id}", DeletePath).Methods("DELETE")
	r.HandleFunc("/api/v1/testDownloader", TestDownloader).Methods("POST")
	r.HandleFunc("/api/v1/downloader", NewDownloader).Methods("POST")
	r.HandleFunc("/api/v1/downloader", GetDownloaders).Methods("GET")
	r.HandleFunc("/api/v1/downloader/{id}", UpdateDownloader).Methods("PUT")
	r.HandleFunc("/api/v1/downloader/{id}", DeleteDownloader).Methods("DELETE")
	r.HandleFunc("/api/v1/schedule", GetSchedule).Methods("GET")
	r.HandleFunc("/api/v1/status", GetStatus).Methods("GET")
	r.HandleFunc("/api/v1/logs", GetLogs).Methods("GET")
	r.HandleFunc("/api/v1/download", GetDownloads).Methods("GET")
	r.HandleFunc("/api/v1/plexAuthToken", FetchPlexAuthToken).Methods("POST")
	r.HandleFunc("/api/v1/testMediaServer", TestMediaServer).Methods("POST")
	r.HandleFunc("/api/v1/mediaServer", NewMediaServer).Methods("POST")
	r.HandleFunc("/api/v1/mediaServer", GetMediaServers).Methods("GET")
	r.HandleFunc("/api/v1/mediaServer/{id}", UpdateMediaServer).Methods("PUT")
	r.HandleFunc("/api/v1/mediaServer/{id}", DeleteMediaServer).Methods("DELETE")

	r.Use(AuthMiddleware)

	if settings.BuildMode == "binary" {
		sfs, err := fs.Sub(conductorr.WebDist, "web/build/dist")
		if err != nil {
			panic(err)
		}
		var staticFS = http.FS(sfs)
		fs := http.FileServer(staticFS)
		r.PathPrefix("/").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path

			// static files
			if strings.Contains(path, ".") || path == "/" {
				fs.ServeHTTP(w, r)
				return
			}

			// default to serve index.html
			f, _ := staticFS.Open("index.html")
			http.ServeContent(w, r, "index.html", time.Now().Add(time.Duration(24*365*10)*time.Hour), f)
		}))
	}

	if settings.ResetUser {
		log.Println("Warning: allowing user registration either because no user exists in the database currenly, or the RESET_USER environment variable has been set. After the signup route successfully registers a user, the route will be disabled until the server exits.")
		http.HandleFunc("/api/v1/signup", SignUp)
	}

	return r
}
