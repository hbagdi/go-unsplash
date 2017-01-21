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

// ImageOpt denotes properties of any Image
type ImageOpt struct {
	Height int `json:"h,omitempty" url:"h"`
	Width  int `json:"w,omitempty" url:"w"`
}

// UserService interacts with /users endpoint
type UserService service

// User returns a User with username and optional profile image size ImageOpt
func (us *UserService) User(username string, imageOpt *ImageOpt) (*User, error) {
	if "" == username {
		return nil, &IllegalArgumentError{ErrString: "Username cannot be null"}
	}
	var user *User
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
	err = json.Unmarshal(*resp.body, &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

type portfolioResponse struct {
	URL *URL `json:"url"`
}

// Portfolio returns a User with username and optional profile image size ImageOpt
func (us *UserService) Portfolio(username string) (*URL, error) {
	if "" == username {
		return nil, &IllegalArgumentError{ErrString: "Username cannot be null"}
	}
	var portfolio portfolioResponse
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
	err = json.Unmarshal(*resp.body, &portfolio)
	if err != nil {
		return nil, err
	}
	return portfolio.URL, nil
}
