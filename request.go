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
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/google/go-querystring/query"
)

type request struct {
	Request *http.Request
}

func newRequest(m method, e string, qs interface{}, body interface{}) (*request, error) {
	if e == "" {
		return nil, &IllegalArgumentError{ErrString: "Endpoint can't be null."}
	}
	//body to be sent in JSON
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	//Create a new request

	httpRequest, err := http.NewRequest(string(m), getEndpoint(base)+e, bytes.NewBuffer(buf))

	if err != nil {
		return nil, err
	}
	//Add query string if any
	if qs != nil {
		values, err := query.Values(qs)
		if err != nil {
			return nil, err
		}
		httpRequest.URL.RawQuery = values.Encode()
	}
	req := new(request)
	req.Request = httpRequest
	req.Request.Header.Add("Content-Type", "application/json")
	return req, nil
}
