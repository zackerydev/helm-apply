package main

import (
	"errors"
	"github.com/zackerydev/helm-apply/cmd"
	"os"
)

type Error struct {
	error
	Code int
}

func main() {
	if err := cmd.New().Execute(); err != nil {
		var cmdErr Error
		switch {
		case errors.As(err, &cmdErr):
			os.Exit(cmdErr.Code)
		default:
			os.Exit(1)
		}
	}
}
