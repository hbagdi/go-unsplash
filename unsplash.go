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
	"log"
	"net/http"
)

// Unsplash wraps the entire Unsplash.com API
type Unsplash struct {
	httpClient *http.Client
	//TODO add rate limit struct
}

//New returns a new Unsplash struct
func New(client *http.Client) *Unsplash {
	unsplash := new(Unsplash)
	if client == nil {
		unsplash.httpClient = http.DefaultClient
	} else {
		unsplash.httpClient = client
	}
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
	rawResp, err := u.httpClient.Do(req.Request)
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
		e := new(JSONUnmarshallingError)
		e.ErrString = err.Error()
		return nil,
			&JSONUnmarshallingError{ErrString: err.Error()}
	}
	log.Println()
	log.Println(user)
	return user, nil
}

// List is a temporary crude test
// func (u *Unsplash) List() {
// 	req, err := http.NewRequest("GET", apiURL+"photos", nil)
// 	if err != nil {
// 		log.Println(err)
// 		os.Exit(1)
// 	}
// 	req.Header.Set("Authorization", "Client-ID "+u.Config.AppID)
// 	res, err := u.httpClient.Do(req)
// 	if err != nil {
// 		log.Println(err)
// 		os.Exit(1)
// 	}
// 	defer res.Body.Close()
// 	log.Println(res.Status)
// 	body, _ := ioutil.ReadAll(res.Body)
// 	log.Println(string(body))
// 	res.Header.Write(os.Stdout)
// 	//log.Println(res.Header.Write(w))
// }
