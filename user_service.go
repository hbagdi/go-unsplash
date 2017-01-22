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
	"encoding/json"
	"fmt"
)

// ProfileImageOpt denotes properties of any Image
type ProfileImageOpt struct {
	Height int `json:"h,omitempty" url:"h"`
	Width  int `json:"w,omitempty" url:"w"`
}

// UserService interacts with /users endpoint
type UserService service

// User returns a User with username and optional profile image size ImageOpt
func (us *UserService) User(username string, imageOpt *ProfileImageOpt) (*User, error) {
	if "" == username {
		return nil, &IllegalArgumentError{ErrString: "Username cannot be null"}
	}
	endpoint := fmt.Sprintf("%v/%v", getEndpoint(users), username)
	req, err := newRequest(GET, endpoint, imageOpt, nil)
	if err != nil {
		return nil, err
	}
	cli := (service)(*us)
	resp, err := cli.do(req)
	if err != nil {
		return nil, err
	}
	var user User
	err = json.Unmarshal(*resp.body, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

type portfolioResponse struct {
	URL *URL `json:"url"`
}

// Portfolio returns a User with username and optional profile image size ImageOpt
func (us *UserService) Portfolio(username string) (*URL, error) {
	if "" == username {
		return nil, &IllegalArgumentError{ErrString: "Username cannot be null"}
	}
	endpoint := fmt.Sprintf("%v/%v/portfolio", getEndpoint(users), username)
	req, err := newRequest(GET, endpoint, nil, nil)
	if err != nil {
		return nil, err
	}
	cli := (service)(*us)
	resp, err := cli.do(req)
	if err != nil {
		return nil, err
	}
	var portfolio portfolioResponse
	err = json.Unmarshal(*resp.body, &portfolio)
	if err != nil {
		return nil, err
	}
	return portfolio.URL, nil
}

// LikedPhotos return an array of liked photos
func (us *UserService) LikedPhotos(username string, opt *ListOpt) (*[]Photo, *Response, error) {
	if "" == username {
		return nil, nil, &IllegalArgumentError{ErrString: "Username cannot be null"}
	}
	if nil == opt {
		opt = defaultListOpt
	}
	if !opt.Valid() {
		return nil, nil, &InvalidListOpt{ErrString: "opt provided is not valid."}
	}
	endpoint := fmt.Sprintf("%v/%v/likes", getEndpoint(users), username)
	req, err := newRequest(GET, endpoint, opt, nil)
	if err != nil {
		return nil, nil, err
	}
	cli := (service)(*us)
	resp, err := cli.do(req)
	if err != nil {
		return nil, nil, err
	}
	likedPhotos := make([]Photo, 0)
	err = json.Unmarshal(*resp.body, &likedPhotos)
	if err != nil {
		return nil, nil, err
	}
	return &likedPhotos, resp, nil
}
