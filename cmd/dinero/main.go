package main

import (
	"github.com/eikc/dinero-cli/internal/cmd"
	"os"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	root := cmd.NewRoot()

	_, err := root.ExecuteC()

	return err
}
