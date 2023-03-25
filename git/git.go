package git

import (
	"github.com/siddharth2010/lil-infra/lib/customError"
	"github.com/siddharth2010/lil-infra/lib/parser"
)

func Run(command parser.ParsedCommand) error {
	switch command.Subcommand {
	case "init":
		return handleInit()
		// create .git , .git/objects, .git/refs, and .git/HEAD ("ref: refs/heads/master\n")
	case "cat-file":
		// cat-file -p <hash>
		return handleCatFile()
		// look into .git/objects/<hash[:2]>/<hash[2:]> for the file
		// decompress using zlib
		// strip headers, and print
	case "hash-object":
		// hash-object -w <file>
		return handleHashObject()
		// add header, compute checksum, compress, and store to .git/objects/<hash[:2]>/<hash[2:]>
	case "ls-tree":
		// ls-tree <hash>
		return handleLsTree()
		// confirm it's a tree, from header, and print the contents
	case "write-tree":
		// write-tree
		return handleWriteTree()
		// create a tree object current directory and store recursively in .git/objects
	case "commit-tree":
		// commit-tree <hash> -m <message>
		return handleCommitTree()
		// write tree, and link to parent commit, and set message
	case "clone":
		// clone <url> <path>
		return handleClone()
	default:
		return customError.NotImplementedError{}
	}
}

func handleInit() error {
	return customError.NotImplementedError{}
}

func handleCatFile() error {
	return customError.NotImplementedError{}
}

func handleHashObject() error {
	return customError.NotImplementedError{}
}

func handleLsTree() error {
	return customError.NotImplementedError{}
}

func handleWriteTree() error {
	return customError.NotImplementedError{}
}

func handleCommitTree() error {
	return customError.NotImplementedError{}
}

func handleClone() error {
	return customError.NotImplementedError{}
}
