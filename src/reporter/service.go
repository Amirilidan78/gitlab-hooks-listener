package reporter

import (
	"github.com/labstack/echo/v4"
	"gitlab-hooks-listener/pkg/telegram"
	"gitlab-hooks-listener/src/project"
)

type GitlabReporter interface {
	CommentWebhook()
	DeploymentWebhook()
	FeatureFlagWebhook()
	GroupWebhook()
	IssueWebhook()
	JobWebhook()
	MergeRequestWebhook()
	PipelineWebhook()
	PushWebhook()
	ReleaseWebhook()
	SubGroupWebhook()
	TagWebhook()
	WikiPageWebhook()
}

type gitlabReporter struct {
	c echo.Context
	p project.Project
	t telegram.Telegram
}

func NewGitlabReporter(c echo.Context, p project.Project) GitlabReporter {
	return &gitlabReporter{c: c, p: p, t: telegram.GetTelegramClient()}
}
