package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <Owner>/<Repo>")
	}
	command := os.Args[1]
	repo := os.Args[2]
	switch command {
	case "Repo":
		info, err := GetRepoInfo(repo)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		PrintRepoInfo(info)

	default:
		fmt.Println("Unknown command:", command)
	}
}
