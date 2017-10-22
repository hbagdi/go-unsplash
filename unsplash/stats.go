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
	"strconv"
	"strings"
)

// GlobalStats shows the total photo stats of Unsplash.com
type GlobalStats struct {
	TotalPhotos        uint64
	PhotoDownloads     uint64
	Photos             uint64 `json:"photos,omitempty"`
	Downloads          uint64 `json:"downloads,omitempty"`
	Views              uint64 `json:"views,omitempty"`
	Likes              uint64 `json:"likes,omitempty"`
	Photographers      uint64 `json:"photographers,omitempty"`
	Pixels             uint64 `json:"pixels,omitempty"`
	DownloadsPerSecond uint64 `json:"downloads_per_second,omitempty"`
	ViewsPerSecond     uint64 `json:"views_per_second,omitempty"`
	Developers         uint64 `json:"developers,omitempty"`
	Applications       uint64 `json:"applications,omitempty"`
	Requests           uint64 `json:"requests,omitempty"`
}

// processNumber converts a string or float or int representation into an int
// this hack is needed because the API strangely returns a float value in quotes in the JSON response for this endpoint
func processNumber(i interface{}) uint64 {
	var n uint64
	switch v := i.(type) {
	case uint64:
		n = v
	case string:
		s := strings.Split(v, ".")
		n, _ = strconv.ParseUint(s[0], 10, 64)
	case float64:
		n = uint64(v)
	}
	return n
}

// UnmarshalJSON converts a JSON string representation of GlobalStats into a struct
func (gs *GlobalStats) UnmarshalJSON(b []byte) error {
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		return err
	}
	m := f.(map[string]interface{})
	if v, ok := m["photos"]; ok {
		gs.Photos = processNumber(v)
		gs.TotalPhotos = gs.Photos
	}
	if v, ok := m["downloads"]; ok {
		gs.Downloads = processNumber(v)
		gs.PhotoDownloads = gs.Downloads
	}
	if v, ok := m["views"]; ok {
		gs.Views = processNumber(v)
	}
	if v, ok := m["likes"]; ok {
		gs.Likes = processNumber(v)
	}
	if v, ok := m["photographers"]; ok {
		gs.Photographers = processNumber(v)
	}
	if v, ok := m["pixels"]; ok {
		gs.Pixels = processNumber(v)
	}
	if v, ok := m["downloads_per_second"]; ok {
		gs.DownloadsPerSecond = processNumber(v)
	}
	if v, ok := m["views_per_second"]; ok {
		gs.ViewsPerSecond = processNumber(v)
	}
	if v, ok := m["developers"]; ok {
		gs.Developers = processNumber(v)
	}
	if v, ok := m["applications"]; ok {
		gs.Applications = processNumber(v)
	}
	if v, ok := m["requests"]; ok {
		gs.Requests = processNumber(v)
	}
	return nil
}
func (gs *GlobalStats) String() string {
	var buf bytes.Buffer
	buf.WriteString("Global Stats: ")

	buf.WriteString(" Photos[")
	buf.WriteString(strconv.FormatUint(gs.Photos, 10))
	buf.WriteString("]")

	buf.WriteString(" Downloads[")
	buf.WriteString(strconv.FormatUint(gs.Downloads, 10))
	buf.WriteString("]")

	buf.WriteString(" Views[")
	buf.WriteString(strconv.FormatUint(gs.Views, 10))
	buf.WriteString("]")

	buf.WriteString(" Likes[")
	buf.WriteString(strconv.FormatUint(gs.Likes, 10))
	buf.WriteString("]")

	buf.WriteString(" Photographers[")
	buf.WriteString(strconv.FormatUint(gs.Photographers, 10))
	buf.WriteString("]")

	buf.WriteString(" Pixels[")
	buf.WriteString(strconv.FormatUint(gs.Pixels, 10))
	buf.WriteString("]")

	buf.WriteString(" DownloadsPerSecond[")
	buf.WriteString(strconv.FormatUint(gs.DownloadsPerSecond, 10))
	buf.WriteString("]")

	buf.WriteString(" ViewsPerSecond[")
	buf.WriteString(strconv.FormatUint(gs.ViewsPerSecond, 10))
	buf.WriteString("]")

	buf.WriteString(" Developers[")
	buf.WriteString(strconv.FormatUint(gs.Developers, 10))
	buf.WriteString("]")

	buf.WriteString(" Applications[")
	buf.WriteString(strconv.FormatUint(gs.Applications, 10))
	buf.WriteString("]")

	buf.WriteString(" Requests[")
	buf.WriteString(strconv.FormatUint(gs.Requests, 10))
	buf.WriteString("]")
	return buf.String()
}

// MonthStats shows the overall Unsplash stats for the past 30 days.
type MonthStats struct {
	Downloads        uint64 `json:"downloads"`
	Views            uint64 `json:"views"`
	Likes            uint64 `json:"likes"`
	NewPhotos        uint64 `json:"new_photos"`
	NewPhotographers uint64 `json:"new_photographers"`
	NewPixels        uint64 `json:"new_pixels"`
	NewDevelopers    uint64 `json:"new_developers"`
	NewApplications  uint64 `json:"new_applications"`
	NewRequests      uint64 `json:"new_requests"`
}

// UnmarshalJSON converts a JSON string representation of GlobalStats into a struct
func (stats *MonthStats) UnmarshalJSON(b []byte) error {
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		return err
	}
	m := f.(map[string]interface{})
	if v, ok := m["downloads"]; ok {
		stats.Downloads = processNumber(v)
	}
	if v, ok := m["views"]; ok {
		stats.Views = processNumber(v)
	}
	if v, ok := m["likes"]; ok {
		stats.Likes = processNumber(v)
	}
	if v, ok := m["new_photographers"]; ok {
		stats.NewPhotographers = processNumber(v)
	}
	if v, ok := m["new_pixels"]; ok {
		stats.NewPixels = processNumber(v)
	}
	if v, ok := m["new_photos"]; ok {
		stats.NewPhotos = processNumber(v)
	}
	if v, ok := m["new_requests"]; ok {
		stats.NewRequests = processNumber(v)
	}
	if v, ok := m["new_developers"]; ok {
		stats.NewDevelopers = processNumber(v)
	}
	if v, ok := m["new_applications"]; ok {
		stats.NewApplications = processNumber(v)
	}
	return nil
}
func (stats *MonthStats) String() string {
	var buf bytes.Buffer
	buf.WriteString("Monthly Stats: ")

	buf.WriteString(" Downloads[")
	buf.WriteString(strconv.FormatUint(stats.Downloads, 10))
	buf.WriteString("]")

	buf.WriteString(" Views[")
	buf.WriteString(strconv.FormatUint(stats.Views, 10))
	buf.WriteString("]")

	buf.WriteString(" Likes[")
	buf.WriteString(strconv.FormatUint(stats.Likes, 10))
	buf.WriteString("]")

	buf.WriteString(" New Photos[")
	buf.WriteString(strconv.FormatUint(stats.NewPhotos, 10))
	buf.WriteString("]")

	buf.WriteString(" New Photographers[")
	buf.WriteString(strconv.FormatUint(stats.NewPhotographers, 10))
	buf.WriteString("]")

	buf.WriteString(" New Pixels[")
	buf.WriteString(strconv.FormatUint(stats.NewPixels, 10))
	buf.WriteString("]")

	buf.WriteString(" New Applications[")
	buf.WriteString(strconv.FormatUint(stats.NewApplications, 10))
	buf.WriteString("]")

	buf.WriteString(" New Developers[")
	buf.WriteString(strconv.FormatUint(stats.NewDevelopers, 10))
	buf.WriteString("]")

	buf.WriteString(" New Requests[")
	buf.WriteString(strconv.FormatUint(stats.NewRequests, 10))
	buf.WriteString("]")

	return buf.String()
}
