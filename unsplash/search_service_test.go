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
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestSearchPhotos(T *testing.T) {
	log.SetOutput(ioutil.Discard)
	assert := assert.New(T)
	assert.Nil(nil)
	log.SetOutput(ioutil.Discard)
	unsplash := setup()
	var opt SearchOpt
	photos, resp, err := unsplash.Search.Photos(&opt)
	assert.NotNil(err)
	assert.Nil(resp)
	assert.Nil(photos)
	opt.Query = "Nature"
	photos, _, err = unsplash.Search.Photos(&opt)
	log.Println(len(*photos.Results))
	assert.NotNil(photos)
	assert.Nil(err)
	log.Println(photos)

	photos, resp, err = unsplash.Search.Photos(nil)
	assert.NotNil(err)
	assert.Nil(resp)
	assert.Nil(photos)
}

func TestSearchUsers(T *testing.T) {
	log.SetOutput(ioutil.Discard)
	assert := assert.New(T)
	assert.Nil(nil)
	log.SetOutput(ioutil.Discard)
	unsplash := setup()
	var opt SearchOpt
	users, resp, err := unsplash.Search.Users(&opt)
	assert.NotNil(err)
	assert.Nil(resp)
	assert.Nil(users)
	opt.Query = "Nature"
	users, _, err = unsplash.Search.Users(&opt)
	log.Println(len(*users.Results))
	assert.NotNil(users)
	assert.Nil(err)
	log.Println(users)

	users, resp, err = unsplash.Search.Users(nil)
	assert.NotNil(err)
	assert.Nil(resp)
	assert.Nil(users)
}

func TestSearchCollections(T *testing.T) {
	log.SetOutput(ioutil.Discard)
	assert := assert.New(T)
	assert.Nil(nil)
	log.SetOutput(ioutil.Discard)
	unsplash := setup()
	var opt SearchOpt
	collections, resp, err := unsplash.Search.Collections(&opt)
	assert.NotNil(err)
	assert.Nil(resp)
	assert.Nil(collections)
	opt.Query = "Nature"
	collections, _, err = unsplash.Search.Collections(&opt)
	assert.NotNil(collections)
	assert.Nil(err)
	log.Println(collections)
	log.Println(len(*collections.Results))

	collections, resp, err = unsplash.Search.Collections(nil)
	assert.NotNil(err)
	assert.Nil(resp)
	assert.Nil(collections)
}

func rogueSearchServiceTest(T *testing.T, responder httpmock.Responder) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	log.SetOutput(os.Stdout)
	qs := "?page=1&per_page=10&query=InnerPeace"
	httpmock.RegisterResponder("GET", getEndpoint(base)+getEndpoint(searchCollections)+qs,
		responder)
	httpmock.RegisterResponder("GET", getEndpoint(base)+getEndpoint(searchUsers)+qs,
		responder)
	httpmock.RegisterResponder("GET", getEndpoint(base)+getEndpoint(searchPhotos)+qs,
		responder)

	unsplash := setup()
	assert := assert.New(T)
	var opt SearchOpt
	opt.Query = "InnerPeace"
	photos, resp, err := unsplash.Search.Photos(&opt)
	assert.Nil(photos)
	assert.Nil(resp)
	assert.NotNil(err)
	log.Println(err)

	collections, resp, err := unsplash.Search.Collections(&opt)
	assert.Nil(collections)
	assert.Nil(resp)
	assert.NotNil(err)
	log.Println(err)
	users, resp, err := unsplash.Search.Users(&opt)
	assert.Nil(users)
	assert.Nil(resp)
	assert.NotNil(err)
	log.Println(err)
}

func TestSearchRogueStuff(T *testing.T) {
	rogueSearchServiceTest(T, httpmock.NewStringResponder(200, `Bad ass Bug flow`))
	rogueSearchServiceTest(T, nil)
}
