package main

// RepoInfo represents basic information about a GitHub repository.
type RepoInfo struct {
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	HTMLURL     string `json:"html_url"`
	Stargazers  int    `json:"stargazers_count"`
	Forks       int    `json:"forks_count"`
	OpenIssues  int    `json:"open_issues_count"`
	Owner       Owner  `json:"owner"`
}

// Owner represents the owner of the repository.
type Owner struct {
	Login   string `json:"login"`
	HTMLURL string `json:"html_url"`
}

// Commit represents a commit in the repository.
type Commit struct {
	SHA     string      `json:"sha"`
	Commit  CommitInner `json:"commit"`
	HTMLURL string      `json:"html_url"`
}

// CommitInner holds commit details.
type CommitInner struct {
	Message string       `json:"message"`
	Author  CommitAuthor `json:"author"`
}

// CommitAuthor holds author details for a commit.
type CommitAuthor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Date  string `json:"date"`
}
