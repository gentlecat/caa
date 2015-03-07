// Package caa provides an easy way to access Cover Art Archive.
/*
Cover Art Archive (CAA) is a project with a goal to make cover art images
available to everyone on the Internet in an organized and convinient
format. More informantion is available at https://coverartarchive.org/.

There are two supported MusicBrainz ID (MBID) types for retrieval from CAA:
 - release
 - release group
*/
package caa

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type thumbnails struct {
	Large string `json:"large"`
	Small string `json:"small"`
}

type Image struct {
	ID         string     `json:"id"`
	Image      string     `json:"image"` // Full URL to the original image
	Types      []string   `json:"types"` // List of zero or more types for the image (https://musicbrainz.org/doc/Cover_Art/Types)
	IsFront    bool       `json:"front"`
	IsBack     bool       `json:"back"`
	Thumbnails thumbnails `json:"thumbnails"`
	Approved   bool       `json:"approved"` // Whether the image was approved by the MusicBrainz edit system
	Comment    string     `json:"comment"`  // A free text comment
	Edit       int        `json:"edit"`     // Edit ID on MusicBrainz
}

var (
	Scheme = "http" // http or https
	Host   = "coverartarchive.org"
)

func getServerURL() string { return Scheme + "://" + Host }

func getImages(url string) ([]Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("Failed to load images. Retuned status: " + resp.Status + ".")
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	type jason struct {
		Images []Image
	}

	respParsed := jason{}
	err = json.Unmarshal(body, &respParsed)
	if err != nil {
		return nil, err
	}

	return respParsed.Images, nil
}

// GetRelease returns list of cover art images for release with a specified MusicBrainz ID.
func GetRelease(mbid string) ([]Image, error) {
	return getImages(getServerURL() + "/release/" + mbid)
}

// GetReleaseGroup returns list of cover art images for release group with a specified MusicBrainz ID.
func GetReleaseGroup(mbid string) ([]Image, error) {
	return getImages(getServerURL() + "/release-group/" + mbid)
}
