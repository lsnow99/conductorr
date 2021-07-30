package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/lsnow99/conductorr/app"
	"github.com/lsnow99/conductorr/dbstore"
	_ "github.com/lsnow99/conductorr/internal/csl"
	"github.com/lsnow99/conductorr/logger"
	"github.com/lsnow99/conductorr/routes"
	"github.com/lsnow99/conductorr/scheduler"
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
		if err := app.DM.RegisterDownloader(downloader.ID, downloader.DownloaderType, downloader.Name, downloader.Config); err != nil {
			logger.LogDanger(err)
		}
	}

	// Initialize the downloads
	downloads, err := dbstore.GetActiveDownloads()
	if err != nil {
		logger.LogToStdout(err)
		os.Exit(1)
	}
	for _, download := range downloads {
		if err != nil {
			logger.LogToStdout(err)
			os.Exit(1)
		}
		if download.MediaID.Valid {
			app.DM.RegisterDownload(download.ID, int(download.MediaID.Int32), download.FriendlyName, download.Status, download.Identifier)
		} else {
			// This download was registered with a downloader that no longer exists. Record it as an error
			err = dbstore.UpdateDownloadStatusByIdentifier(download.Identifier, "error")
			if err != nil {
				logger.LogToStdout(err)
			}
		}
	}

	// Initialize indexers
	indexers, err := dbstore.GetIndexers()
	if err != nil {
		logger.LogToStdout(err)
		os.Exit(1)
	}
	for _, indexer := range indexers {
		app.IM.RegisterIndexer(indexer.ID, indexer.DownloadType, indexer.UserID, indexer.Name, indexer.ApiKey, indexer.BaseUrl, indexer.ForMovies, indexer.ForSeries)
	}

	// Initialize media servers
	mediaServers, err := dbstore.GetMediaServers()
	if err != nil {
		logger.LogToStdout(err)
		os.Exit(1)
	}
	for _, mediaServer := range mediaServers {
		if err := app.MSM.RegisterMediaServer(mediaServer.ID, mediaServer.MediaServerType, mediaServer.Name, mediaServer.Config); err != nil {
			logger.LogToStdout(err)
			os.Exit(1)
		}
	}

	scheduler.StartTasks()

	log.Fatal(serveRoutes(8282))
}

func serveRoutes(port int) error {

	http.Handle("/", routes.GetRouter())

	if settings.DebugMode {
		log.Println("Warning: starting in debug mode")
	}

	log.Printf("Listening on :%d\n", port)
	return http.ListenAndServe(":"+strconv.Itoa(port), http.DefaultServeMux)
}
