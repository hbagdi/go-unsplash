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
	"errors"
	"fmt"
)

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

// Related returns a list of collections related to collections with id.
func (cs *CollectionsService) Related(id string, opt *ListOpt) (*[]Collection, *Response, error) {
	if "" == id {
		return nil, nil, &IllegalArgumentError{ErrString: "Collection ID cannot be nil"}
	}
	s := (service)(*cs)
	endpoint := fmt.Sprintf("%v/%v/%v", getEndpoint(collections), id, "related")
	return s.getCollections(opt, endpoint)
}

// Collection returns a collection with id.
func (cs *CollectionsService) Collection(id string) (*Collection, *Response, error) {
	if "" == id {
		return nil, nil, &IllegalArgumentError{ErrString: "Collection ID cannot be nil"}
	}
	endpoint := fmt.Sprintf("%v/%v", getEndpoint(collections), id)
	req, err := newRequest(GET, endpoint, nil, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err := cs.client.do(req)
	if err != nil {
		return nil, nil, err
	}
	var collection Collection
	err = json.Unmarshal(*resp.body, &collection)
	if err != nil {
		return nil, nil, err
	}
	return &collection, resp, nil
}

//CollectionOpt shows various available optional parameters available
//during creatioin of collection
type CollectionOpt struct {
	Title       *string `url:"title,omitempty"`
	Description *string `url:"description,omitempty"`
	Private     *bool   `url:"private,omitempty"`
}

//Create creates a new collection on the authenticated  user's profile.
func (cs *CollectionsService) Create(opt *CollectionOpt) (*Collection, *Response, error) {
	if nil == opt {
		return nil, nil, &IllegalArgumentError{ErrString: "Opt cannot be nil"}
	}
	if *opt.Title == "" {
		return nil, nil, &IllegalArgumentError{ErrString: "Need to provide a title for the new collection."}
	}
	req, err := newRequest(POST, getEndpoint(collections), opt, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err := cs.client.do(req)
	if err != nil {
		return nil, nil, err
	}
	var collection Collection
	err = json.Unmarshal(*resp.body, &collection)
	if err != nil {
		return nil, nil, err
	}
	if resp.httpResponse.StatusCode != 201 {
		return nil, nil, errors.New("failed to create the collection")
	}
	return &collection, resp, nil
}

//Update updates an existing collection on the authenticated  user's profile.
func (cs *CollectionsService) Update(collectionID int, opt *CollectionOpt) (*Collection, *Response, error) {
	if nil == opt {
		return nil, nil, &IllegalArgumentError{ErrString: "Opt cannot be nil"}
	}
	if collectionID == 0 {
		return nil, nil, &IllegalArgumentError{ErrString: "collectionID cannot be nil."}
	}
	endpoint := fmt.Sprintf("%v/%v", getEndpoint(collections), collectionID)
	req, err := newRequest(PUT, endpoint, opt, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err := cs.client.do(req)
	if err != nil {
		return nil, nil, err
	}
	var collection Collection
	err = json.Unmarshal(*resp.body, &collection)
	if err != nil {
		return nil, nil, err
	}
	return &collection, resp, nil
}

//Delete deletes a collection on the authenticated user's profile.
func (cs *CollectionsService) Delete(collectionID int) (*Response, error) {
	if collectionID == 0 {
		return nil, &IllegalArgumentError{ErrString: "CollectionID cannot be empty or zero."}
	}
	endpoint := fmt.Sprintf("%v/%v", getEndpoint(collections), collectionID)
	req, err := newRequest(DELETE, endpoint, nil, nil)
	if err != nil {
		return nil, err
	}
	resp, err := cs.client.do(req)
	if err != nil {
		return nil, err
	}
	if resp.httpResponse.StatusCode != 204 {
		return nil, errors.New("failed to delete the collection")
	}
	return resp, nil
}

type addPhoto struct {
	Photo string `url:"photo_id"`
}

//AddPhoto adds a photo to a collection owned by an authenticated user.
func (cs *CollectionsService) AddPhoto(collectionID int, photoID string) (*Response, error) {
	if collectionID == 0 {
		return nil, &IllegalArgumentError{ErrString: "CollectionID cannot be empty or zero."}
	}
	if photoID == "" {
		return nil, &IllegalArgumentError{ErrString: "PhotoID cannot be empty or zero."}
	}
	opt := &addPhoto{photoID}
	endpoint := fmt.Sprintf("%v/%v/add", getEndpoint(collections), collectionID)
	req, err := newRequest(POST, endpoint, opt, nil)
	if err != nil {
		return nil, err
	}
	resp, err := cs.client.do(req)
	if err != nil {
		return nil, err
	}
	if resp.httpResponse.StatusCode != 201 {
		return nil, errors.New("failed to add photo to the collection")
	}
	return resp, nil
}

//RemovePhoto removes a photo from a collection owned by an authenticated user.
func (cs *CollectionsService) RemovePhoto(collectionID int, photoID string) (*Response, error) {
	if collectionID == 0 {
		return nil, &IllegalArgumentError{ErrString: "CollectionID cannot be empty or zero."}
	}
	if photoID == "" {
		return nil, &IllegalArgumentError{ErrString: "PhotoID cannot be empty or zero."}
	}
	opt := &addPhoto{photoID}
	endpoint := fmt.Sprintf("%v/%v/remove", getEndpoint(collections), collectionID)
	req, err := newRequest(DELETE, endpoint, opt, nil)
	if err != nil {
		return nil, err
	}
	resp, err := cs.client.do(req)
	if err != nil {
		return nil, err
	}
	if resp.httpResponse.StatusCode != 200 {
		return nil, errors.New("failed to remove photo from the collection")
	}
	return resp, nil
}
