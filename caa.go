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
	"io/ioutil"
	"net/http"
)

type Image struct {
	ID             string   `json:"id"`
	Image          string   `json:"image"`    // Full URL to the original image
	Approved       bool     `json:"approved"` // Whether the image was approved by the MusicBrainz edit system
	Comment        string   `json:"comment"`  // A free text comment
	Edit           int      `json:"edit"`     // Edit ID on MusicBrainz
	LargeThumbnail string   `json:"thumbnails.large"`
	SmallThumbnail string   `json:"thumbnails.small"`
	Types          []string `json:"types"`
	IsFront        bool     `json:"front"`
	IsBack         bool     `json:"back"`
}

var (
	Scheme = "http" // http or https
	Host   = "coverartarchive.org"
)

func getServerURL() string { return Scheme + "://" + Host }

// GetRelease returns list of cover art images for release group with a specified MusicBrainz ID.
func GetRelease(mbid string) ([]Image, error) {
	resp, err := http.Get(getServerURL() + "/release/" + mbid)
	if err != nil {
		return nil, err
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
