package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchRepoInfo(repo string) (*RepoInfo, error) {
	URL := fmt.Sprintf("https://api.github.com/repos/%s", repo)
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var info RepoInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, err
	}

	// Fetch latest commit
	commitsURL := fmt.Sprintf("https://api.github.com/repos/%s/commits", repo)
	commitsResp, err := http.Get(commitsURL)
	if err != nil {
		return nil, err
	}
	defer commitsResp.Body.Close()
	var commits []Commit
	if err := json.NewDecoder(commitsResp.Body).Decode(&commits); err != nil {
		return nil, err
	}

	// Optionally print latest commit info
	if len(commits) > 0 {
		fmt.Println("Latest Commit:")
		fmt.Println("SHA:", commits[0].SHA)
		fmt.Println("Message:", commits[0].Commit.Message)
		fmt.Println("Author:", commits[0].Commit.Author.Name)
		fmt.Println("Date:", commits[0].Commit.Author.Date)
		fmt.Println("URL:", commits[0].HTMLURL)
	}

	return &info, nil
}
