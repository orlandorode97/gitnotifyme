package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gen2brain/beeep"
)

func main() {

	http.HandleFunc("/events", receiveEvent)
	fmt.Println("Starting webhook")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}

func receiveEvent(w http.ResponseWriter, r *http.Request) {
	payload := Payload{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	if err = json.Unmarshal(body, &payload); err != nil {
		log.Fatal(err)
	}

	prNumber := strconv.Itoa(payload.PullRequest.PRNumber)
	var title, message string

	switch payload.Action {
	case Closed, Reopened:
		title = "A PR is " + string(payload.Action)
		message = fmt.Sprintf("The user %s %s a PR: #%s %s from %s", payload.PullRequest.User.Login, payload.Action, prNumber, payload.PullRequest.Title, payload.Repository.FullName)
		if err = sendNotification(title, message); err != nil {
			log.Fatal(err)
		}
	case Created, Submitted:
		title = "New PR comment "
		message = fmt.Sprintf("The user %s has comment the PR #%s %s from %s", payload.Sender.Login, prNumber, payload.PullRequest.Title, payload.Repository.FullName)
		if err = sendNotification(title, message); err != nil {
			log.Fatal(err)
		}
	case Assigned:
		title = "You've been assigned"
		message = fmt.Sprintf("You have been assigned to the PR #%s %s from %s", prNumber, payload.PullRequest.Title, payload.Repository.FullName)
		if err = sendNotification(title, message); err != nil {
			log.Fatal(err)
		}
	case ReviewRequested:
		title = "You've been requested"
		message = fmt.Sprintf("You have been requested to review the PR #%s %s from %s", prNumber, payload.PullRequest.Title, payload.Repository.FullName)
		if err = sendNotification(title, message); err != nil {
			log.Fatal(err)
		}
	case ReviewRequestRemoved:
		title = "You've been removed to review"
		message = fmt.Sprintf("You have been removed to review the PR #%s %s from %s", prNumber, payload.PullRequest.Title, payload.Repository.FullName)
		if err = sendNotification(title, message); err != nil {
			log.Fatal(err)
		}
	}
}

func sendNotification(title, message string) error {
	return beeep.Notify(title, message, "")
}
