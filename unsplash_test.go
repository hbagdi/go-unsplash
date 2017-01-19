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

func getAppAuth() *AuthConfig {
	var config AuthConfig
	appID, ok := os.LookupEnv("unsplash_appID")
	if !ok {
		log.Println("unsplash_appID env varible not set. Stopping tests.")
		os.Exit(1)
	}
	config.AppID = appID
	return &config
}

func getUserAuth() *AuthConfig {
	config := getAppAuth()
	secret, ok := os.LookupEnv("unsplash_secret")
	if !ok {
		log.Println("unsplash_secret env varible not set. Stopping tests.")
		os.Exit(1)
	}
	config.Secret = secret
	token, ok := os.LookupEnv("unsplash_usertoken")
	if !ok {
		log.Println("unsplash_usertoken env varible not set. Stopping tests.")
		os.Exit(1)
	}
	config.AuthToken = token
	return config
}

func TestUnsplash(T *testing.T) {
	assert := assert.New(T)
	//test setup
	t := 2
	assert.Equal(2, t)

	unsplash, err := New(nil)
	assert.Nil(unsplash)
	assert.NotNil(err)
	_, ok := err.(*InvalidAuthCredentialsError)
	assert.NotNil(err.Error())
	assert.Equal(true, ok)

	var config AuthConfig
	unsplash, err = New(&config)
	assert.Nil(unsplash)
	assert.NotNil(err)
	_, ok = err.(*InvalidAuthCredentialsError)
	assert.NotNil(err.Error())
	assert.Equal(true, ok)

	c := getAppAuth()
	unsplash, err = New(c)
	assert.Nil(err)
	assert.NotNil(unsplash)
	unsplash.List()
}
