package parser

import "github.com/siddharth2010/lil-infra/lib/customError"

const (
	Redis  = iota
	Git    = iota
	Grep   = iota
	Sqlite = iota
	Docker = iota
)

type ParsedCommand struct {
	Command    rune
	Options    []string
	Subcommand string
	Args       []string
}

func Parse(command []string) (ParsedCommand, error) {
	return ParsedCommand{}, customError.NotImplementedError{}
}
