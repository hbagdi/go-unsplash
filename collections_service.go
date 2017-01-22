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

// CollectionsService interacts with /users endpoint
type CollectionsService service

// All returns a list of all collections on unsplash.
// Note that some fields in collection structs from this result will be missing.
// Use Collection() method to get all details of the Collection.
func (cs *CollectionsService) All(opt *ListOpt) (*[]Collection, *Response, error) {
	s := (service)(*cs)
	return s.getCollections(opt, getEndpoint(collections))
}

// Featured returns a list of featured collections on unsplash.
// Note that some fields in collection structs from this result will be missing.
// Use Collection() method to get all details of the Collection.
func (cs *CollectionsService) Featured(opt *ListOpt) (*[]Collection, *Response, error) {
	s := (service)(*cs)
	return s.getCollections(opt, getEndpoint(collections)+"/featured")
}

// Curated returns a list of curated collections on unsplash.
// Note that some fields in collection structs from this result will be missing.
// Use Collection() method to get all details of the Collection.
func (cs *CollectionsService) Curated(opt *ListOpt) (*[]Collection, *Response, error) {
	s := (service)(*cs)
	return s.getCollections(opt, getEndpoint(collections)+"/curated")
}
