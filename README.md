# Unsplash API client

[![GoDoc](https://godoc.org/github.com/hardikbagdi/go-unsplash?status.svg)](https://godoc.org/github.com/hardikbagdi/go-unsplash/unsplash)
[![Go Report Card](https://goreportcard.com/badge/github.com/hardikbagdi/go-unsplash)](https://goreportcard.com/report/github.com/hardikbagdi/go-unsplash)
[![Build Status](https://travis-ci.org/hardikbagdi/go-unsplash.svg?branch=master)](https://travis-ci.org/hardikbagdi/go-unsplash)
[![Coverage Status](https://coveralls.io/repos/github/hardikbagdi/go-unsplash/badge.svg?branch=master)](https://coveralls.io/github/hardikbagdi/go-unsplash?branch=master)

A wrapper for the [Unsplash API](https://unsplash.com/developers).

[Unsplash.com](https://unsplash.com) provides free
([CC0-licensed](https://unsplash.com/license)) high-resolution photos.
As many as you want without any pesky API rate limits.

## Documentation  
- [Installation](#installation)
- [Dependencies](#dependencies)
- [API terms and guidelines](#api-guidelines)
- [Registering your App](#registration)
- [Contact for help](#help)
- [Usage](#usage)
- [License](#license)

## Installation

```go
go get github.com/hardikbagdi/go-unsplash
```

## Dependencies
This library has a single dependency on Google's
[go-querystring](https://github.com/google/go-querystring/query).


## API Guidelines

- API is [open and free]
(https://community.unsplash.com/developersblog/the-unsplash-api-is-now-open-free).
- API [Terms & Conditions](https://unsplash.com/api-terms)
- API [Guidelines]
(https://community.unsplash.com/developersblog/unsplash-api-guidelines)
- [Hotlinking](https://unsplash.com/documentation#hotlinking)
the Unsplash image files is encouraged.


## Registration
[Sign up](https://unsplash.com/join) on Unsplash.com and register as a
[developer](https://unsplash.com/developers).  
You can then
[create a new application](https://unsplash.com/oauth/applications/new) and
use the AppID and Secret for authentication.

## Help
Please open an issue in this repository if you need help or want to report a bug.  
Mail at the e-mail address in the license if needed.

## Usage

- [Importing](#importing)
- [Authentication](#authentication)
- [Creating an instance](#creating-an-instance)
- [Error handling](#error-handling)
- [Response struct](#response-struct)
- [Pagination](#pagination) - paging through results
- [unsplash.Photos](#photos)
  - [Random](#random) - get random photo(s)
  - [All](#all-photos) - get all photos on unplash.com
  - [Curated](#curated-photos) - returns a list of curated photos
  - [Photo](#photo) -get details of a single photo
  - [Like](#like) - like a photo on the authenticated users' behalf
  - [Unlike](#unlike) - unlike a photo on the authenticated users' behalf
  - [Download link](#download-link) - get download link of a photo
  - [Stats](#stats) - statistics of a photo
- [unsplash.Collections](#collections)
  - [All](#all-collections) - return all collections on unsplash.com
  - [Collection](#collection) - get details about a single existing collection
  - [Curated](#curated-collections) -  get a list of curated collections
  - [Featured](#featured-collections) - get a list of featured collections
  - [Related](#related-collections) - get a list of related collections of a particular collection
  - [Create](#create-collection) - create a collection for the authenticated user
  - [Delete](#delete-collection) - delete a collection
  - [Update](#update-collection) - update a collection's description
  - [Add Photo](#add-photo) - add a photo to a collection
  - [Remove Photo](#remove-photo) - remove a photo from a collection
- [unsplash.Users](#users)
  - [User](#user) -  get details about a single user
  - [Portfolio](#portfolio) - get link to a user's portfolio
  - [Liked Photos](#liked-photos) - get list of photos a user has liked
  - [Photos](#user-photos) - get list of photos a user has uploaded to unsplash.com
  - [Collections](#user-collections) - collections of a user
- [unsplash.Search](#search)
  - [Photos](#search-photos) - search photos
  - [Collections](#search-collections) - search collections
  - [Users](#search-users) - search users


### Importing

Once you've installed the library using [`go get`](#installation),
import it to your go project:
```go
import "github.com/hardikbagdi/go-unsplash/unsplash"
```

### Authentication
Authentication is not handled by directly by go-unsplash.  
Instead, pass an [`http.Client`](https://godoc.org/net/http#Client)
that can handle authentication for you.  
You can use libraries such as [oauth2](https://godoc.org/golang.org/x/oauth2).  
Please note that all calls will include the OAuth token and hence,
[`http.Client`](https://godoc.org/net/http#Client)
should not be shared between users.

Note that if you're just using actions that require the public permission scope,
only the AppID is required.

### Creating an instance
An instance of unsplash can be created using `New()`.  
The http.Client supplied will be used to make requests to the API.

```go
ts := oauth2.StaticTokenSource(
  &oauth2.Token{AccessToken: "Your-access-token"},
)
client := oauth2.NewClient(oauth2.NoContext, ts)
//use the http.Client to instantiate unsplash
unsplash := unsplash.New(client)  
// requests can be now made to the API
randomPhoto, _ , err := unsplash.RandomPhoto(nil)
```

### Error handling
All API calls return an `error` as second or third return object.
All successful calls will return nil in place of this return.
Further, go-unsplash has errors defined as types for better error handling.  
```go
randomPhoto, _ , err := unsplash.RandomPhoto(nil)
if err != nil {
  //handle error
}
```

### Response struct
Most API methods return a
[`*Response`](https://godoc.org/github.com/hardikbagdi/go-unsplash/unsplash#Response)
along-with the result of the call.  
This struct contains paging and rate-limit information.

### Pagination
Pagination is currently supported by supplying a page
number in the
[`ListOpt`](https://godoc.org/github.com/hardikbagdi/go-unsplash/unsplash#ListOpt).
The `NextPage` field in Response can be used to  get the next page number.
```go
searchOpt := &SearchOpt{Query : "Batman"}
photos, resp, err := unsplash.Search.Photos(searchOpt)

if err != nil {
  return
}
// process photos
for _,photo := range *photos {
  fmt.Println(*photo.ID)
}
// get next
if !resp.HasNextPage {
  return
}
searchOpt.Page = resp.NextPage
photos, resp ,err = unsplash.Search.Photos(searchOpt)
//photos now has next page of the search result
```
### Photos
Unsplash.Photos is of type PhotosService.  
It provides various methods for querying the /photos endpoint of the API.

###### Random
You can get a single random photo or multiple depending upon opt.
If opt is nil, then a single random photo is returned.
Random photos satisfy all the parameters specified in `*RandomPhotoOpt`.
```go
photos, resp, err := unsplash.Photos.Random(nil)
assert.Nil(err)
assert.NotNil(photos)
assert.NotNil(resp)
assert.Equal(1, len(*photos))
var opt RandomPhotoOpt
opt.Count = 3
photos, resp, err = unsplash.Photos.Random(&opt)
assert.Nil(err)
assert.NotNil(photos)
assert.NotNil(resp)
assert.Equal(3, len(*photos))
```

###### All photos
Get all photos on unsplash.com.  
Obviously, this is a huge list and hence can be paginated through.
```go
opt := new(unsplash.ListOpt)
opt.Page = 1
opt.PerPage = 10

if !opt.Valid() {
	fmt.Println("error with opt")
	return
}
count := 0
for {
	photos, resp, err := un.Photos.All(opt)

	if err != nil {
		fmt.Println("error")
		return
	}
	//process photos
	for _, c := range *photos {
		fmt.Printf("%d : %d\n", count, *c.ID)
		count += 1
	}
	//go for next page
	if !resp.HasNextPage {
		return
	}
	opt.Page = resp.NextPage
}
```

###### Curated Photos
Get all curated photos on unsplash.com.  
Obviously, this is a huge list and hence can be paginated through.
```go
opt := new(unsplash.ListOpt)
opt.Page = 1
opt.PerPage = 10

if !opt.Valid() {
	fmt.Println("error with opt")
	return
}
count := 0
for {
	photos, resp, err := un.Photos.Curated(opt)

	if err != nil {
		fmt.Println("error")
		return
	}
	//process photos
	for _, c := range *photos {
		fmt.Printf("%d : %d\n", count, *c.ID)
		count += 1
	}
	//go for next page
	if !resp.HasNextPage {
		return
	}
	opt.Page = resp.NextPage
}
```

###### Photo
Get details of a specific photo.
```go
photo, resp, err := unsplash.Photos.Photo("9BoqXzEeQqM", nil)
assert.NotNil(photo)
assert.NotNil(resp)
assert.Nil(err)
fmt.Println(photo)
//photo is of type *Unsplash.Photo

// you can also specify a PhotoOpt to get a custom size or cropped photo
var opt PhotoOpt
opt.Height = 400
opt.Width = 600
photo, resp, err = unsplash.Photos.Photo("9BoqXzEeQqM", &opt)
assert.NotNil(photo)
assert.NotNil(resp)
assert.Nil(err)
log.Println(photo)
//photo.Urls.Custom will have the cropped photo
// See PhotoOpt for more details
```

###### Like
TODO

###### Unlike
TODO

###### Download Link
TODO

###### Stats
TODO


### Collections
TODO

###### All collections
TODO

###### Curated collections
TODO

###### Featured collections
TODO

###### Related collections
TODO

###### Collection
TODO

###### Create collection
TODO

###### Delete collection
TODO

###### Update collection
TODO

###### Add photo
TODO

###### Remove photo
TODO

### Users
TODO

###### User
TODO

###### Portfolio
TODO

###### Liked Photos
TODO

###### User photos
TODO

###### User collections
TODO

### Search
TODO

###### Search photos
TODO

###### Search collections
TODO

###### Search users
TODO


## License

 Copyright (c) 2017 Hardik Bagdi <hbagdi1@binghamton.edu>

 MIT License

 Permission is hereby granted, free of charge, to any person obtaining
 a copy of this software and associated documentation files (the
 "Software"), to deal in the Software without restriction, including
 without limitation the rights to use, copy, modify, merge, publish,
 distribute, sublicense, and/or sell copies of the Software, and to
 permit persons to whom the Software is furnished to do so, subject to
 the following conditions:

 The above copyright notice and this permission notice shall be
 included in all copies or substantial portions of the Software.

 THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
 NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
 LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
 OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
 WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
