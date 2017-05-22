# Unsplash API client

[![GoDoc](https://godoc.org/github.com/hardikbagdi/go-unsplash?status.svg)](https://godoc.org/github.com/hardikbagdi/go-unsplash/unsplash)
[![Go Report Card](https://goreportcard.com/badge/github.com/hardikbagdi/go-unsplash)](https://goreportcard.com/report/github.com/hardikbagdi/go-unsplash)
[![Build Status](https://travis-ci.org/hardikbagdi/go-unsplash.svg?branch=master)](https://travis-ci.org/hardikbagdi/go-unsplash)
[![Coverage Status](https://coveralls.io/repos/github/hardikbagdi/go-unsplash/badge.svg?branch=master)](https://coveralls.io/github/hardikbagdi/go-unsplash?branch=master)

A wrapper for the [Unsplash API](https://unsplash.com/developers).

[Unsplash.com](https://unsplash.com) provides free ([ CC0-licensed](https://unsplash.com/license)) high-resolution photos.
As many as you want without any pesky API rate limits.

## Documentation
- [Installation](#installation)
- [Dependencies](#dependencies)
- [API terms and guidelines](#api-guidelines)
- [Registering your App](#registration)
- [Usage](#usage)
  - [Authentication](authentication)
- [License](#license)

## Installation

```go
go get github.com/hardikbagdi/go-unsplash
```

## Dependencies
This library has a single dependency on Google's [go-querystring](https://github.com/google/go-querystring/query).


## API Guidelines

- API is [open and free](https://community.unsplash.com/developersblog/the-unsplash-api-is-now-open-free).
- API [Terms & Conditions](https://unsplash.com/api-terms)
- API [Guidelines](https://community.unsplash.com/developersblog/unsplash-api-guidelines)
- [Hotlinking](https://unsplash.com/documentation#hotlinking) the
Unsplash image files is encouraged.


## Registration
[Sign up](https://unsplash.com/join) on Unsplash.com and register as a [developer](https://unsplash.com/developers).  
You can then
[create a new application](https://unsplash.com/oauth/applications/new) and
use the AppID and Secret for authentication.

## Usage
Once you've installed the library using [`go get`](#installation),
import it to your go project:
```go
import github.com/hardikbagdi/go-unsplash/unsplash
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

 TODO add more auth info and example.

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

#### Error types
 TODO

### Response object
Most API methods return a
[`*Response`](https://godoc.org/github.com/hardikbagdi/go-unsplash/unsplash#Response)
along-with the result of the call.  
This struct contains paging and rate-limit information.

### Paging
Paging is currently supported by supplying a page
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
// get next
if !resp.HasNextPage {
  return
}
searchOpt.Page := resp.NextPage
photos, resp ,err = unsplash.Search.Photos(searchOpt)
//photos now has next page of photos
```
### Photos
TODO

### Collections
TODO

### Users
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
