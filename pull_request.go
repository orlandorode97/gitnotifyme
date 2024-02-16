package main

type PullRequest struct {
	URL      string `json:"url"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	User     User   `json:"user"`
	PRNumber int    `json:"number"`
}

type Repository struct {
	FullName string `json:"full_name"`
}
