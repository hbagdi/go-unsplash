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
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var userJSONString = `{
  "uid": "ZVXqdNfJF50",
  "id": "ZVXqdNfJF50",
  "numeric_id": 464648,
  "username": "hbagdi",
  "name": "Hardik Bagdi",
  "first_name": "Hardik",
  "last_name": "Bagdi",
  "portfolio_url": null,
  "bio": null,
  "location": null,
  "total_likes": 0,
  "total_photos": 0,
  "total_collections": 0,
  "followed_by_user": false,
  "following_count": 1,
  "followers_count": 0,
  "downloads": 0,
  "profile_image": {
    "small": "https://images.unsplash.com/placeholder-avatars/extra-large.jpg?ixlib=rb-0.3.5&q=80&fm=jpg&crop=faces&cs=tinysrgb&fit=crop&h=32&w=32&s=0ad68f44c4725d5a3fda019bab9d3edc",
    "medium": "https://images.unsplash.com/placeholder-avatars/extra-large.jpg?ixlib=rb-0.3.5&q=80&fm=jpg&crop=faces&cs=tinysrgb&fit=crop&h=64&w=64&s=356bd4b76a3d4eb97d63f45b818dd358",
    "large": "https://images.unsplash.com/placeholder-avatars/extra-large.jpg?ixlib=rb-0.3.5&q=80&fm=jpg&crop=faces&cs=tinysrgb&fit=crop&h=128&w=128&s=ee8bbf5fb8d6e43aaaa238feae2fe90d"
  },
  "photos": [],
  "completed_onboarding": false,
  "badge": null,
  "links": {
    "self": "https://api.unsplash.com/users/hbagdi",
    "html": "http://unsplash.com/@hbagdi",
    "photos": "https://api.unsplash.com/users/hbagdi/photos",
    "likes": "https://api.unsplash.com/users/hbagdi/likes",
    "portfolio": "https://api.unsplash.com/users/hbagdi/portfolio",
    "following": "https://api.unsplash.com/users/hbagdi/following",
    "followers": "https://api.unsplash.com/users/hbagdi/followers"
  }
}`

func TestUser(T *testing.T) {
	assert := assert.New(T)
	log.SetOutput(os.Stdout)
	assert.Equal(42, 42)
	var user User
	bytes := []byte(userJSONString)
	err := json.Unmarshal(bytes, &user)
	assert.Nil(err)
	_, err = json.Marshal(user)
	assert.Nil(err)
	// TODO enable deep testing once badge and photos structs are inplace
	// log.Println("Unmarshalled bytes:", string(bytes))
	// log.Println("Marshalled bytes:", string(marshalledBytes))
	//assert.JSONEq(userJSONString, string(marshalledBytes))
}

func stripSpaces(s string) string {
	return strings.Replace(s, " ", "", -1)
}
