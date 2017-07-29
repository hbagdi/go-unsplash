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
	"strconv"
)

//These constants should be used for OrderBy searches/results.
const (
	Latest  = "latest"
	Oldest  = "oldest"
	Popular = "popular"
)

var orders = []string{"latest", "oldest", "popular"}

// ListOpt should be used for pagination over results
type ListOpt struct {
	Page    int    `url:"page"`
	PerPage int    `url:"per_page"`
	OrderBy string `url:"order_by"` //TODO doc which endpoints obey this
}

var defaultListOpt = &ListOpt{
	Page:    1,
	PerPage: 10,
	OrderBy: Popular,
}

// Valid validates the values in a ListOpt
func (opt *ListOpt) Valid() bool {
	if opt.Page < 0 {
		return false
	}
	if opt.Page == 0 {
		opt.Page = 1
	}
	if opt.PerPage < 0 {
		return false
	}
	if opt.PerPage == 0 {
		opt.PerPage = 10
	}
	if opt.OrderBy == "" {
		opt.OrderBy = Latest
	}
	for _, val := range orders {
		if val == opt.OrderBy {
			return true
		}
	}
	return false
}

func (opt *ListOpt) String() string {
	var buf bytes.Buffer
	if !opt.Valid() {
		buf.WriteString("ListOpt is invalid")
	}
	buf.WriteString("ListOpt: Page[")
	buf.WriteString(strconv.Itoa(opt.Page))
	buf.WriteString("], PerPage[")
	buf.WriteString(strconv.Itoa(opt.PerPage))
	buf.WriteString("], OrderBy[")
	buf.WriteString(opt.OrderBy)
	buf.WriteString("]")
	return buf.String()
}
