package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go Repo <Owner>/<Repo>")
		return
	}
	command := os.Args[1]
	repo := os.Args[2]
	switch command {
	case "Repo":
		info, err := FetchRepoInfo(repo)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		PrintRepoInfo(info)
	default:
		fmt.Println("Unknown command:", command)
	}
}

// PrintRepoInfo prints basic repository information.
func PrintRepoInfo(info *RepoInfo) {
	fmt.Println("Repository Name:", info.Name)
	fmt.Println("Full Name:", info.FullName)
	fmt.Println("Description:", info.Description)
	fmt.Println("URL:", info.HTMLURL)
	fmt.Println("Stars:", info.Stargazers)
	fmt.Println("Forks:", info.Forks)
	fmt.Println("Open Issues:", info.OpenIssues)
	fmt.Println("Owner:", info.Owner.Login)
}
