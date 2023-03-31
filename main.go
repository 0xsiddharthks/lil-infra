package main

import (
	"github.com/0xsiddharthks/lil-infra/cmd"
	"github.com/0xsiddharthks/lil-infra/lib/customError"
)

func main() {
	err := cmd.Run()
	if err != nil {
		customError.HandleError(err)
	}
}
