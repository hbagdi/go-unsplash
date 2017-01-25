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

	"github.com/stretchr/testify/assert"
)

func TestPhotoOpt(T *testing.T) {
	log.SetOutput(os.Stdout)
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
	log.SetOutput(os.Stdout)
	unsplash := setup()
	photo, err := unsplash.Photos.Photo("", nil)
	assert.NotNil(err)
	assert.Nil(photo)

	photo, err = unsplash.Photos.Photo("random", nil)
	assert.NotNil(photo)
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
	photo, err := unsplash.Photos.Photo("random", &opt)
	assert.NotNil(err)
	assert.Nil(photo)
	log.Println(photo)
	opt.Height = 400
	opt.Width = 600
	photo, err = unsplash.Photos.Photo("random", &opt)
	assert.NotNil(photo)
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
	_, ok := err.(*InvalidListOpt)
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
	log.SetOutput(os.Stdout)
	unsplash := setup()
	stats, err := unsplash.Photos.Stats("-HPhkZcJQNk")
	assert.Nil(err)
	assert.NotNil(stats)
	log.Println(stats)

	stats, err = unsplash.Photos.Stats("")
	assert.Nil(stats)
	assert.NotNil(err)
}

func TestDownloadLink(T *testing.T) {
	assert := assert.New(T)
	log.SetOutput(os.Stdout)
	unsplash := setup()
	url, err := unsplash.Photos.DownloadLink("-HPhkZcJQNk")
	assert.Nil(err)
	assert.NotNil(url)
	log.Println(url)

	url, err = unsplash.Photos.DownloadLink("")
	assert.Nil(url)
	assert.NotNil(err)
}
