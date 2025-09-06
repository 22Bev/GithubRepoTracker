package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchRepoInfo(Repo string) (*RepoInfo, error) {
	URL := fmt.Sprintf("https://api.github.com/repos/%s", Repo)
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var info RepoInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, err
	}

	//fetch latest commit
	commitsURL := fmt.Sprintf("https://api.github.com/repos/%s/commits", Repo)
	commitsResp, err := http.Get(commitsURL)
	if err != nil {
		return nil, err
	}
	defer commitsResp.Body.Close()	
	var commits []Commit
	if err := json.NewDecoder(commitsResp.Body).Decode(&commits); err != nil {
		return nil, err
	}
}
