package main

import (
	"fmt"
	"gitlab-hooks-listener/src/project"
)

func main() {

	s := project.GetProject("fxg2ory")

	fmt.Println(s.GetConfig())

}
