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
	"io"
	"io/ioutil"
	"net/http"
)

type service struct {
	client *Unsplash
}

// Unsplash wraps the entire Unsplash.com API
type Unsplash struct {
	client_id   string
	client      *http.Client
	common      service
	Users       *UsersService
	Photos      *PhotosService
	Collections *CollectionsService
	Search      *SearchService
}

//New returns a new Unsplash struct
func New(client *http.Client) *Unsplash {
	if client == nil {
		client = http.DefaultClient
	}
	unsplash := new(Unsplash)
	unsplash.client = client
	unsplash.common.client = unsplash
	unsplash.Users = (*UsersService)(&unsplash.common)
	unsplash.Photos = (*PhotosService)(&unsplash.common)
	unsplash.Collections = (*CollectionsService)(&unsplash.common)
	unsplash.Search = (*SearchService)(&unsplash.common)
	return unsplash
}

func NewWithClientID(client *http.Client, client_id string) *Unsplash {
	r := New(client)
	r.client_id = client_id
	return r
}

func (s *Unsplash) do(req *request) (*Response, error) {
	var err error
	//TODO should this be exported?
	if req == nil {
		return nil,
			&IllegalArgumentError{ErrString: "Request object cannot be nil"}
	}
	req.Request.Header.Set("Accept-Version", "v1")
	if s.client_id != "" {
		req.Request.Header.Set("Authorization", fmt.Sprintf("Client-ID %v", s.client_id))
	}
	//Make the request
	client := s.client
	rawResp, err := client.Do(req.Request)
	if rawResp != nil {
		defer rawResp.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	resp, err := newResponse(rawResp)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(ioutil.Discard, rawResp.Body)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CurrentUser returns details about the authenticated user
func (u *Unsplash) CurrentUser() (*User, *Response, error) {
	var err error
	req, err := newRequest(GET, getEndpoint(currentUser), nil, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err := u.do(req)
	if err != nil {
		return nil, nil, err
	}
	user := new(User)
	err = json.Unmarshal(*resp.body, &user)
	if err != nil {
		return nil, nil,
			&JSONUnmarshallingError{ErrString: err.Error()}
	}
	return user, resp, nil
}

// UpdateCurrentUser updates the current user's private data and returns an update User struct
func (u *Unsplash) UpdateCurrentUser(updateInfo *UserUpdateInfo) (*User, *Response, error) {
	if updateInfo == nil {
		return nil, nil, &IllegalArgumentError{ErrString: "updateInfo cannot be null"}
	}
	endpoint := "me"
	req, err := newRequest(PUT, endpoint, updateInfo, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err := u.do(req)
	if err != nil {
		return nil, nil, err
	}
	user := new(User)
	err = json.Unmarshal(*resp.body, &user)
	if err != nil {
		return nil, nil,
			&JSONUnmarshallingError{ErrString: err.Error()}
	}
	return user, resp, nil
}

// Stats gives the total photos,download since the inception of unsplash.com
// This method is DEPRECATED, USE TotalStats()
func (u *Unsplash) Stats() (*GlobalStats, *Response, error) {
	var err error
	req, err := newRequest(GET, getEndpoint(globalStats), nil, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err := u.do(req)
	if err != nil {
		return nil, nil, err
	}
	globalStats := new(GlobalStats)
	err = json.Unmarshal(*resp.body, &globalStats)
	if err != nil {
		return nil, nil,
			&JSONUnmarshallingError{ErrString: err.Error()}
	}
	return globalStats, resp, nil
}

// TotalStats returns various stats related to unsplash.com since it's inception
func (u *Unsplash) TotalStats() (*GlobalStats, *Response, error) {
	return u.Stats()
}

// MonthStats returns various stats related to unsplash.com for last 30 days
func (u *Unsplash) MonthStats() (*MonthStats, *Response, error) {
	var err error
	req, err := newRequest(GET, getEndpoint(monthStats), nil, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err := u.do(req)
	if err != nil {
		return nil, nil, err
	}
	monthStats := new(MonthStats)
	err = json.Unmarshal(*resp.body, &monthStats)
	if err != nil {
		return nil, nil,
			&JSONUnmarshallingError{ErrString: err.Error()}
	}
	return monthStats, resp, nil
}
