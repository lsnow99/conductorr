package routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lsnow99/conductorr/dbstore"
	"github.com/lsnow99/conductorr/integration"
)

func SearchReleasesManual(w http.ResponseWriter, r *http.Request) {
	mediaIDStr := mux.Vars(r)["id"]
	mediaID, err := strconv.Atoi(mediaIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbMedia, err := dbstore.GetMediaByID(mediaID)
	if err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}

	media := integration.Media{
		Title: dbMedia.Title.String,
	}
	if dbMedia.ContentType.String == "movie" {
		media.ContentType = integration.Movie
	} else if dbMedia.ContentType.String == "series" {
		media.ContentType = integration.TVShow
	}
	media.ImdbID = dbMedia.ImdbID.String
	media.Year = dbMedia.ReleasedAt.Time.Year()

	xnab := integration.NewXnab(0, "7c1yDyAz12ZJDpUunuXyAeFxeFv0jjeI", "https://api.nzbgeek.info", "NZBGeek")
	xnab.TestConnection()
	results, err := xnab.Search(media)
	if err != nil {
		log.Println(err)
	}
	log.Println(results)
	Respond(w, r.Host, err, results, true)
}
