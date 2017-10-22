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

import "bytes"

// ProfileImage contains URLs to profile image of a user
type ProfileImage struct {
	Small  *URL `json:"small"`
	Medium *URL `json:"medium"`
	Large  *URL `json:"large"`
	Custom *URL `json:"custom"`
}

// UserLinks contains URLs to
type UserLinks struct {
	Followers *URL `json:"followers"`
	Following *URL `json:"following"`
	Self      *URL `json:"self"`
	HTML      *URL `json:"html"`
	Photos    *URL `json:"photos"`
	Likes     *URL `json:"likes"`
	Portfolio *URL `json:"portfolio"`
}

// UserBadge contains information about badge for a user
type UserBadge struct {
	Title   *URL    `json:"title,omitempty"`
	Primary *bool   `json:"primary,omitempty"`
	Slug    *string `json:"slug,omitempty"`
	Link    *URL    `json:"link,omitempty"`
}

// User represents a Unsplash.com user
type User struct {
	UID                 *string       `json:"uid"`
	ID                  *string       `json:"id"`
	Username            *string       `json:"username"`
	Name                *string       `json:"name"`
	FirstName           *string       `json:"first_name"`
	CompletedOnboarding *bool         `json:"completed_onboarding"`
	LastName            *string       `json:"last_name,omitempty"`
	PortfolioURL        *URL          `json:"portfolio_url"`
	Bio                 *string       `json:"bio"`
	Location            *string       `json:"location"`
	TotalLikes          *int          `json:"total_likes"`
	TotalPhotos         *int          `json:"total_photos"`
	TotalCollections    *int          `json:"total_collections"`
	FollowedByUser      *bool         `json:"followed_by_user"`
	NumericID           *int          `json:"numeric_id"`
	FollowersCount      *int          `json:"followers_count"`
	FollowingCount      *int          `json:"following_count"`
	Downloads           *int          `json:"downloads"`
	ProfileImage        *ProfileImage `json:"profile_image"`
	Badge               *UserBadge    `json:"badge"`
	Links               *UserLinks    `json:"links,omitempty"`
	Photos              *[]Photo      `json:"photos"`
}

func (u *User) String() string {
	var buf bytes.Buffer
	if u.ID == nil {
		return "User is not valid"
	}
	buf.WriteString("User: Name[")
	buf.WriteString(*u.Name)
	buf.WriteString("], ID[")
	buf.WriteString(*u.ID)
	buf.WriteString("]")
	return buf.String()
}

// UserUpdateInfo is used to update private data of a user
type UserUpdateInfo struct {
	Username          string `url:"username,omitempty"`
	FirstName         string `url:"first_name,omitempty"`
	LastName          string `url:"last_name,omitempty"`
	Bio               string `url:"bio,omitempty"`
	Email             string `url:"email,omitempty"`
	PortfolioURL      string `url:"url,omitempty"`
	Location          string `url:"location,omitempty"`
	InstagramUsername string `url:"instagram_username,omitempty"`
}

func (u *UserUpdateInfo) String() string {
	var buf bytes.Buffer
	buf.WriteString("UserUpdateInfo:")

	buf.WriteString("Username[")
	buf.WriteString(u.Username)
	buf.WriteString("]")

	buf.WriteString("FirstName[")
	buf.WriteString(u.FirstName)
	buf.WriteString("]")

	buf.WriteString("LastName[")
	buf.WriteString(u.LastName)
	buf.WriteString("]")

	buf.WriteString("Bio[")
	buf.WriteString(u.Bio)
	buf.WriteString("]")

	buf.WriteString("Email[")
	buf.WriteString(u.Email)
	buf.WriteString("]")

	buf.WriteString("PortfolioURL[")
	buf.WriteString(u.PortfolioURL)
	buf.WriteString("]")

	buf.WriteString("Location[")
	buf.WriteString(u.Location)
	buf.WriteString("]")

	buf.WriteString("InstagramUsername[")
	buf.WriteString(u.InstagramUsername)
	buf.WriteString("]")

	return buf.String()
}

// UserStatistics represents statistics like downloads, views and likes of an unsplash user
type UserStatistics struct {
	Username  string `json:"username"`
	Downloads struct {
		Total      int `json:"total"`
		Historical struct {
			Change     int    `json:"change"`
			Average    int    `json:"average"`
			Resolution string `json:"resolution"`
			Quantity   int    `json:"quantity"`
			Values     []struct {
				Date  string `json:"date"`
				Value int    `json:"value"`
			} `json:"values"`
		} `json:"historical"`
	} `json:"downloads"`
	Views struct {
		Total      int `json:"total"`
		Historical struct {
			Change     int    `json:"change"`
			Average    int    `json:"average"`
			Resolution string `json:"resolution"`
			Quantity   int    `json:"quantity"`
			Values     []struct {
				Date  string `json:"date"`
				Value int    `json:"value"`
			} `json:"values"`
		} `json:"historical"`
	} `json:"views"`
	Likes struct {
		Total      int `json:"total"`
		Historical struct {
			Change     int    `json:"change"`
			Average    int    `json:"average"`
			Resolution string `json:"resolution"`
			Quantity   int    `json:"quantity"`
			Values     []struct {
				Date  string `json:"date"`
				Value int    `json:"value"`
			} `json:"values"`
		} `json:"historical"`
	} `json:"likes"`
}
