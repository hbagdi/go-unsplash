// Copyright (c) 2017 Hardik Bagdi <hbagdi1@binghamton.edu>
//
// MIT License
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package unsplash

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

// PhotosService interacts with /photos endpoint
type PhotosService service

// PhotoOpt denotes properties of any Image
type PhotoOpt struct {
	Height int `json:"h" url:"h"`
	Width  int `json:"w" url:"w"`
	CropX  int
	CropY  int
	Crop   bool
}

// Valid validates a PhotoOpt
func (p *PhotoOpt) Valid() bool {
	if p.Height <= 0 || p.Width <= 0 || p.CropX < 0 || p.CropY < 0 {
		return false
	}
	return true
}

type rect struct {
	Rect string `url:"rect"`
}

func processPhotoOpt(photoOpt *PhotoOpt) interface{} {
	if !photoOpt.Crop {
		return photoOpt
	}
	var buf bytes.Buffer
	var r rect
	buf.WriteString(strconv.Itoa(photoOpt.CropX) + "," +
		strconv.Itoa(photoOpt.CropY) + "," +
		strconv.Itoa(photoOpt.Width) + "," +
		strconv.Itoa(photoOpt.Height))
	r.Rect = buf.String()
	return r
}

// Photo return a photo with id
func (ps *PhotosService) Photo(id string, photoOpt *PhotoOpt) (*Photo, *Response, error) {
	if "" == id {
		return nil, nil, &IllegalArgumentError{ErrString: "Photo ID cannot be null"}
	}

	// Validation and conversion if necessary of photoOpt
	var opt interface{}
	opt = nil
	if photoOpt != nil {
		if !photoOpt.Valid() {
			return nil, nil, &InvalidPhotoOptError{ErrString: " photoOpt has zero or non-negative values"}
		}
		opt = processPhotoOpt(photoOpt)
	}
	endpoint := fmt.Sprintf("%v/%v", getEndpoint(photos), id)
	req, err := newRequest(GET, endpoint, opt, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err := ps.client.do(req)
	if err != nil {
		return nil, nil, err
	}
	var photo Photo
	err = json.Unmarshal(*resp.body, &photo)
	if err != nil {
		return nil, nil, err
	}
	return &photo, resp, nil
}

// Stats return a stats about a photo with id.
func (ps *PhotosService) Stats(id string) (*PhotoStats, *Response, error) {
	if "" == id {
		return nil, nil, &IllegalArgumentError{ErrString: "Photo ID cannot be null"}
	}
	endpoint := fmt.Sprintf("%v/%v/stats", getEndpoint(photos), id)
	req, err := newRequest(GET, endpoint, nil, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err := ps.client.do(req)
	if err != nil {
		return nil, nil, err
	}
	var stats PhotoStats
	err = json.Unmarshal(*resp.body, &stats)
	if err != nil {
		return nil, nil, err
	}
	return &stats, resp, nil
}

// Statistics return a stats about a photo with id.
func (ps *PhotosService) Statistics(id string, opt *StatsOpt) (*PhotoStatistics, *Response, error) {
	if "" == id {
		return nil, nil, &IllegalArgumentError{ErrString: "Photo ID cannot be null"}
	}
	if opt == nil {
		opt = defaultStatsOpt
	}
	if !opt.Valid() {
		return nil, nil, &InvalidStatsOptError{ErrString: "opt provided is not valid."}
	}
	endpoint := fmt.Sprintf("%v/%v/statistics", getEndpoint(photos), id)
	req, err := newRequest(GET, endpoint, opt, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err := ps.client.do(req)
	if err != nil {
		return nil, nil, err
	}
	var stats PhotoStatistics
	err = json.Unmarshal(*resp.body, &stats)
	if err != nil {
		return nil, nil, err
	}
	return &stats, resp, nil
}

// DownloadLink return the download URL for a photo.
func (ps *PhotosService) DownloadLink(id string) (*URL, *Response, error) {
	if "" == id {
		return nil, nil, &IllegalArgumentError{ErrString: "Photo ID cannot be null"}
	}
	endpoint := fmt.Sprintf("%v/%v/download", getEndpoint(photos), id)
	req, err := newRequest(GET, endpoint, nil, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err := ps.client.do(req)
	if err != nil {
		return nil, nil, err
	}
	var url urlWrapper
	err = json.Unmarshal(*resp.body, &url)
	if err != nil {
		return nil, nil, err
	}
	return url.URL, resp, nil
}

// All returns a list of all photos on unsplash.
// Note that some fields in photo structs from this result will be missing.
// Use Photo() method to get all details of the  Photo.
func (ps *PhotosService) All(listOpt *ListOpt) (*[]Photo, *Response, error) {
	s := (service)(*ps)
	return s.getPhotos(listOpt, getEndpoint(photos))
}

// Curated return a list of all curated photos.
func (ps *PhotosService) Curated(listOpt *ListOpt) (*[]Photo, *Response, error) {
	s := (service)(*ps)
	return s.getPhotos(listOpt, getEndpoint(photos)+"/curated")
}

// RandomPhotoOpt optional parameters for a random photo search
type RandomPhotoOpt struct {
	Height        int         `url:"h,omitempty"`
	Width         int         `url:"w,omitempty"`
	Featured      bool        `url:"featured,omitempty"`
	Username      string      `url:"username,omitempty"`
	SearchQuery   string      `url:"query,omitempty"`
	Count         int         `url:"count,omitempty"`
	Orientation   orientation `url:"orientation,omitempty"`
	CollectionIDs []int       `url:"collections,comma"`
	TopicIDs      []string    `url:"topics,comma"`
}

// Valid validates a RandomPhotoOpt
func (opt *RandomPhotoOpt) Valid() bool {
	//negative values
	if opt.Count < 0 || opt.Height < 0 || opt.Width < 0 {
		return false
	}
	//collections/topics and query can't be used at the same time acc to API documentation
	if (len(opt.CollectionIDs) != 0 || len(opt.TopicIDs) != 0) && opt.SearchQuery != "" {
		return false
	}
	if opt.Count == 0 {
		opt.Count = 1
	}
	if opt.Orientation != "" {
		if opt.Orientation != Landscape && opt.Orientation != Portrait && opt.Orientation != Squarish {
			return false
		}
	}

	return true
}

// Orientation is orientation of a photo
type orientation string

// These constants show possible Orientation types
const (
	Landscape orientation = "landscape"
	Portrait  orientation = "portrait"
	Squarish  orientation = "squarish"
)

var defaultRandomPhotoOpt = &RandomPhotoOpt{Count: 1}

// Random returns random photo(s).
// If opt is nil, then a single random photo is returned by default
func (ps *PhotosService) Random(opt *RandomPhotoOpt) (*[]Photo, *Response, error) {
	if opt == nil {
		opt = defaultRandomPhotoOpt
	}
	if !opt.Valid() {
		return nil, nil, &InvalidListOptError{ErrString: "opt provided is not valid."}
	}
	req, err := newRequest(GET, getEndpoint(photos)+"/random", opt, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err := ps.client.do(req)
	if err != nil {
		return nil, nil, err
	}
	photos := make([]Photo, 0)
	err = json.Unmarshal(*resp.body, &photos)
	if err != nil {
		return nil, nil, err
	}
	return &photos, resp, nil

}

// Like likes a photo on the currently authenticated user's behalf
func (ps *PhotosService) Like(photoID string) (*Photo, *Response, error) {
	if photoID == "" {
		return nil, nil, &IllegalArgumentError{ErrString: "PhotoID cannot be null"}
	}
	endpoint := fmt.Sprintf("%v/%v/like", getEndpoint(photos), photoID)
	req, err := newRequest(POST, endpoint, nil, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err := ps.client.do(req)
	if err != nil {
		return nil, nil, err
	}
	var photo Photo
	err = json.Unmarshal(*resp.body, &photo)
	if err != nil {
		return nil, nil, err
	}
	return &photo, resp, nil
}

// Unlike likes a photo on the currently authenticated user's behalf
func (ps *PhotosService) Unlike(photoID string) (*Photo, *Response, error) {
	if photoID == "" {
		return nil, nil, &IllegalArgumentError{ErrString: "PhotoID cannot be null"}
	}
	endpoint := fmt.Sprintf("%v/%v/like", getEndpoint(photos), photoID)
	req, err := newRequest(DELETE, endpoint, nil, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err := ps.client.do(req)
	if err != nil {
		return nil, nil, err
	}
	var photo Photo
	err = json.Unmarshal(*resp.body, &photo)
	if err != nil {
		return nil, nil, err
	}
	return &photo, resp, nil
}
