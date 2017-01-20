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
	"net/http"
)

type service struct {
	httpClient *http.Client
}

// Unsplash wraps the entire Unsplash.com API
type Unsplash struct {
	common *service
	Users  *UserService
	//TODO add rate limit struct
}

//New returns a new Unsplash struct
func New(client *http.Client) *Unsplash {
	unsplash := new(Unsplash)
	unsplash.common = new(service)
	if client == nil {
		unsplash.common.httpClient = http.DefaultClient
	} else {
		unsplash.common.httpClient = client
	}
	unsplash.Users = (*UserService)(unsplash.common)
	return unsplash
}

func (u *Unsplash) do(req *request) (*response, error) {
	var err error
	//TODO should this be exported?
	if req == nil {
		return nil,
			&IllegalArgumentError{ErrString: "Request object cannot be nil"}
	}
	//TODO add rate limiting support, API is erronous at the moment

	//Make the request
	client := u.common.httpClient
	rawResp, err := client.Do(req.Request)
	if err != nil {
		return nil, err
	}
	resp, err := newResponse(rawResp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CurrentUser returns details about the authenticated user
func (u *Unsplash) CurrentUser() (*User, error) {
	var err error
	req, err := newRequest(GET, currentUser, nil)
	if err != nil {
		return nil, err
	}
	resp, err := u.do(req)
	if err != nil {
		return nil, err
	}
	user := new(User)
	err = json.Unmarshal(*resp.body, &user)
	if err != nil {
		return nil,
			&JSONUnmarshallingError{ErrString: err.Error()}
	}
	return user, nil
}

// Stats gives the total photos,download since the inception of unsplash.com
func (u *Unsplash) Stats() (*GlobalStats, error) {
	var err error
	req, err := newRequest(GET, globalStats, nil)
	if err != nil {
		return nil, err
	}
	resp, err := u.do(req)
	if err != nil {
		return nil, err
	}
	globalStats := new(GlobalStats)
	err = json.Unmarshal(*resp.body, &globalStats)
	if err != nil {
		return nil,
			&JSONUnmarshallingError{ErrString: err.Error()}
	}
	return globalStats, nil
}
