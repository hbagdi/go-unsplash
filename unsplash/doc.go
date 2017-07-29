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

/*
Package unsplash provides a RESTful go binding for https//:unsplash.com API.


Usage

Use the following import path:
    import "github.com/hbagdi/go-unsplash/unsplash"

Authentication

Authentication is not handled by directly by go-unsplash.
Instead, pass an http.Client
that can handle authentication for you.
You can use libraries such as https://godoc.org/golang.org/x/oauth2.
Please note that all calls will include the OAuth token and hence, http.Client
should not be shared between users.
Note that if you're just using actions that require the public permission scope,
only the AppID is required.


Creating an instance

An instance of unsplash can be created using New().
The http.Client supplied will be used to make requests to the API.

	ts := oauth2.StaticTokenSource(
	  &oauth2.Token{AccessToken: "Your-access-token"},
	)
	client := oauth2.NewClient(oauth2.NoContext, ts)
	//use the http.Client to instantiate unsplash
	unsplash := unsplash.New(client)
	// requests can be now made to the API
	randomPhoto, _ , err := unsplash.RandomPhoto(nil)

Error handling

All API calls return an error as second or third return object.
All successful calls will return nil in place of this return.
Further, go-unsplash has errors defined as types for better error handling.

	randomPhoto, _ , err := unsplash.RandomPhoto(nil)
	if err != nil {
	  //handle error
	}

Pagination

Pagination is supported by supplying a page
number in the ListOpt.
The NextPage field in Response can be used to  get the next page number.

	var allPhotos []*unsplash.Photo
	searchOpt := &unsplash.SearchOpt{Query: "Batman"}
	for {
		photos, resp, err := unsplash.Search.Photos(searchOpt)

		if err != nil {
			return
		}

		allPhotos = append(allPhotos, photos)

		if !resp.HasNextPage {
			break
		}
		searchOpt.Page = resp.NextPage
	}

*/
package unsplash
