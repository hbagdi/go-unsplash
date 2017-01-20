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
	"net/http"
)

type response struct {
	Response                                *http.Response
	MorePage                                bool
	body                                    *[]byte
	FirstPage, PrevPage, NextPage, LastPage URL
	err                                     error
}

func (r *response) Errored() error {
	return r.err
}
func (r *response) CheckForErrors() error {
	return nil
}

func newResponse(r *http.Response) (*response, error) {
	if nil == r {
		return nil,
			&IllegalArgumentError{ErrString: "*http.Response cannot be null"}
	}
	resp := new(response)
	resp.Response = r

	defer r.Body.Close()
	err := resp.CheckForErrors()
	if err != nil {
		return nil, err
	}
	buf, err := ioutil.ReadAll(r.Body)
	resp.body = &buf
	if err != nil {
		return nil, err
	}
	return resp, nil
}
