package middlwares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	appConfig "gitlab-hooks-listener/config"
	"gitlab-hooks-listener/src/project"
	"net/http"
	"strconv"
)

const GitlabSecretHeader = "Gitlab-Token"

const (
	gitlab       = "gitlab"
	defaultRealm = "Restricted"
)

type SecureGitlabConfig struct {
	Skipper       middleware.Skipper
	GitlabProject project.Project
	Realm         string
}

type SecureGitlabValidator func(string, string, echo.Context) (bool, error)

var DefaultBasicAuthConfig = SecureGitlabConfig{
	Skipper: middleware.DefaultSkipper,
	Realm:   defaultRealm,
}

func SecureGitlab(GitlabProject project.Project) echo.MiddlewareFunc {
	c := DefaultBasicAuthConfig
	c.GitlabProject = GitlabProject
	return BasicAuthWithConfig(c)
}

func BasicAuthWithConfig(config SecureGitlabConfig) echo.MiddlewareFunc {

	if config.Skipper == nil {
		config.Skipper = DefaultBasicAuthConfig.Skipper
	}

	if config.Realm == "" {
		config.Realm = defaultRealm
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			if config.Skipper(c) {
				return next(c)
			}

			auth := c.Request().Header.Get(GitlabSecretHeader)

			if auth != "" {
				if auth == appConfig.NewConfig().GetString("http-server.gitlab-secret-token") {
					return next(c)
				}
			}

			realm := defaultRealm
			if config.Realm != defaultRealm {
				realm = strconv.Quote(config.Realm)
			}

			// Need to return `401` for browsers to pop-up login box.
			c.Response().Header().Set(echo.HeaderWWWAuthenticate, gitlab+" realm="+realm)
			return echo.NewHTTPError(http.StatusForbidden, "invalid gitlab token")
		}
	}
}
