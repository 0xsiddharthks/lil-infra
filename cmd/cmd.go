package cmd

import (
	"os"

	"github.com/0xsiddharthks/lil-infra/git"
	"github.com/0xsiddharthks/lil-infra/lib/customError"
	"github.com/0xsiddharthks/lil-infra/lib/parser"
)

func Run() error {
	parsedCommand, err := parser.Parse(os.Args[1:])
	if err != nil {
		return err
	}

	switch parsedCommand.Command {
	case parser.Git:
		return git.Run(parsedCommand)
	case parser.Grep:
	case parser.Sqlite:
	case parser.Docker:
	case parser.Redis:
		return customError.NotImplementedError{}
	}
	return customError.NotImplementedError{}
}
