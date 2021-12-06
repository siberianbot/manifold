package main

import (
	"fmt"
	"github.com/mitchellh/cli"
	"log"
	"manifold/internal/backend"
	"os"
)

type buildCommand struct {
	backend backend.Interface
}

func (cmd *buildCommand) Help() string {
	content := `%s

Build project or workspace defined in configuration. Path to configuration should 
be valid Manifold configuration file or path to directory with configuration. Path 
to current directory is used if path is not defined.

%s`

	return fmt.Sprintf(content, usage(BuildCommandName, UsageBuildArgs), version())
}

func (cmd *buildCommand) Synopsis() string {
	return "Build project or workspace defined in configuration"
}

func (cmd *buildCommand) Run(args []string) int {
	buildOptions := backend.BuildOptions{}

	if len(os.Args) == 2 {
		buildOptions.Path, _ = os.Getwd()
	} else {
		buildOptions.Path = os.Args[2]
	}

	err := cmd.backend.Build(buildOptions)

	if err != nil {
		log.Println(fmt.Sprintf("build failed: %v", err))
		return 1
	}

	log.Println("done")
	return 0
}

func buildCommandFactory(b backend.Interface) (cli.Command, error) {
	return &buildCommand{backend: b}, nil
}
