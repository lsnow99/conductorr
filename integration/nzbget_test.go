package integration

import (
	"testing"
)

func TestNZBGetConnection(t *testing.T) {
	nzbget := NewNZBGet("nzbget", "tegbzn6789", "http://localhost:6789")
	err := nzbget.Init()
	if err != nil {
		t.Fatal(err)
	}
	if err := nzbget.TestConnection(); err != nil {
		t.Fatal(err)
	}
}
