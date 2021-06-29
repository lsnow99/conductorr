package routes

import (
	"fmt"
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

	if !dbMedia.ProfileID.Valid {
		Respond(w, r.Host, fmt.Errorf("media must have profile assigned in order to search"), nil, true)
		return
	}
	profile, err := dbstore.GetProfileByID(int(dbMedia.ProfileID.Int32))
	if err != nil {
		Respond(w, r.Host, fmt.Errorf("error loading profile for media"), nil, true)
		return
	}
	if !profile.Filter.Valid || !profile.Sorter.Valid {
		Respond(w, r.Host, fmt.Errorf("profile must have sorter and filter defined"), nil, true)
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

	indexers, err := getIndexers()
	if err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}

	results := make([]integration.Release, 0)
	for _, indexer := range indexers {
		indexer.TestConnection()
		indexerResults, err := indexer.Search(&media)
		if err != nil {
			Respond(w, r.Host, err, nil, true)
			return
		}
		results = append(results, indexerResults...)
	}

	included, excluded, err := integration.FilterReleases(results, profile.Filter.String)
	if err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}
	for index := range excluded {
		excluded[index].Warnings = append(excluded[index].Warnings, "Release did not pass filter")
	}

	releases := append(included, excluded...)
	err = integration.SortReleases(&releases, profile.Sorter.String)
	if err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}

	Respond(w, r.Host, nil, releases, true)
}


func getIndexers() ([]*integration.Xnab, error) {
	dbIndexers, err := dbstore.GetIndexers()
	if err != nil {
		return nil, err
	}
	indexers := make([]*integration.Xnab, len(dbIndexers))
	for i, indexer := range dbIndexers {
		indexers[i] = integration.NewXnab(0, indexer.ApiKey, indexer.BaseUrl, indexer.Name, indexer.DownloadType)
	}
	return indexers, nil
}