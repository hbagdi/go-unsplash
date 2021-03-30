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
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestUserProfile(T *testing.T) {
	T.Skip()
	assert := assert.New(T)
	log.SetOutput(ioutil.Discard)
	unsplash := setup()
	profileImageOpt := &ProfileImageOpt{Height: 120, Width: 400}
	user, err := unsplash.Users.User("lukechesser", profileImageOpt)
	assert.Nil(err)
	assert.NotNil(user)
	log.Println(user)
	assert.NotNil(user.Photos)
	photos := *user.Photos
	assert.NotNil(photos[0])
	log.Println(len(photos))
	photo := photos[0]
	assert.NotNil(&photo)
	pi := user.ProfileImage
	assert.NotNil(pi)
	assert.NotNil(pi.Medium)
	assert.NotNil(pi.Custom)
	log.Println(user.ProfileImage.Custom.String())

	user, err = unsplash.Users.User("hbagdi", nil)
	assert.Nil(err)
	assert.NotNil(user)
	//log.Println(user)
	pi = user.ProfileImage
	assert.NotNil(pi)
	assert.Nil(pi.Custom)

	user, err = unsplash.Users.User("", nil)
	assert.Nil(user)
	assert.NotNil(err)
	iae, ok := err.(*IllegalArgumentError)
	assert.NotNil(iae)
	assert.Equal(true, ok)

	user, err = unsplash.Users.User(" batmanIsNotAuser", nil)
	assert.Nil(user)
	assert.NotNil(err)
	nfe, ok := err.(*NotFoundError)
	assert.NotNil(nfe)
	assert.Equal(true, ok)
}

func TestUserPortfolio(T *testing.T) {
	assert := assert.New(T)
	log.SetOutput(ioutil.Discard)
	unsplash := setup()
	url, err := unsplash.Users.Portfolio("hbagdi")
	assert.Nil(err)
	assert.NotNil(url)
	log.Println("URL is : ", url.String())
	url, err = unsplash.Users.Portfolio("gopher")
	assert.Nil(err)
	assert.NotNil(url)
	assert.Equal(url.String(), "https://wikipedia.org/wiki/Gopher")
	log.Println("URL is : ", url.String())

	url, err = unsplash.Users.Portfolio("")
	assert.Nil(url)
	assert.NotNil(err)
	iae, ok := err.(*IllegalArgumentError)
	assert.NotNil(iae)
	assert.Equal(true, ok)
}

func TestUserStatistics(T *testing.T) {
	assert := assert.New(T)
	log.SetOutput(ioutil.Discard)
	unsplash := setup()
	stats, resp, err := unsplash.Users.Statistics("lukechesser", nil)
	assert.Nil(err)
	assert.NotNil(stats)
	assert.NotNil(resp)
	assert.NotNil(30, stats.Downloads.Historical.Quantity)
	log.Println(stats)

	stats, resp, err = unsplash.Users.Statistics("lukechesser", &StatsOpt{Quantity: 10})
	assert.Nil(err)
	assert.NotNil(stats)
	assert.NotNil(resp)
	assert.NotNil(30, stats.Downloads.Historical.Quantity)
	log.Println(stats)

	stats, resp, err = unsplash.Users.Statistics("lukechesser", &StatsOpt{Resolution: "sd"})
	assert.NotNil(err)
	assert.Nil(resp)

	stats, resp, err = unsplash.Users.Statistics("lukechesser", &StatsOpt{Quantity: 31})
	assert.NotNil(err)
	assert.Nil(resp)

	stats, resp, err = unsplash.Users.Statistics("", nil)
	assert.Nil(stats)
	assert.Nil(resp)
	assert.NotNil(err)
}
func TestLikedPhotos(T *testing.T) {
	assert := assert.New(T)
	log.SetOutput(ioutil.Discard)
	unsplash := setup()
	// hopefully cofounder won't change his username
	photos, resp, err := unsplash.Users.LikedPhotos("lukechesser", nil)
	assert.Nil(err)
	//check pagination
	assert.NotNil(resp)
	log.Println(resp)
	assert.Equal(true, resp.HasNextPage)
	assert.Equal(2, resp.NextPage)
	// lastPage := resp.LastPage
	//check photos
	assert.NotNil(photos)
	assert.Equal(10, len(*photos))

	opt := *defaultListOpt
	opt.Page = 2
	opt.PerPage = 42
	photos, resp, err = unsplash.Users.LikedPhotos("lukechesser", &opt)
	assert.Nil(err)
	log.Println(err)
	assert.NotNil(resp)
	log.Println(resp)
	assert.Equal(true, resp.HasNextPage)
	assert.Equal(3, resp.NextPage)
	assert.Equal(1, resp.PrevPage)
	assert.NotNil(photos)
	assert.Equal(30, len(*photos))

	photos, resp, err = unsplash.Users.LikedPhotos("lukechesser", &ListOpt{PerPage: -1})
	assert.Nil(photos)
	assert.Nil(resp)
	assert.NotNil(err)
	_, ok := err.(*InvalidListOptError)
	assert.Equal(true, ok)

	photos, resp, err = unsplash.Users.LikedPhotos("", nil)
	assert.NotNil(err)
	assert.Nil(photos)
	assert.Nil(resp)
}
func TestUserPhotos(T *testing.T) {
	assert := assert.New(T)
	//TODO write better tests
	log.SetOutput(ioutil.Discard)
	unsplash := setup()
	// hopefully cofounder won't change his username
	_, resp, err := unsplash.Users.Photos("lukechesser", nil)
	assert.Nil(err)
	//check pagination
	assert.NotNil(resp)
	log.Println(resp)

	photos, resp, err := unsplash.Users.Photos("", nil)
	assert.NotNil(err)
	assert.Nil(photos)
	assert.Nil(resp)

}
func TestUserCollections(T *testing.T) {
	T.Skip()
	assert := assert.New(T)
	//TODO write better tests
	log.SetOutput(ioutil.Discard)
	unsplash := setup()
	// hopefully cofounder won't change his username
	_, resp, err := unsplash.Users.Collections("gopher", nil)
	assert.Nil(err)
	//check pagination
	assert.NotNil(resp)
	log.Println(resp)

	photos, resp, err := unsplash.Users.Collections("", nil)
	assert.NotNil(err)
	assert.Nil(photos)
	assert.Nil(resp)
}

func rogueUserServiceTests(T *testing.T, responder httpmock.Responder) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	log.SetOutput(ioutil.Discard)
	httpmock.RegisterResponder("GET", getEndpoint(base)+getEndpoint(users)+"/gopher",
		responder)
	httpmock.RegisterResponder("GET", getEndpoint(base)+getEndpoint(users)+"/gopher/portfolio",
		responder)

	unsplash := setup()
	assert := assert.New(T)
	user, err := unsplash.Users.User("gopher", nil)
	assert.Nil(user)
	assert.NotNil(err)
	log.Println(err)

	url, err := unsplash.Users.Portfolio("gopher")
	assert.Nil(url)
	assert.NotNil(err)
	log.Println(err)
}

func TestUserServiceRogueStuff(T *testing.T) {
	rogueUserServiceTests(T, httpmock.NewStringResponder(200, `Bad ass Bug flow`))
	rogueUserServiceTests(T, nil)
}
