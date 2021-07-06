package integration

import (
	"testing"

	"github.com/hekmon/transmissionrpc"
)

func TestTransmissionConnection(t *testing.T) {
	transmission, err := NewTransmission("transmission", "transmission", "http://localhost:9091")
	if err != nil {
		t.Fatal(err)
	}
	err = transmission.Init()
	if err != nil {
		t.Fatal(err)
	}
	if err := transmission.TestConnection(); err != nil {
		t.Fatal(err)
	}
}

func TestTransmissionGetDownload(t *testing.T) {
	hash := "bfecf6cc9720da45b18e28f5a08c169bbd279840"

	client, err := transmissionrpc.New("localhost", "transmission", "transmission", &transmissionrpc.AdvancedConfig{
		HTTPS: false,
		Port:  uint16(9091),
	})

	if err != nil {
		t.Fatal(err)
	}

	torrents, err := client.TorrentGetAllForHashes([]string{hash})
	if err != nil {
		t.Fatal(err)
	}

	for _, torrent := range torrents {
		t.Logf("%v", *torrent.DownloadDir)
	}
}
