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
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllCollections(T *testing.T) {
	assert := assert.New(T)
	log.SetOutput(os.Stdout)
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
	_, ok := err.(*InvalidListOpt)
	assert.Equal(true, ok)

}
