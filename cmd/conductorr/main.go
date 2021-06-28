package main

import (
	"log"
	"net/http"
	"strconv"

	_ "github.com/lsnow99/conductorr/internal/csl"
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

	http.Handle("/", routes.GetRouter())

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
