package routes

import (
	"net/http"
	"time"

	"github.com/lsnow99/conductorr/dbstore"
)

type Event struct {
	Timestamp time.Time `json:"timestamp,omitempty"`
	MediaID   int       `json:"media_id,omitempty"`
	Title     string    `json:"title,omitempty"`
}

func GetSchedule(w http.ResponseWriter, r *http.Request) {
	medias, err := dbstore.GetAllMedia()
	if err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}

	events := make([]Event, len(medias))
	for i, media := range medias {
		if !media.Title.Valid || !media.ReleasedAt.Valid {
			continue
		}
		events[i] = Event{
			Title: media.Title.String,
			MediaID: media.ID,
			Timestamp: media.ReleasedAt.Time,
		}
	}
	Respond(w, r.Host, nil, events, true)
}