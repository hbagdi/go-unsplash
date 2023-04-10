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

type method string

const (
	// GET is HTTP GET request
	GET method = "GET"
	// POST is HTTP POST request
	POST method = "POST"
	// PUT is HTTP PUT request
	PUT method = "PUT"
	// DELETE is HTTP DELETE request
	DELETE method = "DELETE"
)

const (
	currentUserEndpoint       = "me"
	globalStatsEndpoint       = "stats/total"
	monthStatsEndpoint        = "stats/month"
	usersEndpoint             = "users"
	photosEndpoint            = "photos"
	collectionsEndpoint       = "collections"
	searchEndpoint            = "search"
	searchUserEndpoint        = searchEndpoint + "/" + usersEndpoint
	searchPhotosEndpoint      = searchEndpoint + "/" + photosEndpoint
	searchCollectionsEndpoint = searchEndpoint + "/" + collectionsEndpoint
)

var apiBaseURL = "https://api.unsplash.com/"

type endpoint int

const (
	base endpoint = iota
	currentUser
	globalStats
	monthStats
	users
	photos
	collections
	searchUsers
	searchPhotos
	searchCollections
)

var mapURL map[endpoint]string

func init() {
	mapURL = make(map[endpoint]string)
	mapURL[base] = apiBaseURL
	mapURL[currentUser] = currentUserEndpoint
	mapURL[globalStats] = globalStatsEndpoint
	mapURL[monthStats] = monthStatsEndpoint
	mapURL[users] = usersEndpoint
	mapURL[photos] = photosEndpoint
	mapURL[collections] = collectionsEndpoint
	mapURL[searchUsers] = searchUserEndpoint
	mapURL[searchPhotos] = searchPhotosEndpoint
	mapURL[searchCollections] = searchCollectionsEndpoint
}

func getEndpoint(e endpoint) string {
	return mapURL[e]
}

func SetupBaseUrl(url string) {
	apiBaseURL = url
	mapURL[base] = apiBaseURL
}
