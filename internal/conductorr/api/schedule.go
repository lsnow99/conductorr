package api

import (
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/lsnow99/conductorr/internal/conductorr/app"
	"github.com/lsnow99/conductorr/internal/conductorr/dbstore"
)

type Event struct {
	Timestamp   time.Time `json:"timestamp,omitempty"`
	MediaID     int       `json:"mediaID,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	ContentType string    `json:"contentType,omitempty"`
	SeasonNum   int       `json:"seasonNum,omitempty"`
	EpisodeNum  int       `json:"episodeNum,omitempty"`
	SeriesID    int       `json:"seriesID,omitempty"`
	SeriesTitle string    `json:"seriesTitle,omitempty"`
	Monitoring  bool      `json:"monitoring,omitempty"`
	PathOK      bool      `json:"pathOk,omitempty"`
}

func GetSchedule(w http.ResponseWriter, r *http.Request) {
	dateFromUnixStr := r.URL.Query().Get("dateFrom")
	dateToUnixStr := r.URL.Query().Get("dateTo")

	dateFromUnix, err := strconv.ParseInt(dateFromUnixStr, 10, 64)
	if err != nil {
		Respond(w, r, err, nil, true)
	}
	dateToUnix, err := strconv.ParseInt(dateToUnixStr, 10, 64)
	if err != nil {
		Respond(w, r, err, nil, true)
	}

	dateFrom := time.Unix(dateFromUnix, 0)
	dateTo := time.Unix(dateToUnix, 0)

	medias, err := dbstore.GetMediaInRange(r.Context(), dateFrom, dateTo)
	if err != nil {
		Respond(w, r, err, nil, true)
		return
	}

	mediaMap := make(map[int]dbstore.Media)
	for _, media := range medias {
		mediaMap[media.ID] = media
	}

	pathStatuses := app.LM.GetAllMediaPathStatuses()

	events := make([]Event, 0, len(medias))
	for _, media := range mediaMap {
		if !media.Title.Valid || !media.ReleasedAt.Valid {
			continue
		}
		pathOK, inMap := pathStatuses[media.ID]
		event := Event{
			Title:       media.Title.String,
			MediaID:     media.ID,
			Timestamp:   media.ReleasedAt.Time,
			Description: media.Description.String,
			ContentType: media.ContentType.String,
			Monitoring:  media.Monitoring,
			PathOK:      pathOK && inMap,
		}
		if media.ContentType.String == "episode" {
			event.EpisodeNum = int(media.ItemNumber.Int32)
			season, ok := mediaMap[int(media.ParentMediaID.Int32)]
			if !ok {
				continue
			}
			event.SeasonNum = int(season.ItemNumber.Int32)
			series, ok := mediaMap[int(season.ParentMediaID.Int32)]
			if !ok {
				continue
			}
			event.SeriesID = series.ID
			event.SeriesTitle = series.Title.String
		} else if media.ContentType.String == "movie" {
		} else {
			// Ignore series or season content types
			continue
		}
		events = append(events, event)
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i].Timestamp.Before(events[j].Timestamp)
	})
	Respond(w, r, nil, events, true)
}
