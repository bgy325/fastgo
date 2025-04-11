package main

import (
	"os"

	"github.com/bgy325/fastgo/cmd/fg-apiserver/app"

	_ "go.uber.org/automaxprocs"
)

func main() {
	command := app.NewFastGOCommand()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
