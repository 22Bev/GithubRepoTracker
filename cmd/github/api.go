package github

import (
	"fmt"

	"encoding/json"
	
	"net/http"
)

func fetchRepo(ownerRepo string) (*Repo, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s", ownerRepo)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var repo Repo
	if err := json.NewDecoder(resp.Body).Decode(&repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

func fetchCommits(ownerRepo string) ([]Commit, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/commits?per_page=5", ownerRepo)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var commits []Commit
	if err := json.NewDecoder(resp.Body).Decode(&commits); err != nil {
		return nil, err
	}
	return commits, nil
}
