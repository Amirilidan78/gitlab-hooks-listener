package http_server

import (
	"github.com/labstack/echo/v4"
	"gitlab-hooks-listener/config"
	"sync"
)

var lock = &sync.Mutex{}

var singleton *httpServer

type HttpServer interface {
	initialServer()
	getServerPort() string
	GetEcho() *echo.Echo
	Listen()
}

type httpServer struct {
	c config.Config
	e *echo.Echo
}

func GetHttpServer() HttpServer {

	if singleton == nil {
		lock.Lock()
		defer lock.Unlock()
		singleton = &httpServer{config.NewConfig(), echo.New()}
		singleton.initialServer()
	}

	return singleton
}
