package api

import (
	"gitlab-hooks-listener/pkg/http_server"
)

type Api interface {
	StartHttpServer()
	RegisterRoutes()
	PrintRegisteredRoutes()
}

type api struct {
	hs http_server.HttpServer
}

func NewApiService() Api {

	return &api{http_server.GetHttpServer()}
}
