package app

import (
	"fmt"
	"os"
)

func CLI(args []string) int {
	var app appEnv
	err := app.fromArgs(args)
	if err != nil {
		return 2
	}
	if err = app.run(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		return 1
	}
	
	return 0
}

type appEnv struct {
}

func (app *appEnv) fromArgs(args []string) error {
	return nil
}

func (app *appEnv) run() error {
	return nil
}
