package api

import (
	"encoding/json"
	"net/http"
	"sort"
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

type Interval struct {
	MinVal float64 `json:"minVal,omitempty"`
	MaxVal float64 `json:"maxVal,omitempty"`
}

func GetSchedule(w http.ResponseWriter, r *http.Request) {
	intervalsStr := r.URL.Query().Get("intervals")

	intervals := make([]Interval, 0)

	err := json.Unmarshal([]byte(intervalsStr), &intervals)
	if err != nil {
		Respond(w, r, err, nil, true)
	}

	dateIntervals := make([]dbstore.DateInterval, 0, len(intervals))
	for _, interval := range intervals {
		dateIntervals = append(dateIntervals, dbstore.DateInterval{
			DateFrom: time.Unix(int64(interval.MinVal), 0),
			DateTo:   time.Unix(int64(interval.MaxVal), 0),
		})
	}

	medias, err := dbstore.GetMediaInIntervals(r.Context(), dateIntervals)
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
