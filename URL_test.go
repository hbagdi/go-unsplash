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
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	correctData = []string{
		"{\"URL\":\"https://unsplash.com/@hbagdi/likes\"}",
		"{\"URL\":\"https://unsplash.com/documentation#get-the-users-profile\"}",
		"{\"URL\":\"https://en.wikipedia.org/wiki/42_(number)\"}",
	}
	badJSON = "{\"Anything that can go wrong, will go wrong.\"}"
	badURL  = "{\"URL\":\"httpps://en.shttps:////\"}"
)

type URLWrapper struct {
	OURL *URL `json:"URL,omitempty"`
}

func TestURL(T *testing.T) {
	assert := assert.New(T)
	log.SetOutput(ioutil.Discard)

	for _, value := range correctData {
		bytes := []byte(value)
		var url URLWrapper
		err := json.Unmarshal(bytes, &url)
		assert.Nil(err)
		log.Println("Unmarshalled string: ", url.OURL.String())
		marshalBytes, err := json.Marshal(url)
		assert.Nil(err)
		log.Println("Marshalled string: ", string(marshalBytes))
		assert.Equal(string(marshalBytes), value)
	}

	log.SetOutput(os.Stdout)
	var url URLWrapper
	bytes := []byte(badJSON)
	err := json.Unmarshal(bytes, &url)
	assert.NotNil(err)
	log.Println(err)
	assert.Nil(url.OURL)

	var url2 URL
	err = url2.UnmarshalJSON([]byte(badURL))
	assert.NotNil(err)

	var url3 URL
	err = url3.UnmarshalJSON([]byte(badJSON))
	assert.NotNil(err)

}
