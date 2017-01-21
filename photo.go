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

import "time"

//ExifData for an image
type ExifData struct {
	Make         string `json:"make"`
	Model        string `json:"model"`
	ExposureTime string `json:"exposure_time"`
	Aperture     string `json:"aperture"`
	FocalLength  string `json:"focal_length"`
	Iso          int    `json:"iso"`
}

// Photo represents a photo on unsplash.com
type Photo struct {
	ID                     *string       `json:"id"`
	CreatedAt              *time.Time    `json:"created_at"`
	Width                  int           `json:"width"`
	Height                 int           `json:"height"`
	Color                  *string       `json:"color"`
	Downloads              int           `json:"downloads"`
	Likes                  int           `json:"likes"`
	LikedByUser            bool          `json:"liked_by_user"`
	Exif                   *ExifData     `json:"exif"`
	CurrentUserCollections []interface{} `json:"current_user_collections"`
	Urls                   *struct {
		Raw     *URL `json:"raw"`
		Full    *URL `json:"full"`
		Regular *URL `json:"regular"`
		Small   *URL `json:"small"`
		Thumb   *URL `json:"thumb"`
	} `json:"urls"`
	Links *struct {
		Self             *URL `json:"self"`
		HTML             *URL `json:"html"`
		Download         *URL `json:"download"`
		DownloadLocation *URL `json:"download_location"`
	} `json:"links"`
	Photographer *User `json:"user"`
}
