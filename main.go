package main

import (
	"github.com/siddharth2010/lil-infra/cmd"
	"github.com/siddharth2010/lil-infra/lib/customError"
)

func main() {
	err := cmd.Run()
	if err != nil {
		customError.HandleError(err)
	}
}
