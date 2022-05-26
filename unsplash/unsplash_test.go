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
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type AuthConfig struct {
	AuthToken string
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

func randName(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func getUserAuth() *AuthConfig {
	token, ok := os.LookupEnv("UNSPLASH_USERTOKEN")
	fmt.Println(token)
	if !ok {
		return nil
	}
	return &AuthConfig{
		AuthToken: token,
	}
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
	T.Skip()
	assert := assert.New(T)
	log.SetOutput(ioutil.Discard)
	unsplash := setup()
	assert.NotNil(unsplash)
	assert.NotNil(unsplash.common)
	assert.NotNil(unsplash.client)
	tstats, resp, err := unsplash.TotalStats()
	assert.Nil(err)
	assert.NotNil(tstats)
	assert.NotNil(resp)
	stats, resp, err := unsplash.Stats()
	assert.Nil(err)
	assert.NotNil(stats)
	//FIXME
	if stats.Photos <= 0 || stats.Downloads <= 0 || stats.Views <= 0 || stats.Likes <= 0 || stats.Photographers <= 0 || stats.Pixels <= 0 || stats.DownloadsPerSecond <= 0 || stats.ViewsPerSecond <= 0 || stats.Developers <= 0 || stats.Applications <= 0 || stats.Requests <= 0 {
		assert.Fail("GlobalStats struct has a zero field: %s\n", stats.String())
	}
	assert.NotNil(resp)
	log.Println(stats)

	//Disabling the test since API almost always times out
	// monthlyStats, resp, err := unsplash.MonthStats()
	// assert.Nil(err)
	// assert.NotNil(resp)
	// assert.NotNil(monthlyStats)

	unsplash = New(nil)
	stats, resp, err = unsplash.Stats()
	assert.Nil(stats)
	assert.Nil(resp)
	assert.NotNil(err)

	var s service
	_, err = s.client.do(nil)
	assert.NotNil(err)
}

func TestUnsplashRogueServer(T *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", getEndpoint(base)+getEndpoint(currentUser),
		httpmock.NewStringResponder(200, `Bad ass Bug flow`))
	httpmock.RegisterResponder("GET", getEndpoint(base)+getEndpoint(globalStats),
		httpmock.NewStringResponder(200, `Bad ass Bug flow`))
	unsplash := setup()
	assert := assert.New(T)
	user, resp, err := unsplash.CurrentUser()
	assert.Nil(user)
	assert.Nil(resp)
	assert.NotNil(err)
	log.SetOutput(ioutil.Discard)
	log.Println(err)
	stats, resp, err := unsplash.Stats()
	assert.Nil(stats)
	assert.Nil(resp)
	assert.NotNil(err)
}

func TestUnsplashRogueNetwork(T *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", getEndpoint(base)+getEndpoint(currentUser),
		nil)
	httpmock.RegisterResponder("GET", getEndpoint(base)+getEndpoint(globalStats),
		nil)
	unsplash := setup()
	assert := assert.New(T)
	user, resp, err := unsplash.CurrentUser()
	assert.Nil(user)
	assert.Nil(resp)
	assert.NotNil(err)
	log.SetOutput(ioutil.Discard)
	log.Println(err)
	stats, resp, err := unsplash.Stats()
	assert.Nil(stats)
	assert.Nil(resp)
	assert.NotNil(err)
}

func TestUpdateCurrentUser(T *testing.T) {
	log.SetOutput(ioutil.Discard)
	assert := assert.New(T)
	unsplash := setup()
	assert.NotNil(unsplash)
	newUserName := "lukechesser"

	user, resp, err := unsplash.UpdateCurrentUser(nil)
	assert.NotNil(err)
	assert.Nil(user)
	assert.Nil(resp)
	log.Println(err.Error())

	user, resp, err = unsplash.UpdateCurrentUser(&UserUpdateInfo{Username: newUserName})
	assert.NotNil(err)
	assert.Nil(user)
	assert.Nil(resp)
	log.Println(err.Error())
}

func TestClientIDPropagation(T *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	expected_client_id := "HARDCODED_TEST"

	httpmock.RegisterResponder("GET", getEndpoint(base)+getEndpoint(searchPhotos),
		func(r *http.Request) (*http.Response, error) {
			client_id_header := r.Header.Get("Authorization")

			assert.Equal(T, fmt.Sprintf("Client-ID %v", expected_client_id), client_id_header)

			return httpmock.NewStringResponse(200, `{"total":1,"total_pages":1,"results":[{"id":"eOLpJytrbsQ","created_at":"2014-11-18T14:35:36-05:00","width":4000,"height":3000,"color":"#A7A2A1","blur_hash":"LaLXMa9Fx[D%~q%MtQM|kDRjtRIU","likes":286,"liked_by_user":false,"description":"A man drinking a coffee.","user":{"id":"Ul0QVz12Goo","username":"ugmonk","name":"Jeff Sheldon","first_name":"Jeff","last_name":"Sheldon","instagram_username":"instantgrammer","twitter_username":"ugmonk","portfolio_url":"http://ugmonk.com/","profile_image":{"small":"https://images.unsplash.com/profile-1441298803695-accd94000cac?ixlib=rb-0.3.5&q=80&fm=jpg&crop=faces&cs=tinysrgb&fit=crop&h=32&w=32&s=7cfe3b93750cb0c93e2f7caec08b5a41","medium":"https://images.unsplash.com/profile-1441298803695-accd94000cac?ixlib=rb-0.3.5&q=80&fm=jpg&crop=faces&cs=tinysrgb&fit=crop&h=64&w=64&s=5a9dc749c43ce5bd60870b129a40902f","large":"https://images.unsplash.com/profile-1441298803695-accd94000cac?ixlib=rb-0.3.5&q=80&fm=jpg&crop=faces&cs=tinysrgb&fit=crop&h=128&w=128&s=32085a077889586df88bfbe406692202"},"links":{"self":"https://api.unsplash.com/users/ugmonk","html":"http://unsplash.com/@ugmonk","photos":"https://api.unsplash.com/users/ugmonk/photos","likes":"https://api.unsplash.com/users/ugmonk/likes"}},"current_user_collections":[],"urls":{"raw":"https://images.unsplash.com/photo-1416339306562-f3d12fefd36f","full":"https://hd.unsplash.com/photo-1416339306562-f3d12fefd36f","regular":"https://images.unsplash.com/photo-1416339306562-f3d12fefd36f?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=1080&fit=max&s=92f3e02f63678acc8416d044e189f515","small":"https://images.unsplash.com/photo-1416339306562-f3d12fefd36f?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=400&fit=max&s=263af33585f9d32af39d165b000845eb","thumb":"https://images.unsplash.com/photo-1416339306562-f3d12fefd36f?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=200&fit=max&s=8aae34cf35df31a592f0bef16e6342ef"},"links":{"self":"https://api.unsplash.com/photos/eOLpJytrbsQ","html":"http://unsplash.com/photos/eOLpJytrbsQ","download":"http://unsplash.com/photos/eOLpJytrbsQ/download"}}]}`), nil
		})

	u := NewWithClientID(&http.Client{Timeout: time.Duration(1) * time.Second}, expected_client_id)

	opt := SearchOpt{
		Page:    1,
		PerPage: 1,
		Query:   "nowhere",
	}
	photo, _, err := u.Search.Photos(&opt)
	assert.Equal(T, nil, err)
	assert.NotEqual(T, nil, photo)

	info := httpmock.GetCallCountInfo()
	assert.Equal(T, 1, info[fmt.Sprintf("GET %v%v", getEndpoint(base), getEndpoint(searchPhotos))])
}
