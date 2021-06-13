package integration

import (
	"testing"
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
