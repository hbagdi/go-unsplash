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

// PhotoService interacts with /photos endpoint
type PhotoService service

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
func (us *PhotoService) Photo(id string, photoOpt *PhotoOpt) (*Photo, error) {
	if "" == id {
		return nil, &IllegalArgumentError{ErrString: "Photo ID cannot be null"}
	}

	// Validation and conversion if necessary of photoOpt
	var opt interface{}
	opt = nil
	if photoOpt != nil {
		if !photoOpt.Valid() {
			return nil, &InvalidPhotoOpt{ErrString: " photoOpt has zero or non-negative values"}
		}
		opt = processPhotoOpt(photoOpt)
	}
	endpoint := fmt.Sprintf("%v/%v", getEndpoint(photos), id)
	req, err := newRequest(GET, endpoint, opt, nil)
	if err != nil {
		return nil, err
	}
	cli := (service)(*us)
	resp, err := cli.do(req)
	if err != nil {
		return nil, err
	}
	var photo Photo
	err = json.Unmarshal(*resp.body, &photo)
	if err != nil {
		return nil, err
	}
	return &photo, nil
}
