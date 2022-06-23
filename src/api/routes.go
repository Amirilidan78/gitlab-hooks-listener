package api

import (
	"encoding/json"
	"fmt"
	"gitlab-hooks-listener/src/api/handlers"
)

func (a *api) RegisterRoutes() {

	router := a.hs.GetEcho()

	tpg := router.Group("/projects")
	tpg.POST("/:name/events", handlers.HandleGitlabEvent)

}

func (a *api) PrintRegisteredRoutes() {

	data, err := json.MarshalIndent(a.hs.GetEcho().Routes(), "", "  ")

	if err == nil {
		fmt.Println(data)
	}
}
