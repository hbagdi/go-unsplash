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
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

type AuthConfig struct {
	AppID, Secret, AuthToken string
}

func authFromFile() *AuthConfig {
	bytes, err := ioutil.ReadFile("auth.json")
	if err != nil {
		return nil
	}
	var config AuthConfig
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil
	}
	return &config
}
func getAppAuth() *AuthConfig {
	var config AuthConfig
	appID, ok := os.LookupEnv("unsplash_appID")
	if !ok {
		return nil
	}
	config.AppID = appID
	return &config
}

func getUserAuth() *AuthConfig {
	config := getAppAuth()
	secret, ok := os.LookupEnv("unsplash_secret")
	if !ok {
		return nil
	}
	config.Secret = secret
	token, ok := os.LookupEnv("unsplash_usertoken")
	if !ok {
		return nil
	}
	config.AuthToken = token
	return config
}

func setup() *Unsplash {
	var c *AuthConfig
	c = getUserAuth()
	if c == nil {
		c = authFromFile()
		if c == nil {
			log.Println("Couldn't read auth token. Stopping tests.")
			os.Exit(1)
		}
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.AuthToken},
	)
	client := oauth2.NewClient(oauth2.NoContext, ts)
	return New(client)
}
func TestUnsplash(T *testing.T) {
	assert := assert.New(T)
	unsplash := setup()
	assert.NotNil(unsplash)
	assert.NotNil(unsplash.common)
	assert.NotNil(unsplash.common.httpClient)
	stats, err := unsplash.Stats()
	assert.Nil(err)
	assert.NotNil(stats)
	log.Println(stats)

	unsplash = New(nil)
	stats, err = unsplash.Stats()
	assert.Nil(stats)
	assert.NotNil(err)

	var s service
	_, err = s.do(nil)
	assert.NotNil(err)
}

func TestRogueServer(T *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", getEndpoint(base)+getEndpoint(currentUser),
		httpmock.NewStringResponder(200, `Bad ass Bug flow`))
	httpmock.RegisterResponder("GET", getEndpoint(base)+getEndpoint(globalStats),
		httpmock.NewStringResponder(200, `Bad ass Bug flow`))
	unsplash := setup()
	assert := assert.New(T)
	user, err := unsplash.CurrentUser()
	assert.Nil(user)
	assert.NotNil(err)
	log.SetOutput(os.Stdout)
	log.Println(err)
	stats, err := unsplash.Stats()
	assert.Nil(stats)
	assert.NotNil(err)
}

func TestRogueNetwork(T *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", getEndpoint(base)+getEndpoint(currentUser),
		nil)
	httpmock.RegisterResponder("GET", getEndpoint(base)+getEndpoint(globalStats),
		nil)
	unsplash := setup()
	assert := assert.New(T)
	user, err := unsplash.CurrentUser()
	assert.Nil(user)
	assert.NotNil(err)
	log.SetOutput(os.Stdout)
	log.Println(err)
	stats, err := unsplash.Stats()
	assert.Nil(stats)
	assert.NotNil(err)
}
