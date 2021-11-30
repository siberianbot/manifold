package main

import (
	"github.com/mitchellh/cli"
	"log"
	"manifold/internal/backend"
	"os"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	b := backend.NewBackend()

	c := cli.NewCLI(Name, version())
	c.Args = os.Args[1:]
	c.HelpFunc = helpFunc
	c.Commands = map[string]cli.CommandFactory{
		BuildCommandName: func() (cli.Command, error) { return buildCommandFactory(b) },
	}

	exitCode, err := c.Run()

	if err != nil {
		log.Println(err)
		return -1
	}

	return exitCode
}
