# Unsplash API client

[![GoDoc](https://godoc.org/github.com/hardikbagdi/go-unsplash?status.svg)](https://godoc.org/github.com/hbagdi/go-unsplash/unsplash)
[![Go Report Card](https://goreportcard.com/badge/github.com/hbagdi/go-unsplash)](https://goreportcard.com/report/github.com/hbagdi/go-unsplash)
[![Build Status](https://travis-ci.org/hbagdi/go-unsplash.svg?branch=master)](https://travis-ci.org/hbagdi/go-unsplash) 
[![Coverage Status](https://coveralls.io/repos/github/hbagdi/go-unsplash/badge.svg?branch=master)](https://coveralls.io/github/hbagdi/go-unsplash?branch=master)

A wrapper for the [Unsplash API](https://unsplash.com/developers).

[Unsplash.com](https://unsplash.com) provides free ([CC0-licensed](https://unsplash.com/license)) high-resolution photos. As many as you want without any pesky API rate limits.

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
go get github.com/hbagdi/go-unsplash/unsplash
```

## Dependencies

This library has a single dependency on Google's [go-querystring](https://github.com/google/go-querystring/query).

## API Guidelines

- API is [open and free] (<https://community.unsplash.com/developersblog/the-unsplash-api-is-now-open-free>).
- API [Terms & Conditions](https://unsplash.com/api-terms)
- API [Guidelines] (<https://community.unsplash.com/developersblog/unsplash-api-guidelines>)
- [Hotlinking](https://unsplash.com/documentation#hotlinking) the Unsplash image files is encouraged.

## Registration

[Sign up](https://unsplash.com/join) on Unsplash.com and register as a [developer](https://unsplash.com/developers).<br>
You can then [create a new application](https://unsplash.com/oauth/applications/new) and use the AppID and Secret for authentication.

## Help

Please open an issue in this repository if you need help or want to report a bug.<br>
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
  - [Curated](#curated-collections) - get a list of curated collections
  - [Featured](#featured-collections) - get a list of featured collections
  - [Related](#related-collections) - get a list of related collections of a particular collection
  - [Create](#create-collection) - create a collection for the authenticated user
  - [Delete](#delete-collection) - delete a collection
  - [Update](#update-collection) - update a collection's description
  - [Add Photo](#add-photo) - add a photo to a collection
  - [Remove Photo](#remove-photo) - remove a photo from a collection

- [unsplash.Users](#users)

  - [User](#user) - get details about a single user
  - [Portfolio](#portfolio) - get link to a user's portfolio
  - [Liked Photos](#liked-photos) - get list of photos a user has liked
  - [Photos](#user-photos) - get list of photos a user has uploaded to unsplash.com
  - [Collections](#user-collections) - collections of a user

- [unsplash.Search](#search)

  - [Photos](#search-photos) - search photos
  - [Collections](#search-collections) - search collections
  - [Users](#search-users) - search users

### Importing

Once you've installed the library using [`go get`](#installation), import it to your go project:

```go
import "github.com/hbagdi/go-unsplash/unsplash"
```

### Authentication

Authentication is not handled by directly by go-unsplash.<br>
Instead, pass an [`http.Client`](https://godoc.org/net/http#Client) that can handle authentication for you.<br>
You can use libraries such as [oauth2](https://godoc.org/golang.org/x/oauth2).<br>
Please note that all calls will include the OAuth token and hence, [`http.Client`](https://godoc.org/net/http#Client) should not be shared between users.

Note that if you're just using actions that require the public permission scope, only the AppID is required.

### Creating an instance

An instance of unsplash can be created using `New()`.<br>
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

All API calls return an `error` as second or third return object. All successful calls will return nil in place of this return. Further, go-unsplash has errors defined as types for better error handling.

```go
randomPhoto, _ , err := unsplash.RandomPhoto(nil)
if err != nil {
  //handle error
}
```

### Response struct

Most API methods return a [`*Response`](https://godoc.org/github.com/hbagdi/go-unsplash/unsplash#Response) along-with the result of the call.<br>
This struct contains paging and rate-limit information.

### Pagination

Pagination is currently supported by supplying a page number in the [`ListOpt`](https://godoc.org/github.com/hbagdi/go-unsplash/unsplash#ListOpt). The `NextPage` field in Response can be used to get the next page number.

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

Unsplash.Photos is of type PhotosService.<br>
It provides various methods for querying the /photos endpoint of the API.

#### Random

You can get a single random photo or multiple depending upon opt. If opt is nil, then a single random photo is returned. Random photos satisfy all the parameters specified in `*RandomPhotoOpt`.

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

#### All photos

Get all photos on unsplash.com.<br>
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

#### Curated Photos

Get all curated photos on unsplash.com.<br>
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

#### Photo

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

#### Like

```go
//Like a random photo
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
```

#### Unlike

Same way as Like except call `Unlike()`.

#### Download Link

Get download URL for a photo.

```go
url, resp, err := unsplash.Photos.DownloadLink("-HPhkZcJQNk")
assert.Nil(err)
assert.NotNil(url)
assert.NotNil(resp)
log.Println(url)
```

#### Stats

Statistics for a specific photo

```go
stats, resp, err := unsplash.Photos.Stats("-HPhkZcJQNk")
assert.Nil(err)
assert.NotNil(stats)
assert.NotNil(resp)
log.Println(stats)
```

### Collections

Various details about collection(s).

#### All collections

```go
collections, resp, err = unsplash.Collections.All(nil)
assert.Nil(err)
assert.NotNil(resp)
assert.NotNil(collections)
opt := new(ListOpt)
opt.Page = 2
opt.PerPage = 10
//get the second page
collections, resp, err = unsplash.Collections.All(opt)
assert.Nil(err)
assert.NotNil(resp)
assert.NotNil(collections)
```

#### Curated collections

```go
collections, resp, err := unsplash.Collections.Curated(nil)
assert.Nil(err)
assert.NotNil(resp)
assert.NotNil(collections)
```

#### Featured collections

Same as Curated, but use `Featured()` instead.

#### Related collections

Get collections related to a collection.

```go
collections, resp, err := unsplash.Collections.Related("296", nil)
assert.Nil(err)
assert.NotNil(resp)
assert.NotNil(collections)
//page through if necessary
```

#### Collection

Details about a specific collection

```go
collection, resp, err := unsplash.Collections.Collection("910")
assert.Nil(err)
assert.NotNil(resp)
assert.NotNil(collection)
```

#### Create collection

Create a collection on behalf of the authenticated user.

```go
var opt CollectionOpt
title := "Test42"
opt.Title = &title
collection, resp, err := unsplash.Collections.Create(&opt)
assert.Nil(err)
assert.NotNil(resp)
assert.NotNil(collection)
```

#### Delete collection

Let's delete the collection just created above.

```go
//get list of collections of a user
collections, resp, err := unsplash.Users.Collections("gopher", nil)
assert.Nil(err)
assert.NotNil(resp)
assert.NotNil(collections)
//take the first one
collection := (*collections)[0]
assert.NotNil(collection)
//delete it
resp, err = unsplash.Collections.Delete(*collection.ID)
assert.NotNil(resp)
assert.Nil(err)
```

#### Update collection

```go
//get a user's collection
collections, resp, err := unsplash.Users.Collections("gopher", nil)
assert.Nil(err)
assert.NotNil(resp)
assert.NotNil(collections)
// take the first one
collection := (*collections)[0]
assert.NotNil(collection)
log.Println(*collection.ID)
//random title
var opt CollectionOpt
title := "Test43" + strconv.Itoa(rand.Int())
opt.Title = &title
//update the title
col, resp, err := unsplash.Collections.Update(*collection.ID, &opt)
assert.Nil(err)
assert.NotNil(resp)
assert.NotNil(col)
```

#### Add photo

```go
photos, resp, err := unsplash.Photos.Random(nil)
photo := (*photos)[0]
//get a user's collection
collections, resp, err := unsplash.Users.Collections("gopher", nil)
assert.Nil(err)
collection := (*collections)[0]
//add the photo
resp, err = unsplash.Collections.AddPhoto(*collection.ID, *photo.ID)
assert.Nil(err)
```

#### Remove photo

```go
//remove a photo
_, _ = unsplash.Collections.RemovePhoto(*collection.ID, *photo.ID)
```

### Users

Details about an unsplash.com users.

#### User

Details about unsplash.com users.

```go
profileImageOpt := &ProfileImageOpt{Height: 120, Width: 400}
//or pass a nil as second arg
user, err := unsplash.Users.User("lukechesser", profileImageOpt)
assert.Nil(err)
assert.NotNil(user)

//OR, get the currently authenticated user
user, resp, err := unsplash.CurrentUser()
assert.Nil(user)
assert.Nil(resp)
assert.NotNil(err)
```

#### Portfolio

```go
url, err = unsplash.Users.Portfolio("gopher")
assert.Nil(err)
assert.NotNil(url)
assert.Equal(url.String(), "https://wikipedia.org/wiki/Gopher")
```

#### Liked Photos

```go
photos, resp, err := unsplash.Users.LikedPhotos("lukechesser", nil)
assert.Nil(err)
assert.NotNil(photos)
assert.NotNil(resp)
```

#### User photos

Get photos a users has uploaded on unsplash.com

```go
photos, resp, err := unsplash.Users.Photos("lukechesser", nil)
assert.Nil(err)
assert.NotNil(resp)
assert.NotNil(photos)
```

#### User collections

Get a list of collections created by the user.

```go
collections, resp, err := unsplash.Users.Collections("gopher", nil)
assert.Nil(err)
assert.NotNil(resp)
assert.NotNil(collections)
```

### Search

Search for photos, collections or users.

#### Search photos

```go
var opt SearchOpt
//an empty search will be erroneous
photos, resp, err := unsplash.Search.Photos(&opt)
assert.NotNil(err)
assert.Nil(resp)
assert.Nil(photos)
opt.Query = "Nature"
//Search for photos tageed "Nature"
photos, _, err = unsplash.Search.Photos(&opt)
log.Println(len(*photos.Results))
assert.NotNil(photos)
assert.Nil(err)
```

#### Search collections

```go
var opt SearchOpt
opt.Query = "Nature"
collections, _, err = unsplash.Search.Collections(&opt)
assert.NotNil(collections)
assert.Nil(err)
log.Println(len(*collections.Results))
```

#### Search users

```go
var opt SearchOpt
opt.Query = "Nature"
users, _, err = unsplash.Search.Users(&opt)
log.Println(len(*users.Results))
assert.NotNil(users)
assert.Nil(err)
```

## License

Copyright (c) 2017 Hardik Bagdi [hbagdi1@binghamton.edu](mailto:hbagdi1@binghamton.edu)

MIT License

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
