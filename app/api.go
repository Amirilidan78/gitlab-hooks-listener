package main

import (
	"gitlab-hooks-listener/src/api"
)

func main() {

	service := api.NewApiService()

	service.RegisterRoutes()

	service.PrintRegisteredRoutes()

	service.StartHttpServer()

}
