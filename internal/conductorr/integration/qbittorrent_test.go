package integration

import "testing"

func TestConnect(t *testing.T) {
	q, err := NewQBittorrent("admin", "adminadmin", "http://localhost:8080")
	if err != nil {
		t.Fatal(err)
	}
	if err := q.TestConnection(); err != nil {
		t.Fatal(err)
	}
}

func TestAddTorrent(t *testing.T) {
	q, err := NewQBittorrent("admin", "adminadmin", "http://localhost:8080")
	if err != nil {
		t.Fatal(err)
	}
	str, err := q.AddRelease(Release{
		DownloadURL: "magnet:?xt=urn:btih:CFAB036F9FD9F03D7A6A35C5E3118227E227D4CF&dn=Baggage%20Claim%20(2013)%20720p%20BrRip%20x264%20-%20YIFY&tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.openbittorrent.com%3A6969%2Fannounce&tr=udp%3A%2F%2F9.rarbg.to%3A2710%2Fannounce&tr=udp%3A%2F%2F9.rarbg.me%3A2780%2Fannounce&tr=udp%3A%2F%2F9.rarbg.to%3A2730%2Fannounce&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337&tr=http%3A%2F%2Fp4p.arenabg.com%3A1337%2Fannounce&tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce&tr=udp%3A%2F%2Ftracker.tiny-vps.com%3A6969%2Fannounce&tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce",
	})
	t.Logf("%s: %v", str, err)
}
