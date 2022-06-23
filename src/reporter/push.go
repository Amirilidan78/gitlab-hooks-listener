package reporter

import (
	"encoding/json"
	"log"
	"strconv"
)

type PushWebhookResponse struct {
	Username     string           `json:"user_name"`
	Project      ProjectResponse  `json:"project"`
	TotalCommits int              `json:"total_commits_count"`
	Commits      []CommitResponse `json:"commits"`
}

func (gr *gitlabReporter) PushWebhook() {

	if gr.p.GetEvents().Push == false {
		return
	}

	message := preparePushWebhookMessage(gr)

	err := gr.t.SendMessage(message)

	log.Print(err)
}

func preparePushWebhookMessage(gr *gitlabReporter) string {

	body := PushWebhookResponse{}

	err := json.NewDecoder(gr.c.Request().Body).Decode(&body)

	if err != nil {
		log.Print(err)
		return "err in parsing data"
	}

	username := body.Username
	projectName := body.Project.Name
	totalCommits := strconv.Itoa(body.TotalCommits)
	commits := body.Commits

	message := gr.t.Bold(gr.p.GetTitles().Push)
	message = message + gr.t.NewLine()
	message = message + gr.t.Bold(username)
	message = message + gr.t.Text("pushed")
	message = message + gr.t.Bold(totalCommits)
	message = message + gr.t.Text("commits on project ")
	message = message + gr.t.Bold(projectName)
	message = message + gr.t.Divider()

	for _, commit := range commits {
		message = message + gr.t.Bold(commit.Author.Name+": ")
		message = message + gr.t.Copy(commit.Message)
		message = message + gr.t.NewLine()

		if len(commit.Added) > 0 {
			message = message + gr.t.Copy("Added files")
			message = message + gr.t.NewLine()
			for _, file := range commit.Added {
				message = message + gr.t.Copy(file)
				message = message + gr.t.NewLine()
			}
		}

		if len(commit.Modified) > 0 {
			message = message + gr.t.Copy("Modified files")
			message = message + gr.t.NewLine()
			for _, file := range commit.Modified {
				message = message + gr.t.Copy(file)
				message = message + gr.t.NewLine()
			}
		}

		if len(commit.Removed) > 0 {
			message = message + gr.t.Copy("Removed files")
			message = message + gr.t.NewLine()
			for _, file := range commit.Removed {
				message = message + gr.t.Copy(file)
				message = message + gr.t.NewLine()
			}
		}

		message = message + gr.t.Divider()
	}

	return message
}
