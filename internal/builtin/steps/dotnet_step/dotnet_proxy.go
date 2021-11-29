package dotnet_step

import (
	"io"
	"log"
	"manifold/internal/errors"
	"manifold/internal/steps"
	"os/exec"
)

const (
	Name = "dotnet"
)

type dotnetStep struct {
	path string
}

func (dotnetStep) Name() string {
	return Name
}

type dotnetProxy struct {
	//
}

func (dotnetProxy) newStep(definition interface{}) (steps.Step, error) {
	path, ok := definition.(string)

	if !ok || path == "" {
		return nil, errors.NewError("dotnet step should be a valid path to project or solution file")
	}

	return &dotnetStep{path: path}, nil
}

func (dotnetProxy) executeStep(step steps.Step) error {
	dotnetCmd := exec.Command("dotnet", "build", step.(dotnetStep).path)

	mw := io.MultiWriter(log.Writer())
	dotnetCmd.Stdout = mw
	dotnetCmd.Stderr = mw

	return dotnetCmd.Run()
}

func PopulateOptions(options *steps.ProviderOptions) {
	proxy := new(dotnetProxy)
	options.Factories[Name] = proxy.newStep
	options.Executors[Name] = proxy.executeStep
}