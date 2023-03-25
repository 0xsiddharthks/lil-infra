package customError

import (
	"fmt"
	"os"
)

type NotImplementedError struct {
}

func (e NotImplementedError) Error() string {
	return "Not implemented yet"
}

type MyError struct {
	message string
}

func (e MyError) Error() string {
	return e.message
}

func HandleError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %s", err)
	os.Exit(1)
}
