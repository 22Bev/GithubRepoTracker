package github

import "fmt"

func HandleRepo(ownerRepo string) error {
	repo, err := fetchRepo(ownerRepo)
	if err != nil {
		return err
	}

	commits, err := fetchCommits(ownerRepo)
	if err != nil {
		return err
	}

	fmt.Println("Repository:", repo.FullName)
	fmt.Println("Description:", repo.Description)
	fmt.Println("Stars:", repo.Stars)
	fmt.Println("Forks:", repo.Forks)

	if len(commits) > 0 {
		latest := commits[0]
		fmt.Println("Latest Commit:", latest.Commit.Message)
		fmt.Printf("By %s on %s\n", latest.Commit.Author.Name, latest.Commit.Author.Date)
	}

	return nil
}
