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
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	apiURL = "https://api.unsplash.com/"
)

//AuthConfig defines various authorization header
type AuthConfig struct {
	AppID     string
	Secret    string
	AuthToken string
}

// Unsplash wraps the entire Unsplash.com API
type Unsplash struct {
	httpClient http.Client
	Config     AuthConfig
}

//New returns a new Unsplash struct
func New(config *AuthConfig) (*Unsplash, error) {
	if config == nil {
		return nil, &InvalidAuthCredentialsError{}
	}
	if config.AppID == "" {
		return nil, &InvalidAuthCredentialsError{}
	}
	unsplash := new(Unsplash)
	unsplash.Config = *config
	return unsplash, nil
}

// CurrentUser returns details about the authenticated user
func (u *Unsplash) CurrentUser() (*User, error) {
	req, err := http.NewRequest("GET", apiURL+"me", nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	req.Header.Set("Authorization", "Bearer "+u.Config.AuthToken)
	req.Header.Set("Content-Type", "application/json")
	res, err := u.httpClient.Do(req)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer res.Body.Close()
	log.Println(res.Status)
	body, _ := ioutil.ReadAll(res.Body)
	res.Header.Write(os.Stdout)
	log.Println()
	log.Println(string(body))
	var user User
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&user)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println()
	log.Println(user)
	return nil, err
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
