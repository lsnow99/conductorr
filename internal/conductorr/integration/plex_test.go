package integration

import (
	"testing"

	"github.com/lsnow99/conductorr/internal/conductorr/services/plex"
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

func TestPlexGetPaths(t *testing.T) {
	token, err := plex.RetrievePlexToken("jschuler99@gmail.com", "%3PeAIO$d&l9")
	if err != nil {
		t.Fatal(err)
	}
	plex := NewPlex(token, "https://n0tflix.mediaserver.pro")
	paths, err := plex.GetMediaPaths()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(paths)
}
