package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lsnow99/conductorr/internal/conductorr/app"
	"github.com/lsnow99/conductorr/internal/conductorr/dbstore"
	"github.com/lsnow99/conductorr/internal/conductorr/integration"
)

func SearchReleasesManual(w http.ResponseWriter, r *http.Request) {
	mediaIDStr := mux.Vars(r)["id"]
	mediaID, err := strconv.Atoi(mediaIDStr)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}

	allReleases, _, err := getReleasesForMediaID(mediaID)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}

	Respond(w, r.Header.Get("hostname"), nil, allReleases, true)
}

func SearchReleasesAuto(w http.ResponseWriter, r *http.Request) {
	mediaIDStr := mux.Vars(r)["id"]
	mediaID, err := strconv.Atoi(mediaIDStr)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}

	_, filteredReleases, err := getReleasesForMediaID(mediaID)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}

	app.DM.AutoDownload(mediaID, filteredReleases)

	Respond(w, r.Header.Get("hostname"), nil, len(filteredReleases), true)
}

func getReleasesForMediaID(mediaID int) ([]integration.Release, []integration.Release, error) {

	dbMedia, err := dbstore.GetMediaByID(mediaID)
	if err != nil {
		return nil, nil, err
	}

	profileID, err := getProfileID(mediaID)
	if err != nil {
		return nil, nil, err
	}

	if profileID < 1 {
		return nil, nil, fmt.Errorf("no profile assigned")
	}
	profile, err := dbstore.GetProfileByID(profileID)
	if err != nil {
		return nil, nil, fmt.Errorf("error loading profile for media")
	}
	if !profile.Filter.Valid || !profile.Sorter.Valid {
		return nil, nil, fmt.Errorf("profile must have sorter and filter defined")
	}

	var runtime *int64
	var results []integration.Release
	if dbMedia.ContentType.String == "movie" {
		var imdbID *string
		if dbMedia.ImdbID.Valid {
			imdbID = new(string)
			*imdbID = dbMedia.ImdbID.String
		}
		nzbs, err := app.IM.SearchMovie(dbMedia.Title.String, dbMedia.ReleasedAt.Time.Year(), imdbID)
		if err != nil {
			return nil, nil, err
		}
		runtime = new(int64)
		*runtime = int64(dbMedia.Runtime.Int32)
		results = append(results, nzbs...)
	} else if dbMedia.ContentType.String == "episode" {
		dbSeason, err := dbstore.GetMediaByID(int(dbMedia.ParentMediaID.Int32))
		if err != nil {
			return nil, nil, fmt.Errorf("episode has no parent")
		}
		dbSeries, err := dbstore.GetMediaByID(int(dbSeason.ParentMediaID.Int32))
		if err != nil {
			return nil, nil, fmt.Errorf("season has no parent")
		}
		var imdbID *string
		if dbSeries.ImdbID.Valid {
			imdbID = new(string)
			*imdbID = dbSeries.ImdbID.String
		}
		var tvdbID *int
		if dbSeries.TvdbID.Valid {
			tvdbID = new(int)
			*tvdbID = int(dbSeries.TvdbID.Int32)
		}
		nzbs, err := app.IM.SearchEpisode(int(dbSeason.Number.Int32), int(dbMedia.Number.Int32), dbSeries.Title.String, tvdbID, imdbID)
		if err != nil {
			return nil, nil, err
		}
		runtime = new(int64)
		*runtime = int64(dbSeries.Runtime.Int32)
		results = append(results, nzbs...)
	} else if dbMedia.ContentType.String == "season" {
		dbSeries, err := dbstore.GetMediaByID(int(dbMedia.ParentMediaID.Int32))
		if err != nil {
			return nil, nil, fmt.Errorf("season has no parent")
		}
		var imdbID *string
		if dbSeries.ImdbID.Valid {
			imdbID = new(string)
			*imdbID = dbSeries.ImdbID.String
		}
		var tvdbID *int
		if dbSeries.TvdbID.Valid {
			tvdbID = new(int)
			*tvdbID = int(dbSeries.TvdbID.Int32)
		}
		nzbs, err := app.IM.SearchSeason(int(dbMedia.Number.Int32), dbSeries.Title.String, tvdbID, imdbID)
		if err != nil {
			return nil, nil, err
		}
		runtime = new(int64)
		*runtime = int64(dbSeries.Runtime.Int32)
		results = append(results, nzbs...)

		// TODO: Search individual episodes as fallback
	} else if dbMedia.ContentType.String == "series" {
		// TODO: Search all seasons
	}

	included, excluded, err := integration.FilterReleases(results, profile.Filter.String, runtime)
	if err != nil {
		return nil, nil, err
	}

	history, err := dbstore.GetReleaseHistoryForMedia(mediaID)
	if err != nil {
		return nil, nil, err
	}

	i := 0
	for _, release := range included {
		found := false
		for _, releaseHistory := range history {
			if releaseHistory.ID == release.ID {
				found = true
			}
		}
		if found {
			release.Warnings = append(release.Warnings, "Release failed previously")
			excluded = append(excluded, release)
		} else {
			included[i] = release
			i++
		}
	}
	included = included[:i]

	allReleases := append(included, excluded...)
	err = integration.SortReleases(&allReleases, profile.Sorter.String, runtime)
	if err != nil {
		return nil, nil, err
	}

	filteredReleases := included
	err = integration.SortReleases(&filteredReleases, profile.Sorter.String, runtime)
	if err != nil {
		return nil, nil, err
	}

	return allReleases, filteredReleases, nil
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
