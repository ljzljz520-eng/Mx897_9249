package transfer

import "testing"

func TestConnectionCheck(t *testing.T) {
	if !CheckConnection("new").Reachable {
		t.Fatal("reachable")
	}
	if CheckConnection("offline").Reachable {
		t.Fatal("offline")
	}
}
