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

//These are permission scopes for the Unsplash API for OAuth.
const (
	//Public is default; gives access to read public data.
	Public = "public"
	//ReadUser gives access to read user’s private data.
	ReadUser = "read_user"
	//WriteUser gives access to update the user’s profile.
	WriteUser = "write_user"
	//ReadPhotos gives acess to read private data from the user’s photos.
	ReadPhotos = "read_photos"
	//WritePhotos gives access to update photos on the user’s behalf.
	WritePhotos = "write_photos"
	//WriteLikes gives access to  like/unlike a photo on the user’s behalf.
	WriteLikes = "write_likes"
	//WriteFollowers gives access to follow or unfollow a user on the user’s behalf.
	WriteFollowers = "write_followers"
	//ReadCollections gives access to view a user’s private collections.
	ReadCollections = "read_collections"
	//WriteCollections gives access to create and update a users's collections.
	WriteCollections = "write_collections"
)
