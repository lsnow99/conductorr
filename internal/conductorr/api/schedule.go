package api

import (
	"net/http"
	"time"

	"github.com/lsnow99/conductorr/internal/conductorr/app"
	"github.com/lsnow99/conductorr/internal/conductorr/dbstore"
)

type Event struct {
	Timestamp   time.Time `json:"timestamp,omitempty"`
	MediaID     int       `json:"media_id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	ContentType string    `json:"content_type,omitempty"`
	SeasonNum   int       `json:"season_num,omitempty"`
	EpisodeNum  int       `json:"episode_num,omitempty"`
	SeriesID    int       `json:"series_id,omitempty"`
	SeriesTitle string    `json:"series_title,omitempty"`
	Monitoring  bool      `json:"monitoring,omitempty"`
	PathOK      bool      `json:"path_ok,omitempty"`
}

func GetSchedule(w http.ResponseWriter, r *http.Request) {
	medias, err := dbstore.GetAllMediaMap()
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	pathStatuses := app.LM.GetAllMediaPathStatuses()

	events := make([]Event, 0, len(medias))
	for _, media := range medias {
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
			event.EpisodeNum = int(media.Number.Int32)
			season, ok := medias[int(media.ParentMediaID.Int32)]
			if !ok {
				continue
			}
			event.SeasonNum = int(season.Number.Int32)
			series, ok := medias[int(season.ParentMediaID.Int32)]
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
	Respond(w, r.Header.Get("hostname"), nil, events, true)
}
