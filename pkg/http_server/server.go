package http_server 

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (hs *httpServer) initialServer() {
	hs.e.Use(middleware.Logger())
	hs.e.Use(middleware.Recover())
}

func (hs *httpServer) GetEcho() *echo.Echo {
	return hs.e
}

func (hs *httpServer) getServerPort() string {
	return hs.c.GetString("http-server.port")
}

func (hs *httpServer) Listen() {
	hs.e.Logger.Fatal(hs.e.Start(":" + hs.getServerPort()))
}