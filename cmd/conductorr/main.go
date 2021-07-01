package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/lsnow99/conductorr"
	"github.com/lsnow99/conductorr/app"
	"github.com/lsnow99/conductorr/dbstore"
	_ "github.com/lsnow99/conductorr/internal/csl"
	"github.com/lsnow99/conductorr/logger"
	"github.com/lsnow99/conductorr/routes"
	"github.com/lsnow99/conductorr/settings"
)

func main() {
	if err := dbstore.Init(); err != nil {
		log.Fatal(err)
	}
	defer dbstore.Close()

	// Initialize downloaders
	downloaders, err := dbstore.GetDownloaders()
	if err != nil {
		logger.LogToStdout(err)
		os.Exit(1)
	}
	for _, downloader := range downloaders {
		if err := app.DM.RegisterDownloader(downloader.DownloaderType, downloader.Name, downloader.Config); err != nil {
			logger.LogDanger(err)
		}
	}

	// Initialize indexers
	indexers, err := dbstore.GetIndexers()
	if err != nil {
		logger.LogToStdout(err)
		os.Exit(1)
	}
	for _, indexer := range indexers {
		var userID int
		if indexer.UserID.Valid {
			userID = int(indexer.UserID.Int32)
		}
		app.IM.RegisterIndexer(indexer.DownloadType, userID, indexer.Name, indexer.ApiKey, indexer.BaseUrl, indexer.ForMovies, indexer.ForSeries)
	}

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
