package routes

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lsnow99/conductorr/app"
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
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}

	if !dbMedia.ProfileID.Valid {
		Respond(w, r.Header.Get("hostname"), errors.New("no profile assigned"), nil, true)
		return
	}
	profile, err := dbstore.GetProfileByID(int(dbMedia.ProfileID.Int32))
	if err != nil {
		Respond(w, r.Header.Get("hostname"), fmt.Errorf("error loading profile for media"), nil, true)
		return
	}
	if !profile.Filter.Valid || !profile.Sorter.Valid {
		Respond(w, r.Header.Get("hostname"), fmt.Errorf("profile must have sorter and filter defined"), nil, true)
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

	results, err := app.IM.Search(&media)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}

	included, excluded, err := integration.FilterReleases(results, profile.Filter.String)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	for index := range excluded {
		excluded[index].Warnings = append(excluded[index].Warnings, "Release did not pass filter")
	}

	releases := append(included, excluded...)
	err = integration.SortReleases(&releases, profile.Sorter.String)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}

	Respond(w, r.Header.Get("hostname"), nil, releases, true)
}
