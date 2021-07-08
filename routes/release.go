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

	profileID, err := getProfileID(mediaID)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}

	if profileID < 1 {
		Respond(w, r.Header.Get("hostname"), errors.New("no profile assigned"), nil, true)
		return
	}
	profile, err := dbstore.GetProfileByID(profileID)
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
		media.ImdbID = dbMedia.ImdbID.String
	} else if dbMedia.ContentType.String == "episode" {
		dbSeason, err := dbstore.GetMediaByID(int(dbMedia.ParentMediaID.Int32))
		if err != nil {
			Respond(w, r.Header.Get("hostname"), fmt.Errorf("episode has no parent"), nil, true)
			return
		}
		dbSeries, err := dbstore.GetMediaByID(int(dbSeason.ParentMediaID.Int32))
		if err != nil {
			Respond(w, r.Header.Get("hostname"), fmt.Errorf("episode has no parent"), nil, true)
			return
		}
		media.Title = dbSeries.Title.String
		media.Episode = int(dbMedia.Number.Int32)
		media.Season = int(dbSeason.Number.Int32)
		media.ContentType = integration.TVShow
	}
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

func getProfileID(mediaID int) (int, error) {
	media, err := dbstore.GetMediaByID(mediaID)
	if err != nil {
		return 0, err
	}
	if media.ParentMediaID.Valid {
		return getProfileID(int(media.ParentMediaID.Int32))
	}
	return int(media.ProfileID.Int32), nil
}