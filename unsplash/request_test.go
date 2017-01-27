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

	"github.com/stretchr/testify/assert"
)

func TestRequest(T *testing.T) {
	assert := assert.New(T)
	log.SetOutput(ioutil.Discard)

	request, err := newRequest(GET, "", nil, nil)
	assert.NotNil(err)
	assert.Nil(request)
	iae, ok := err.(*IllegalArgumentError)
	assert.NotNil(iae)
	assert.Equal(true, ok)
}

func TestRequestRogue(T *testing.T) {
	assert := assert.New(T)
	log.SetOutput(ioutil.Discard)
	var ROGUE method
	ROGUE = "ROGUE"
	req, err := newRequest(ROGUE, "", nil, nil)
	assert.Nil(req)
	assert.NotNil(err)
}
