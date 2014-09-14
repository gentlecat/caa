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

	// Back to normal
	Scheme = "http"
	Host = "coverartarchive.org"
}

func TestGetRelease(t *testing.T) {
	_, err := GetRelease("9e2f9f67-cf77-45f9-8c82-0232425f5bb6")
	if err != nil {
		t.Error(err)
	}
}

func TestGetReleaseGroup(t *testing.T) {
	_, err := GetReleaseGroup("c2c696fc-6beb-3dfb-bb15-7bb021ebeb5d")
	if err != nil {
		t.Error(err)
	}
}
