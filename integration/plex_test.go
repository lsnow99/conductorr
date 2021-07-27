package integration

import (
	"testing"

	"github.com/lsnow99/conductorr/services/plex"
)

func TestPlexBasic(t *testing.T) {
	token, err := plex.RetrievePlexToken("email", "password")
	if err != nil {
		t.Fatal(err)
	}
	config := map[string]interface{}{
		"token":    token,
		"base_url": "https://n0tflix.mediaserver.pro",
	}
	plex, err := NewPlexFromConfig(config)
	if err != nil {
		t.Fatal(err)
	}
	if err := plex.TestConnection(); err != nil {
		t.Fatal(err)
	}
}
