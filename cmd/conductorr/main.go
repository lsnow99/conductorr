package main

import (
	"flag"
	"net/http"
	"os"
	"os/signal"
	"runtime/pprof"
	"strconv"
	"syscall"

	"github.com/lsnow99/conductorr/internal/conductorr/api"
	"github.com/lsnow99/conductorr/internal/conductorr/app"
	"github.com/lsnow99/conductorr/internal/conductorr/dbstore"
	"github.com/lsnow99/conductorr/internal/conductorr/logger"
	"github.com/lsnow99/conductorr/internal/conductorr/scheduler"
	"github.com/lsnow99/conductorr/internal/conductorr/settings"
	"github.com/rs/zerolog/log"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	logger.Init()
	settings.Init()

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal().
				Stack().
				Err(err).
				Msg("")
		}
		pprof.StartCPUProfile(f)

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		go func() {
			for range c {
				pprof.StopCPUProfile()
				os.Exit(0)
			}
		}()
	}

	if err := dbstore.Init(); err != nil {
		log.Fatal().
			Stack().
			Err(err).
			Msg("")
	}
	defer dbstore.Close()

	// Initialize downloaders
	downloaders, err := dbstore.GetDownloaders()
	if err != nil {
		log.Fatal().
			Stack().
			Err(err).
			Msg("error getting downloaders from database")
	}
	for _, downloader := range downloaders {
		if err := app.DM.RegisterDownloader(downloader.ID, downloader.DownloaderType, downloader.Name, downloader.Config); err != nil {
			log.Error().
				Stack().
				Err(err).
				Msgf("error registering downloader id %d", downloader.ID)
		}
	}

	// Initialize the downloads
	downloads, err := dbstore.GetActiveDownloads()
	if err != nil {
		log.Fatal().
			Stack().
			Err(err).
			Msgf("error retrieving active downloads from database")
	}
	for _, download := range downloads {
		if download.MediaID.Valid {
			// TODO: verify that this shouldn't throw an error
			var releaseID *string
			if download.ReleaseID.Valid {
				releaseID = &download.ReleaseID.String
			}
			app.DM.RegisterDownload(download.ID, int(download.MediaID.Int32), download.FriendlyName, download.Status, download.Identifier, releaseID)
		} else {
			// This download was attached to a media that no longer exists
			log.Warn().
				Msgf("download %s registered with a downloader that no longer exists, recording it as an error")
			err = dbstore.UpdateDownloadStatusByIdentifier(download.Identifier, "error")
			if err != nil {
				log.Error().
					Stack().
					Err(err).
					Msgf("error updating download %s", download.Identifier)
			}
		}
	}

	// Initialize indexers
	indexers, err := dbstore.GetIndexers()
	if err != nil {
		log.Fatal().
			Stack().
			Err(err).
			Msg("error retrieving indexers from database")
	}
	for _, indexer := range indexers {
		app.IM.RegisterIndexer(indexer.ID, indexer.DownloadType, indexer.UserID, indexer.Name, indexer.ApiKey, indexer.BaseUrl, indexer.ForMovies, indexer.ForSeries, indexer.LastRSSID)
	}

	app.IM.DoTask()

	// Initialize media servers
	mediaServers, err := dbstore.GetMediaServers()
	if err != nil {
		log.Fatal().
			Stack().
			Err(err).
			Msg("error retrieving media servers from database")
	}
	for _, mediaServer := range mediaServers {
		if err := app.MSM.RegisterMediaServer(mediaServer.ID, mediaServer.MediaServerType, mediaServer.Name, mediaServer.Config); err != nil {
			log.Fatal().
				Stack().
				Err(err).
				Msg("error registering media server")
		}
	}

	scheduler.StartTasks()

	log.Fatal().
		Stack().
		Err(serveRoutes(8282)).Msg("")
}

func serveRoutes(port int) error {

	http.Handle("/", api.GetRouter())

	if settings.DebugMode {
		log.Warn().
			Msg("Warning: starting in debug mode")
	}

	log.Info().
		Msgf("Listening on :%d", port)
	return http.ListenAndServe(":"+strconv.Itoa(port), http.DefaultServeMux)
}
