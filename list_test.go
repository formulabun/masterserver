package masterserver

import (
	"testing"
)

func TestListServers(t *testing.T) {
	n, err := ListServers()
	if err != nil {
		t.Fatal("Could not get servers", err)
	}

	if len(n) == 0 {
		t.Fatalf("Expected some results, only got %v", len(n))
	}
}
