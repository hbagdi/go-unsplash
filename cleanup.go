package testing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/hardikbagdi/go-unsplash/unsplash"

	"golang.org/x/oauth2"
)

// AuthConfig is a temp struct for storing credentials
type AuthConfig struct {
	AppID, Secret, AuthToken string
}

func authFromFile() *AuthConfig {
	bytes, err := ioutil.ReadFile("auth.json")
	if err != nil {
		return nil
	}
	var config AuthConfig
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil
	}
	return &config
}

func setup() *unsplash.Unsplash {
	var c *AuthConfig

	c = authFromFile()
	if c == nil {
		log.Println("Couldn't read auth token. Stopping tests.")
		os.Exit(1)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.AuthToken},
	)
	client := oauth2.NewClient(oauth2.NoContext, ts)
	return unsplash.New(client)
}
func main() {
	un := setup()
	fmt.Println("Done auth")
	opt := new(unsplash.ListOpt)
	opt.PerPage = 100
	opt.Page = 1

	//delete all collections of @gopher
	for {
		collections, resp, err := un.Users.Collections("gopher", opt)

		if err != nil {
			fmt.Print("error")
			return
		}
		for _, c := range *collections {
			un.Collections.Delete(*c.ID)
		}
		if !resp.HasNextPage {
			break
		}
		opt.Page = resp.NextPage
	}

	opt.PerPage = 100
	opt.Page = 1

	//unlike all photos for @gopher
	for {
		photos, resp, err := un.Users.LikedPhotos("gopher", opt)

		if err != nil {
			fmt.Print("error")
			return
		}
		for _, c := range *photos {
			un.Photos.Unlike(*c.ID)
		}
		if !resp.HasNextPage {
			break
		}
		opt.Page = resp.NextPage
	}
}
