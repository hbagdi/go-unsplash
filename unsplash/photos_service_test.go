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
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestPhotoOpt(T *testing.T) {
	log.SetOutput(ioutil.Discard)
	assert := assert.New(T)
	var opt, opt2 PhotoOpt
	assert.Equal(false, opt.Valid())

	opt.Height = -42
	assert.Equal(false, opt.Valid())
	opt.Crop = true
	assert.Equal(false, opt.Valid())
	opt.Height = 0
	assert.Equal(false, opt.Valid())
	opt.Height = 100
	opt.Width = 100
	assert.Equal(true, opt.Valid())
	opt.CropX = 3
	opt.CropY = 14
	assert.Equal(true, opt.Valid())
	opt.Crop = true
	v := processPhotoOpt(&opt)
	opt0, ok := v.(rect)
	assert.Equal(true, ok)
	assert.Equal("3,14,100,100", opt0.Rect)
	opt2.Width = 42
	opt2.Height = 42
	v = processPhotoOpt(&opt2)
	_, ok = v.(*PhotoOpt)
	assert.Equal(true, ok)
}

func TestSimplePhoto(T *testing.T) {
	log.SetOutput(ioutil.Discard)
	assert := assert.New(T)
	assert.Nil(nil)
	log.SetOutput(ioutil.Discard)
	unsplash := setup()
	photo, resp, err := unsplash.Photos.Photo("", nil)
	assert.NotNil(err)
	assert.Nil(photo)
	assert.Nil(resp)

	photo, resp, err = unsplash.Photos.Photo("random", nil)
	assert.NotNil(photo)
	assert.NotNil(resp)
	assert.Nil(err)
	log.Println(photo)

	json, err := json.Marshal(photo)
	assert.Nil(err)
	log.Println(string(json))
}

func TestPhotoWithOpt(T *testing.T) {
	log.SetOutput(ioutil.Discard)
	assert := assert.New(T)
	var opt PhotoOpt
	unsplash := setup()
	photo, resp, err := unsplash.Photos.Photo("random", &opt)
	assert.NotNil(err)
	assert.Nil(resp)
	assert.Nil(photo)
	log.Println(photo)
	opt.Height = 400
	opt.Width = 600
	photo, resp, err = unsplash.Photos.Photo("random", &opt)
	assert.NotNil(photo)
	assert.NotNil(resp)
	assert.Nil(err)
	log.Println(photo)
}

func TestAllPhotos(T *testing.T) {
	assert := assert.New(T)
	log.SetOutput(ioutil.Discard)
	unsplash := setup()
	photos, resp, err := unsplash.Photos.All(nil)
	assert.Nil(err)
	//check pagination
	assert.NotNil(resp)
	log.Println(resp)
	assert.Equal(true, resp.HasNextPage)
	assert.Equal(2, resp.NextPage)
	lastPage := resp.LastPage
	//check photos
	assert.NotNil(photos)
	assert.Equal(10, len(*photos))

	opt := *defaultListOpt
	opt.Page = 2
	photos, resp, err = unsplash.Photos.All(&opt)
	assert.Nil(err)
	log.Println(err)
	assert.NotNil(resp)
	log.Println(resp)
	assert.Equal(true, resp.HasNextPage)
	assert.Equal(3, resp.NextPage)
	assert.Equal(1, resp.PrevPage)
	assert.Equal(lastPage, resp.LastPage)
	assert.NotNil(photos)
	assert.Equal(10, len(*photos))

	photos, resp, err = unsplash.Photos.All(&ListOpt{PerPage: -1})
	assert.Nil(photos)
	assert.Nil(resp)
	assert.NotNil(err)
	_, ok := err.(*InvalidListOptError)
	assert.Equal(true, ok)

}
func TestCuratedPhotos(T *testing.T) {
	assert := assert.New(T)
	//TODO write better tests
	log.SetOutput(ioutil.Discard)
	unsplash := setup()
	_, resp, err := unsplash.Photos.Curated(nil)
	assert.Nil(err)
	//check pagination
	assert.NotNil(resp)
	log.Println(resp)
}

func TestPhotoStats(T *testing.T) {
	assert := assert.New(T)
	log.SetOutput(ioutil.Discard)
	unsplash := setup()
	stats, resp, err := unsplash.Photos.Stats("-HPhkZcJQNk")
	assert.Nil(err)
	assert.NotNil(stats)
	assert.NotNil(resp)
	log.Println(stats)

	stats, resp, err = unsplash.Photos.Stats("")
	assert.Nil(stats)
	assert.Nil(resp)
	assert.NotNil(err)
}

func TestDownloadLink(T *testing.T) {
	assert := assert.New(T)
	log.SetOutput(ioutil.Discard)
	unsplash := setup()
	url, resp, err := unsplash.Photos.DownloadLink("-HPhkZcJQNk")
	assert.Nil(err)
	assert.NotNil(url)
	assert.NotNil(resp)
	log.Println(url)

	url, resp, err = unsplash.Photos.DownloadLink("")
	assert.Nil(url)
	assert.Nil(resp)
	assert.NotNil(err)
}

func TestRandomPhotoOpt(T *testing.T) {
	assert := assert.New(T)
	var opt RandomPhotoOpt
	opt.CollectionIDs = []int{42}
	opt.SearchQuery = "Gopher"
	assert.Equal(false, opt.Valid())

	var opt2 RandomPhotoOpt
	assert.Equal(true, opt2.Valid())

	opt2.Orientation = "gopher"
	assert.Equal(false, opt2.Valid())
}
func TestRandomPhoto(T *testing.T) {
	assert := assert.New(T)
	log.SetOutput(ioutil.Discard)
	unsplash := setup()
	photos, resp, err := unsplash.Photos.Random(nil)
	assert.Nil(err)
	assert.NotNil(photos)
	assert.NotNil(resp)
	assert.Equal(1, len(*photos))
	photo := (*photos)[0]
	log.Println(photo.String())
	user := photo.Photographer
	log.Println(user.String())
	var opt RandomPhotoOpt
	opt.Count = 3
	opt.SearchQuery = "Earth"
	opt.Orientation = Landscape
	photos, resp, err = unsplash.Photos.Random(&opt)
	assert.Nil(err)
	assert.NotNil(photos)
	assert.NotNil(resp)
	assert.Equal(3, len(*photos))

	var opt2 RandomPhotoOpt
	opt2.Count = 3
	opt2.CollectionIDs = []int{151842, 203782}
	photos, resp, err = unsplash.Photos.Random(&opt2)
	assert.Nil(err)
	assert.NotNil(photos)
	assert.NotNil(resp)
	assert.Equal(3, len(*photos))

	opt.Count = -1
	photos, resp, err = unsplash.Photos.Random(&opt)
	assert.NotNil(err)
	assert.Nil(photos)
	assert.Nil(resp)
	//log.Println(photos)

}

func TestPhotoLike(T *testing.T) {
	assert := assert.New(T)
	log.SetOutput(ioutil.Discard)
	unsplash := setup()
	photos, resp, err := unsplash.Photos.Random(nil)
	assert.Nil(err)
	assert.NotNil(photos)
	assert.NotNil(resp)
	assert.Equal(1, len(*photos))
	photoid := (*photos)[0].ID
	photo, resp, err := unsplash.Photos.Like(*photoid)
	assert.Nil(err)
	assert.NotNil(photo)
	assert.NotNil(resp)

	photo, resp, err = unsplash.Photos.Like("")
	assert.NotNil(err)
	assert.Nil(photo)
	assert.Nil(resp)
}

func TestPhotoUnlike(T *testing.T) {
	assert := assert.New(T)
	log.SetOutput(ioutil.Discard)
	unsplash := setup()
	photos, resp, err := unsplash.Photos.Random(nil)
	assert.Nil(err)
	assert.NotNil(photos)
	assert.NotNil(resp)
	assert.Equal(1, len(*photos))
	photoid := (*photos)[0].ID
	photo, resp, err := unsplash.Photos.Like(*photoid)
	assert.Nil(err)
	assert.NotNil(photo)
	log.Println(photo.String())
	assert.NotNil(resp)

	photo2, resp, err := unsplash.Photos.Unlike(*photoid)
	assert.Nil(err)
	assert.NotNil(photo2)
	assert.NotNil(resp)
	assert.Equal(photo.ID, photo2.ID)
	assert.Equal(photo.Color, photo2.Color)

	photo, resp, err = unsplash.Photos.Like("")
	assert.NotNil(err)
	assert.Nil(photo)
	assert.Nil(resp)
}

func roguePhotoServiceTest(T *testing.T, responder httpmock.Responder) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	log.SetOutput(ioutil.Discard)

	httpmock.RegisterResponder("GET", getEndpoint(base)+getEndpoint(photos)+"/gopherPhoto",
		responder)
	httpmock.RegisterResponder("GET", getEndpoint(base)+getEndpoint(photos)+"/gopherPhoto/stats",
		responder)
	httpmock.RegisterResponder("GET", getEndpoint(base)+getEndpoint(photos)+"/gopherPhoto/download",
		responder)
	httpmock.RegisterResponder("GET", getEndpoint(base)+getEndpoint(photos)+"/random?count=1",
		responder)
	httpmock.RegisterResponder("POST", getEndpoint(base)+getEndpoint(photos)+"/gopherPhoto/like",
		responder)
	httpmock.RegisterResponder("DELETE", getEndpoint(base)+getEndpoint(photos)+"/gopherPhoto/like",
		responder)
	httpmock.RegisterResponder("GET", getEndpoint(base)+getEndpoint(photos),
		responder)

	unsplash := setup()
	assert := assert.New(T)
	photo, resp, err := unsplash.Photos.Photo("gopherPhoto", nil)
	assert.Nil(photo)
	assert.Nil(resp)
	assert.NotNil(err)
	log.Println(err)

	photoStats, resp, err := unsplash.Photos.Stats("gopherPhoto")
	assert.Nil(photoStats)
	assert.Nil(resp)
	assert.NotNil(err)
	log.Println(err)

	url, resp, err := unsplash.Photos.DownloadLink("gopherPhoto")
	assert.Nil(url)
	assert.Nil(resp)
	assert.NotNil(err)
	log.Println(err)

	photos, resp, err := unsplash.Photos.Random(nil)
	assert.Nil(photos)
	assert.Nil(resp)
	assert.NotNil(err)
	log.Println(err)

	photo, resp, err = unsplash.Photos.Like("gopherPhoto")
	assert.Nil(photo)
	assert.Nil(resp)
	assert.NotNil(err)
	log.Println(err)

	photo, resp, err = unsplash.Photos.Unlike("gopherPhoto")
	assert.Nil(photo)
	assert.Nil(resp)
	assert.NotNil(err)
	log.Println(err)

	photos, resp, err = unsplash.Photos.All(nil)
	assert.Nil(photos)
	assert.Nil(resp)
	assert.NotNil(err)
	log.Println(err)
}

func TestPhotoServiceRogueStuff(T *testing.T) {
	roguePhotoServiceTest(T, httpmock.NewStringResponder(200, `Bad ass Bug flow`))
	roguePhotoServiceTest(T, nil)
}
