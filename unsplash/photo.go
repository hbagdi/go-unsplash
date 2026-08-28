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
	"time"
)

//ExifData for an image
type ExifData struct {
	Make         *string `json:"make"`
	Model        *string `json:"model"`
	ExposureTime *string `json:"exposure_time"`
	Aperture     *string `json:"aperture"`
	FocalLength  *string `json:"focal_length"`
	Iso          *int    `json:"iso"`
}

// Tag lists can be applied to any photo
type Tag struct {
	Type  *string `json:"type"`
	Title *string `json:"title"`
}

// Photo represents a photo on unsplash.com
type Photo struct {
	ID             *string    `json:"id"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	Width          *int       `json:"width"`
	Height         *int       `json:"height"`
	Color          *string    `json:"color"`
	Description    *string    `json:"description"`
	AltDescription *string    `json:"alt_description"`
	BlurHash       *string    `json:"blur_hash"`
	Views          *int       `json:"views"`
	Downloads      *int       `json:"downloads"`
	Likes          *int       `json:"likes"`
	LikedByUser    *bool      `json:"liked_by_user"`
	Exif           *ExifData  `json:"exif"`
	Photographer   *User      `json:"user"`
	Location       *struct {
		Title    *string `json:"title"`
		Name     *string `json:"name"`
		City     *string `json:"city"`
		Country  *string `json:"country"`
		Position *struct {
			Latitude  *float64 `json:"latitude"`
			Longitude *float64 `json:"longitude"`
		} `json:"position"`
	} `json:"location"`
	Tags                   *[]Tag        `json:"tags"`
	CurrentUserCollections *[]Collection `json:"current_user_collections"`
	Urls                   *struct {
		Raw     *URL `json:"raw"`
		Full    *URL `json:"full"`
		Regular *URL `json:"regular"`
		Small   *URL `json:"small"`
		Thumb   *URL `json:"thumb"`
		Custom  *URL `json:"custom"`
	} `json:"urls"`
	Links *struct {
		Self             *URL `json:"self"`
		HTML             *URL `json:"html"`
		Download         *URL `json:"download"`
		DownloadLocation *URL `json:"download_location"`
	} `json:"links"`
}

func (p *Photo) String() string {
	var buf bytes.Buffer
	if p.ID == nil {
		return "Photo is not valid"
	}
	buf.WriteString("Photo: ID[")
	buf.WriteString(*p.ID)
	buf.WriteString("]")
	return buf.String()
}

//PhotoStats shows various stats of the photo returned by /photos/:id/stats endpoint
type PhotoStats struct {
	Downloads int `json:"downloads"`
	Likes     int `json:"likes"`
	Views     int `json:"views"`
	Links     struct {
		Self             string `json:"self"`
		HTML             string `json:"html"`
		Download         string `json:"download"`
		DownloadLocation string `json:"download_location"`
	} `json:"links"`
}

// PhotoStatistics represents statistics like downloads, views and likes of an unsplash photo
type PhotoStatistics struct {
	ID        string `json:"id"`
	Downloads struct {
		Total      int `json:"total"`
		Historical struct {
			Change     int    `json:"change"`
			Resolution string `json:"resolution"`
			Quantity   int    `json:"quantity"`
			Values     []struct {
				Date  string `json:"date"`
				Value int    `json:"value"`
			} `json:"values"`
		} `json:"historical"`
	} `json:"downloads"`
	Views struct {
		Total      int `json:"total"`
		Historical struct {
			Change     int    `json:"change"`
			Resolution string `json:"resolution"`
			Quantity   int    `json:"quantity"`
			Values     []struct {
				Date  string `json:"date"`
				Value int    `json:"value"`
			} `json:"values"`
		} `json:"historical"`
	} `json:"views"`
	Likes struct {
		Total      int `json:"total"`
		Historical struct {
			Change     int    `json:"change"`
			Resolution string `json:"resolution"`
			Quantity   int    `json:"quantity"`
			Values     []struct {
				Date  string `json:"date"`
				Value int    `json:"value"`
			} `json:"values"`
		} `json:"historical"`
	} `json:"likes"`
}
