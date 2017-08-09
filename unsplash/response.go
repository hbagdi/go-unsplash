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
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Response has pagination information whenever applicable
type Response struct {
	httpResponse                            *http.Response
	HasNextPage                             bool
	body                                    *[]byte
	FirstPage, LastPage, NextPage, PrevPage int
	err                                     error
	RateLimit                               int
	RateLimitRemaining                      int
}

func (r *Response) checkForErrors() error {
	switch r.httpResponse.StatusCode {
	case 200, 201, 202, 204, 205:
		return nil
	case 401:
		return &AuthorizationError{ErrString: errStringHelper(r.httpResponse.StatusCode, "Unauthorized request", r.body)}
	case 403:
		if r.RateLimitRemaining == 0 {
			return &RateLimitError{ErrString: errStringHelper(r.httpResponse.StatusCode, "Rate limit exhausted", r.body)}
		}

		return &AuthorizationError{ErrString: errStringHelper(r.httpResponse.StatusCode, "Access forbidden request", r.body)}

	case 404:
		return &NotFoundError{ErrString: errStringHelper(r.httpResponse.StatusCode, "The cat got tired of the Laser", r.body)}
	default:
		return errors.New(errStringHelper(r.httpResponse.StatusCode, "API returned an error", r.body))

	}
}
func errStringHelper(statusCode int, msg string, errBody *[]byte) string {
	var buf bytes.Buffer
	//XXX Writes can fail?
	buf.WriteString(strconv.Itoa(statusCode))
	buf.WriteString(": ")
	buf.WriteString(msg)
	buf.WriteString(", Body: ")
	buf.Write(*errBody)
	return buf.String()
}

func newResponse(r *http.Response) (*Response, error) {
	if nil == r {
		return nil,
			&IllegalArgumentError{ErrString: "*http.Response cannot be null"}
	}
	resp := new(Response)
	resp.httpResponse = r
	//populate first
	resp.populatePagingInfo()
	resp.populateRateLimits()
	//read the response
	buf, err := ioutil.ReadAll(resp.httpResponse.Body)
	if err != nil {
		return nil, err
	}
	resp.body = &buf
	//now check for errors
	err = resp.checkForErrors()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *Response) populateRateLimits() {
	//fails silently
	maxLimit, ok := r.httpResponse.Header["X-Ratelimit-Limit"]
	if ok && len(maxLimit) == 1 {
		r.RateLimit, _ = strconv.Atoi(maxLimit[0])
	}
	rateRemaining, ok := r.httpResponse.Header["X-Ratelimit-Remaining"]
	if ok && len(rateRemaining) == 1 {
		r.RateLimitRemaining, _ = strconv.Atoi(rateRemaining[0])
	}
}

func (r *Response) populatePagingInfo() {
	//fails silently
	rawLinks, ok := r.httpResponse.Header["Link"]
	if !ok || 0 == len(rawLinks) {
		return
	}

	links := strings.Split(rawLinks[0], ",")

	for _, link := range links {
		parts := strings.Split(link, ";")
		if !strings.Contains(parts[0], "page") && !strings.Contains(parts[1], "rel=") {
			continue
		}
		href := parts[0]
		//strip out '<' and '>'
		href = href[strings.Index(href, "<")+1 : strings.Index(href, ">")]
		url, err := url.Parse(href)
		if err != nil {
			continue
		}
		pageString := url.Query().Get("page")
		pageNumber, err := strconv.Atoi(string(pageString))
		if err != nil {
			continue
		}

		switch strings.TrimSpace(parts[1]) {
		case `rel="first"`:
			r.FirstPage = pageNumber
		case `rel="last"`:
			r.LastPage = pageNumber
		case `rel="next"`:
			r.NextPage = pageNumber
			r.HasNextPage = true
		case `rel="prev"`:
			r.PrevPage = pageNumber
		default:
			continue
		}
	}
}
