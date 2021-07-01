package integration

import (
	"testing"

	"github.com/hekmon/transmissionrpc"
)

func TestTransmissionConnection(t *testing.T) {
	transmission := NewTransmission("username", "password", "http://localhost:9091")
	err := transmission.Init()
	if err != nil {
		t.Fatal(err)
	}
	if err := transmission.TestConnection(); err != nil {
		t.Fatal(err)
	}
}

func TestTransmissionGetDownload(t *testing.T) {
	hash := "82d5abffb68a17f37ae0b0d51171a2e8dd45e6fe"

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

	t.Log(torrents)
}
