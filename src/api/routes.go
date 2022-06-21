package api

import "gitlab-hooks-listener/src/api/handlers"

func (a *api) registerRoutes() {

	router := a.hs.GetEcho()

	tpg := router.Group("/api")
	tpg.GET("/test", handlers.HandleTest)

}
