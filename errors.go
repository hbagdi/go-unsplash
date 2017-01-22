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

//The following are implementing error interface

// IllegalArgumentError occurs when the argument to a function are
// messed up
type IllegalArgumentError struct {
	ErrString string
}

func (e IllegalArgumentError) Error() string {
	return e.ErrString
}

// JSONUnmarshallingError occurs due to a unmarshalling error
type JSONUnmarshallingError struct {
	ErrString string
}

func (e JSONUnmarshallingError) Error() string {
	return e.ErrString
}

// AuthorizationError occurs due to a unmarshalling error
type AuthorizationError struct {
	ErrString string
}

func (e AuthorizationError) Error() string {
	return e.ErrString
}

// NotFoundError occurs due to a unmarshalling error
type NotFoundError struct {
	ErrString string
}

func (e NotFoundError) Error() string {
	return e.ErrString
}

// InvalidPhotoOpt occurs due to a unmarshalling error
type InvalidPhotoOpt struct {
	ErrString string
}

func (e InvalidPhotoOpt) Error() string {
	return e.ErrString
}
