package main

import (
	"gitlab-hooks-listener/src/api"
)

func main() {

	service := api.NewApiService()

	service.StartHttpServer()

}
