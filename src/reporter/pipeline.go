package reporter

import (
	"encoding/json"
	"log"
)

type PipelineWebhookResponse struct {
	Project ProjectResponse      `json:"project"`
	Commit  SimpleCommitResponse `json:"commit"`
	Builds  []BuildResponse      `json:"builds"`
}

func (gr *gitlabReporter) PipelineWebhook() {

	if gr.p.GetEvents().Pipeline == false {
		return
	}

	message := preparePipelineWebhookMessage(gr)

	err := gr.t.SendMessage(message)

	log.Print(err)

}

func preparePipelineWebhookMessage(gr *gitlabReporter) string {

	body := PipelineWebhookResponse{}

	err := json.NewDecoder(gr.c.Request().Body).Decode(&body)

	if err != nil {
		log.Print(err)
		return "err in parsing data"
	}

	projectName := body.Project.Name
	builds := body.Builds

	message := gr.t.Bold(gr.p.GetTitles().Pipeline)
	message = message + gr.t.NewLine()
	message = message + gr.t.Bold(projectName)
	message = message + gr.t.Text("status")
	message = message + gr.t.Divider()

	for _, build := range builds {

		message = message + gr.t.Bold(build.Name+": ")

		message = message + gr.t.Text(build.Status)

		message = message + gr.t.NewLine()
	}

	return message
}
