package main

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
	bytes, err := ioutil.ReadFile("../auth.json")
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

//Photos.All()
func photosAll() {
	un := setup()
	fmt.Println("Done auth")
	opt := new(unsplash.ListOpt)
	opt.Page = 1
	opt.PerPage = 100

	if !opt.Valid() {
		fmt.Println("error with opt")
		return
	}
	count := 0
	for {
		photos, resp, err := un.Photos.All(opt)

		if err != nil {
			fmt.Println("error")
			return
		}
		//process photos
		for _, c := range *photos {
			fmt.Printf("%d : %d\n", count, *c.ID)
			count += 1
		}
		//go for next page
		if !resp.HasNextPage {
			return
		}
		opt.Page = resp.NextPage
	}
}

func main() {
	photosAll()
}
