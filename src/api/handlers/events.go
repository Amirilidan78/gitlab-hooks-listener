package handlers

import (
	"github.com/labstack/echo/v4"
	"gitlab-hooks-listener/src/project"
	"gitlab-hooks-listener/src/reporter"
	"net/http"
)

const HeaderGitlabEvent = "X-Gitlab-Event"
const GitlabSecretHeader = "X-Gitlab-Token"

func HandleGitlabEvent(c echo.Context) error {

	p := project.GetProject(c.ParamValues()[0])

	if doesGitlabSecretMach(c, p) == false {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"error": "invalid gitlab secret token",
		})
	}

	r := reporter.NewGitlabReporter(c, p)

	switch c.Request().Header.Get(HeaderGitlabEvent) {
	case "Push Hook":
		r.PushWebhook()
		break
	case "Tag Push Hook":
		r.TagWebhook()
		break
	case "Issue Hook":
		r.IssueWebhook()
		break
	case "Note Hook":
		r.CommentWebhook()
		break
	case "Merge Request Hook":
		r.MergeRequestWebhook()
		break
	case "Wiki Page Hook":
		r.WikiPageWebhook()
		break
	case "Pipeline Hook":
		r.PipelineWebhook()
		break
	case "Job Hook":
		r.JobWebhook()
		break
	case "Deployment Hook":
		r.DeploymentWebhook()
		break
	case "Member Hook":
		r.GroupWebhook()
		break
	case "Subgroup Hook":
		r.SubGroupWebhook()
		break
	case "Feature Flag Hook":
		r.FeatureFlagWebhook()
	case "Release Hook":
		r.ReleaseWebhook()
		break
	default:
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "received invalid event",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"success": true,
	})
}

func doesGitlabSecretMach(c echo.Context, p project.Project) bool {

	return c.Request().Header.Get(GitlabSecretHeader) == p.GetConfig().GitlabSecret
}
