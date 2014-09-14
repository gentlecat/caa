package caa

import "testing"

func TestGetServerURL(t *testing.T) {
	url := getServerURL()
	if url != "http://coverartarchive.org" {
		t.Error("Expected \"http://coverartarchive.org\", got ", url)
	}

	// Trying to set another scheme
	Scheme = "https"
	url = getServerURL()
	if url != "https://coverartarchive.org" {
		t.Error("Expected \"https://coverartarchive.org\", got ", url)
	}

	// Trying to change host
	Host = "fakearchive.org"
	url = getServerURL()
	if url != "https://fakearchive.org" {
		t.Error("Expected \"https://fakearchive.org\", got ", url)
	}
}
