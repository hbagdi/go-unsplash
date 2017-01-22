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

import "encoding/json"

// Collection holds a collection on unsplash.com
type Collection struct {
	ID           *int    `json:"id"`
	Title        *string `json:"title"`
	Description  *string `json:"description"`
	PublishedAt  *string `json:"published_at"`
	Curated      *bool   `json:"curated"`
	Featured     *bool   `json:"featured"`
	TotalPhotos  *int    `json:"total_photos"`
	Private      *bool   `json:"private"`
	ShareKey     *string `json:"share_key"`
	CoverPhoto   *Photo  `json:"cover_photo"`
	Photographer *User   `json:"user"`
	Links        *struct {
		Self    *URL `json:"self"`
		HTML    *URL `json:"html"`
		Photos  *URL `json:"photos"`
		Related *URL `json:"related"`
	} `json:"links"`
}

// All returns a list of all collections on unsplash.
// Note that some fields in collection structs from this result will be missing.
// Use Photo() method to get all details of the  Photo.
func (cs *CollectionsService) All(opt *ListOpt) (*[]Collection, *Response, error) {
	return cs.getCollections(opt, "collections")
}

// getCollections can be used to query any endpoint which
//returns an array of Collections
func (cs *CollectionsService) getCollections(opt *ListOpt, endpoint string) (*[]Collection, *Response, error) {
	if nil == opt {
		opt = defaultListOpt
	}
	if !opt.Valid() {
		return nil, nil, &InvalidListOpt{ErrString: "opt provided is not valid."}
	}
	req, err := newRequest(GET, endpoint, opt, nil)
	if err != nil {
		return nil, nil, err
	}
	cli := (service)(*cs)
	resp, err := cli.do(req)
	if err != nil {
		return nil, nil, err
	}
	collections := make([]Collection, 0)
	err = json.Unmarshal(*resp.body, &collections)
	if err != nil {
		return nil, nil, err
	}
	return &collections, resp, nil
}
