package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	login := "beckitrue"
	getGithubUserInfo(login)

	name, numRepos, err := getGithubUserInfo(login)
	if err == nil {
		fmt.Printf("name: %s, numRepos: %d\n", name, numRepos)
	}
}

type Reply struct {
	Name         string `json:"name"`
	Public_Repos int    `json:"public_repos"`
}

// function to get github user info
func getGithubUserInfo(username string) (string, int, error) {
	url := "https://api.github.com/users/" + url.PathEscape(username)
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalf("error: %s", err)
		return "", 0, err
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error: %s", resp.Status)
		return "", 0, fmt.Errorf("error: %s", resp.Status)
	}

	defer resp.Body.Close()

	var r struct { // anonymous struct - don't need to define a type
		Name     string
		NumRepos int `json:"public_repos"`
	}

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		log.Fatalf("error: %s", err)
		return "", 0, err
	}

	return r.Name, r.NumRepos, nil
}
