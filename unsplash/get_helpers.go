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

// getPhotos can be used to query any endpoint which returns an array of Photos
func (s *service) getPhotos(opt *ListOpt, endpoint string) (*[]Photo, *Response, error) {
	if nil == opt {
		opt = defaultListOpt
	}
	if !opt.Valid() {
		return nil, nil, &InvalidListOptError{ErrString: "opt provided is not valid."}
	}
	req, err := newRequest(GET, endpoint, opt, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err := s.client.do(req)
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

// getCollections can be used to query any endpoint which
//returns an array of Collections
func (s *service) getCollections(opt *ListOpt, endpoint string) (*[]Collection, *Response, error) {
	if nil == opt {
		opt = defaultListOpt
	}
	if !opt.Valid() {
		return nil, nil, &InvalidListOptError{ErrString: "opt provided is not valid."}
	}
	req, err := newRequest(GET, endpoint, opt, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err := s.client.do(req)
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
