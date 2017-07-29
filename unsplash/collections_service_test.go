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
	"math/rand"
	"strconv"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestAllCollections(T *testing.T) {
	assert := assert.New(T)
	log.SetOutput(ioutil.Discard)
	unsplash := setup()
	collections, resp, err := unsplash.Collections.All(nil)
	assert.Nil(err)
	//check pagination
	assert.NotNil(resp)
	log.Println(resp)
	assert.Equal(true, resp.HasNextPage)
	assert.Equal(2, resp.NextPage)
	lastPage := resp.LastPage
	//check collections
	assert.NotNil(collections)
	assert.Equal(10, len(*collections))

	opt := *defaultListOpt
	opt.Page = 2
	collections, resp, err = unsplash.Collections.All(&opt)
	assert.Nil(err)
	log.Println(err)
	assert.NotNil(resp)
	log.Println(resp)
	assert.Equal(true, resp.HasNextPage)
	assert.Equal(3, resp.NextPage)
	assert.Equal(1, resp.PrevPage)
	assert.Equal(lastPage, resp.LastPage)
	assert.NotNil(collections)
	assert.Equal(10, len(*collections))

	collections, resp, err = unsplash.Collections.All(&ListOpt{PerPage: -1})
	assert.Nil(collections)
	assert.Nil(resp)
	assert.NotNil(err)
	_, ok := err.(*InvalidListOptError)
	assert.Equal(true, ok)

}

func TestFeaturedCollections(T *testing.T) {
	assert := assert.New(T)
	log.SetOutput(ioutil.Discard)
	unsplash := setup()
	collections, resp, err := unsplash.Collections.Featured(nil)
	assert.Nil(err)
	//check pagination
	assert.NotNil(resp)
	log.Println(resp)
	assert.Equal(true, resp.HasNextPage)
	assert.Equal(2, resp.NextPage)
	assert.NotNil(collections)
	assert.Equal(10, len(*collections))
}

func TestCuratedCollections(T *testing.T) {
	assert := assert.New(T)
	log.SetOutput(ioutil.Discard)
	unsplash := setup()
	collections, resp, err := unsplash.Collections.Curated(nil)
	assert.Nil(err)
	//check pagination
	assert.NotNil(resp)
	log.Println(resp)
	assert.Equal(true, resp.HasNextPage)
	assert.Equal(2, resp.NextPage)
	assert.NotNil(collections)
	assert.Equal(10, len(*collections))
}

func TestRelatedCollections(T *testing.T) {
	assert := assert.New(T)
	log.SetOutput(ioutil.Discard)
	unsplash := setup()
	collections, resp, err := unsplash.Collections.Related("296", nil)
	assert.Nil(err)
	//check pagination
	assert.NotNil(resp)
	assert.NotNil(collections)
	log.Println(resp)

	collections, resp, err = unsplash.Collections.Related("", nil)
	assert.NotNil(err)
	assert.Nil(collections)
	assert.Nil(resp)
}

func TestSimpleCollection(T *testing.T) {
	assert := assert.New(T)
	log.SetOutput(ioutil.Discard)
	unsplash := setup()
	collection, resp, err := unsplash.Collections.Collection("910")
	assert.Nil(err)
	assert.NotNil(resp)
	assert.NotNil(collection)
	log.Println(resp)

	collection, resp, err = unsplash.Collections.Collection("")
	assert.NotNil(err)
	assert.Nil(collection)
	assert.Nil(resp)
}

func TestCreateCollection(T *testing.T) {
	assert := assert.New(T)
	log.SetOutput(ioutil.Discard)
	unsplash := setup()
	var opt CollectionOpt
	title := "Test42"
	opt.Title = &title
	collection, resp, err := unsplash.Collections.Create(&opt)
	assert.Nil(err)
	assert.NotNil(resp)
	assert.NotNil(collection)
	resp, err = unsplash.Collections.Delete(*collection.ID)
	assert.NotNil(resp)
	assert.Nil(err)

	title = ""
	collection, resp, err = unsplash.Collections.Create(&opt)
	assert.Nil(resp)
	assert.Nil(collection)
	assert.NotNil(err)

	collection, resp, err = unsplash.Collections.Create(nil)
	assert.Nil(resp)
	assert.Nil(collection)
	assert.NotNil(err)
}

func TestUpdateCollection(T *testing.T) {
	assert := assert.New(T)
	log.SetOutput(ioutil.Discard)
	unsplash := setup()

	//get a user's collection
	collections, resp, err := unsplash.Users.Collections("gopher", nil)
	assert.Nil(err)
	assert.NotNil(resp)
	assert.NotNil(collections)
	collection := (*collections)[0]
	assert.NotNil(collection)
	log.Println(*collection.ID)
	//random title
	var opt CollectionOpt
	title := "Test43" + strconv.Itoa(rand.Int())
	opt.Title = &title
	col, resp, err := unsplash.Collections.Update(*collection.ID, &opt)
	assert.Nil(err)
	assert.NotNil(resp)
	assert.NotNil(col)

	col, resp, err = unsplash.Collections.Update(0, &opt)
	assert.Nil(resp)
	assert.Nil(col)
	assert.NotNil(err)

	col, resp, err = unsplash.Collections.Update(246, nil)
	assert.Nil(resp)
	assert.Nil(col)
	assert.NotNil(err)

	col, resp, err = unsplash.Collections.Update(0, nil)
	assert.Nil(resp)
	assert.Nil(col)
	assert.NotNil(err)
}

func TestDeleteCollection(T *testing.T) {
	assert := assert.New(T)
	log.SetOutput(ioutil.Discard)
	unsplash := setup()
	var opt CollectionOpt
	title := "Test42"
	opt.Title = &title
	collection, resp, err := unsplash.Collections.Create(&opt)
	assert.Nil(err)
	assert.NotNil(resp)
	assert.NotNil(collection)

	resp, err = unsplash.Collections.Delete(*collection.ID)
	assert.NotNil(resp)
	assert.Nil(err)

	resp, err = unsplash.Collections.Delete(0)
	assert.NotNil(err)
	assert.Nil(resp)
}

func TestAddPhoto(T *testing.T) {
	assert := assert.New(T)
	log.SetOutput(ioutil.Discard)
	unsplash := setup()

	//get a random photo
	photos, resp, err := unsplash.Photos.Random(nil)
	assert.Nil(err)
	assert.NotNil(resp)
	assert.NotNil(photos)
	assert.Equal(1, len(*photos))
	photo := (*photos)[0]
	assert.NotNil(photo)

	//get a user's collection
	collections, resp, err := unsplash.Users.Collections("gopher", nil)
	assert.Nil(err)
	assert.NotNil(resp)
	assert.NotNil(collections)

	collection := (*collections)[0]
	assert.NotNil(collection)

	//add the photo
	resp, err = unsplash.Collections.AddPhoto(*collection.ID, *photo.ID)
	assert.Nil(err)
	assert.NotNil(resp)

	//empty things
	resp, err = unsplash.Collections.AddPhoto(0, "photoID")
	assert.NotNil(err)
	assert.Nil(resp)
	resp, err = unsplash.Collections.AddPhoto(296, "")
	assert.NotNil(err)
	assert.Nil(resp)
	resp, err = unsplash.Collections.AddPhoto(0, "")
	assert.NotNil(err)
	assert.Nil(resp)

}

func TestRemovePhoto(T *testing.T) {
	assert := assert.New(T)
	log.SetOutput(ioutil.Discard)
	unsplash := setup()

	//get a random photo
	photos, resp, err := unsplash.Photos.Random(nil)
	assert.Nil(err)
	assert.NotNil(resp)
	assert.NotNil(photos)
	assert.Equal(1, len(*photos))
	photo := (*photos)[0]
	assert.NotNil(photo)

	//get a user's collection
	collections, resp, err := unsplash.Users.Collections("gopher", nil)
	assert.Nil(err)
	assert.NotNil(resp)
	assert.NotNil(collections)

	collection := (*collections)[0]
	assert.NotNil(collection)

	//add the photo
	resp, err = unsplash.Collections.AddPhoto(*collection.ID, *photo.ID)
	assert.Nil(err)
	assert.NotNil(resp)

	//remove the photo
	_, _ = unsplash.Collections.RemovePhoto(*collection.ID, *photo.ID)
	// API is being unreliable at the moment. Returns 403 sometimes
	// could be because of back-to-back requests?
	// assert.Nil(err)
	// assert.NotNil(resp)

	//empty stuff
	//empty things
	resp, err = unsplash.Collections.RemovePhoto(0, "photoID")
	assert.NotNil(err)
	assert.Nil(resp)
	resp, err = unsplash.Collections.RemovePhoto(296, "")
	assert.NotNil(err)
	assert.Nil(resp)
	resp, err = unsplash.Collections.RemovePhoto(0, "")
	assert.NotNil(err)
	assert.Nil(resp)
}

func rogueCollectionServiceTest(T *testing.T, responder httpmock.Responder) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	log.SetOutput(ioutil.Discard)

	httpmock.RegisterResponder("GET", getEndpoint(base)+getEndpoint(collections)+"/gopherCollection",
		responder)
	httpmock.RegisterResponder("POST", getEndpoint(base)+getEndpoint(collections)+"?title=gopherCollection",
		responder)
	httpmock.RegisterResponder("PUT", getEndpoint(base)+getEndpoint(collections)+"/4242?title=gopherCollection",
		responder)
	httpmock.RegisterResponder("POST", getEndpoint(base)+getEndpoint(collections)+"/4242/add?photo_id=gopherPhoto",
		responder)
	httpmock.RegisterResponder("DELETE", getEndpoint(base)+getEndpoint(collections)+"/4242/remove?photo_id=gopherPhoto",
		responder)
	httpmock.RegisterResponder("DELETE", getEndpoint(base)+getEndpoint(collections)+"/4242",
		responder)
	httpmock.RegisterResponder("GET", getEndpoint(base)+getEndpoint(collections),
		responder)

	unsplash := setup()
	assert := assert.New(T)
	collection, resp, err := unsplash.Collections.Collection("gopherCollection")
	assert.Nil(collection)
	assert.Nil(resp)
	assert.NotNil(err)
	log.Println(err)

	var opt CollectionOpt
	title := "gopherCollection"
	opt.Title = &title
	collection, resp, err = unsplash.Collections.Create(&opt)
	assert.Nil(collection)
	assert.Nil(resp)
	assert.NotNil(err)
	log.Println(err)

	collection, resp, err = unsplash.Collections.Update(4242, &opt)
	assert.Nil(collection)
	assert.Nil(resp)
	assert.NotNil(err)
	log.Println(err)

	resp, err = unsplash.Collections.Delete(4242)
	assert.Nil(resp)
	assert.NotNil(err)
	log.Println(err)

	resp, err = unsplash.Collections.AddPhoto(4242, "gopherPhoto")
	assert.NotNil(err)
	assert.Nil(resp)
	log.Println(err)

	cols, resp, err := unsplash.Collections.All(nil)
	assert.Nil(cols)
	assert.Nil(resp)
	assert.NotNil(err)
	log.Println(err)
}

func TestCollectionServiceRogueStuff(T *testing.T) {
	rogueCollectionServiceTest(T, httpmock.NewStringResponder(200, `Bad ass Bug flow`))
	rogueCollectionServiceTest(T, nil)
}

func TestRemovePhotoRogue(T *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	log.SetOutput(ioutil.Discard)

	httpmock.RegisterResponder("DELETE", getEndpoint(base)+getEndpoint(collections)+"/4242/remove?photo_id=gopherPhoto",
		httpmock.NewStringResponder(202, `Bad ass Bug flow`))

	unsplash := setup()
	assert := assert.New(T)
	resp, err := unsplash.Collections.RemovePhoto(4242, "gopherPhoto")
	assert.NotNil(err)
	assert.Nil(resp)
	log.Println(err)

	httpmock.RegisterResponder("DELETE", getEndpoint(base)+getEndpoint(collections)+"/4242/remove?photo_id=gopherPhoto",
		nil)

	resp, err = unsplash.Collections.RemovePhoto(4242, "gopherPhoto")
	assert.Nil(resp)
	assert.NotNil(err)
	log.Println(err)
}
