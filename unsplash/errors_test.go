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

func TestErrors(T *testing.T) {
	assert := assert.New(T)
	log.SetOutput(ioutil.Discard)
	var err error

	err = &IllegalArgumentError{ErrString: "I'm Joker."}
	iae2, ok := err.(*IllegalArgumentError)
	assert.NotNil(iae2)
	log.Println(iae2)
	assert.Equal(true, ok)

	err = &AuthorizationError{ErrString: "I'm Batman."}
	ae, ok := err.(*AuthorizationError)
	assert.NotNil(ae)
	log.Println(ae)
	assert.Equal(true, ok)

	err = &JSONUnmarshallingError{ErrString: "I'm Batman."}
	jue, ok := err.(*JSONUnmarshallingError)
	assert.NotNil(jue)
	log.Println(jue)
	assert.Equal(true, ok)

	err = &NotFoundError{ErrString: "I'll find you and I'll kill you."}
	nfe, ok := err.(*NotFoundError)
	assert.NotNil(nfe)
	log.Println(nfe)
	assert.Equal(true, ok)

	err = &InvalidPhotoOpt{ErrString: "."}
	ipo, ok := err.(*InvalidPhotoOpt)
	assert.NotNil(ipo)
	log.Println(ipo)
	assert.Equal(true, ok)

	err = &InvalidListOpt{ErrString: "."}
	ilo, ok := err.(*InvalidListOpt)
	assert.NotNil(ilo)
	log.Println(ilo)
	assert.Equal(true, ok)
}
