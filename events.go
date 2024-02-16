package main

type action string

const (
	Closed               action = "closed"
	Reopened             action = "reopened"
	Submitted            action = "submitted"
	Assigned             action = "assigned"
	ReviewRequested      action = "review_requested"
	ReviewRequestRemoved action = "review_request_removed"
	Created              action = "created"
)

type Payload struct {
	Action            action      `json:"action"`
	PullRequest       PullRequest `json:"pull_request"`
	Repository        Repository  `json:"repository"`
	Review            Comment     `json:"review"`
	Comment           Comment     `json:"comment"`
	Assignee          User        `json:"assignee"`
	RequestedReviewer User        `json:"requested_reviewer"`
	Sender            User        `json:"sender"`
}
