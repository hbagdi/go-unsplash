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
	TotalPhotos    int `json:"total_photos,omitempty"`
	PhotoDownloads int `json:"photo_downloads,omitempty"`
	BatchDownloads int `json:"batch_downloads,omitempty"`
}

// processNumber converts a string or float or int representation into an int
// this hack is needed because the API strangely returns a float value in quotes in the JSON response for this endpoint
func processNumber(i interface{}) int {
	switch v := i.(type) {
	case int:
		return v
	case string:
		s := strings.Split(i.(string), ".")
		n, _ := strconv.Atoi(s[0])
		return n
	case float64:
		return int(v)
	}
	return 0
}

// UnmarshalJSON converts a JSON string representation of GlobalStats into a struct
func (gs *GlobalStats) UnmarshalJSON(b []byte) error {
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		return err
	}
	m := f.(map[string]interface{})
	if _, ok := m["total_photos"]; ok {
		gs.TotalPhotos = processNumber(m["total_photos"])
	}
	if _, ok := m["photo_downloads"]; ok {
		gs.PhotoDownloads = processNumber(m["photo_downloads"])
	}
	if _, ok := m["batch_downloads"]; ok {
		gs.BatchDownloads = processNumber(m["batch_downloads"])
	}
	return nil
}
func (gs *GlobalStats) String() string {
	var buf bytes.Buffer
	buf.WriteString("\nGlobal Stats:\n")
	buf.WriteString("Total Photos: " + strconv.Itoa(gs.TotalPhotos) + "\n")
	buf.WriteString("Total downloads: " + strconv.Itoa(gs.PhotoDownloads) + "\n")
	buf.WriteString("Batch downloads: " + strconv.Itoa(gs.BatchDownloads) + "\n")
	return buf.String()
}
