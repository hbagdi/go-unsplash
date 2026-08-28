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

// SearchService interacts with /search endpoint
type SearchService service

// SearchOpt should be used to query /search endpoint
type SearchOpt struct {
	Page        int     `url:"page"`
	PerPage     int     `url:"per_page"`
	Query       string  `url:"query"`
	Color       *string `url:"color,omitempty"`
	Orientation *string `url:"orientation,omitempty"`
}

// Valid validates a SearchOpt
func (opt *SearchOpt) Valid() bool {
	if opt.Query == "" {
		return false
	}
	// default params
	if opt.Page == 0 {
		opt.Page = 1
	}
	if opt.PerPage == 0 {
		opt.PerPage = 10
	}
	return true
}

// Users can be used to query any endpoint which returns an array of users.
func (ss *SearchService) Users(opt *SearchOpt) (*UserSearchResult, *Response, error) {
	if nil == opt {
		return nil, nil, &IllegalArgumentError{ErrString: "SearchOpt cannot be nil"}
	}
	if !opt.Valid() {
		return nil, nil, &InvalidListOptError{ErrString: "Search query cannot be empty."}
	}
	req, err := newRequest(GET, getEndpoint(searchUsers), opt, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err := ss.client.do(req)
	if err != nil {
		return nil, nil, err
	}
	var users UserSearchResult
	err = json.Unmarshal(*resp.body, &users)
	if err != nil {
		return nil, nil, err
	}
	return &users, resp, nil
}

// Photos queries the search endpoint to search for photos.
func (ss *SearchService) Photos(opt *SearchOpt) (*PhotoSearchResult, *Response, error) {
	if nil == opt {
		return nil, nil, &IllegalArgumentError{ErrString: "SearchOpt cannot be nil"}
	}
	if !opt.Valid() {
		return nil, nil, &InvalidListOptError{ErrString: "Search query cannot be empty."}
	}
	req, err := newRequest(GET, getEndpoint(searchPhotos), opt, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err := ss.client.do(req)
	if err != nil {
		return nil, nil, err
	}
	var photos PhotoSearchResult
	err = json.Unmarshal(*resp.body, &photos)
	if err != nil {
		return nil, nil, err
	}
	return &photos, resp, nil
}

// Collections queries the search endpoint to search for collections.
func (ss *SearchService) Collections(opt *SearchOpt) (*CollectionSearchResult, *Response, error) {
	if nil == opt {
		return nil, nil, &IllegalArgumentError{ErrString: "SearchOpt cannot be nil"}
	}
	if !opt.Valid() {
		return nil, nil, &InvalidListOptError{ErrString: "Search query cannot be empty."}
	}
	req, err := newRequest(GET, getEndpoint(searchCollections), opt, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err := ss.client.do(req)
	if err != nil {
		return nil, nil, err
	}
	var collections CollectionSearchResult
	err = json.Unmarshal(*resp.body, &collections)
	if err != nil {
		return nil, nil, err
	}
	return &collections, resp, nil
}
