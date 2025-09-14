package main

import (
	"fmt"
	"os"

	"github.com/22Bev/GithubRepoTracker/cmd/github"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go repo <owner>/<repo>")
		return
	}

	command := os.Args[1]
	repo := os.Args[2]

	switch command {
	case "repo":
		err := github.HandleRepo(repo)
		if err != nil {
			fmt.Println("Error:", err)
		}
	default:
		fmt.Println("Unknown command:", command)
	}
}
