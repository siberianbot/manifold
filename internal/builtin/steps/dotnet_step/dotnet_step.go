package dotnet_step

import (
	"io"
	"log"
	"manifold/internal/errors"
	"manifold/internal/step"
	"manifold/internal/step/provider"
	"manifold/internal/utils"
	"os/exec"
)

const (
	name = "dotnet"
)

type dotnetStep struct {
	path string
}

func (dotnetStep) Name() string {
	return name
}

type dotnetProxy struct {
	//
}

func (dotnetProxy) CreateFrom(value interface{}) (step.Step, error) {
	path, ok := value.(string)

	if !ok || path == "" {
		return nil, errors.NewError("dotnet step should be a valid path to project or solution file")
	}

	return &dotnetStep{path: path}, nil
}

func (dotnetProxy) Execute(context step.ExecutorContext) error {
	targetPath := utils.BuildPath(context.Dir(), context.Step().(*dotnetStep).path)
	dotnetCmd := exec.Command("dotnet", "build", targetPath)

	mw := io.MultiWriter(log.Writer())
	dotnetCmd.Stdout = mw
	dotnetCmd.Stderr = mw

	return dotnetCmd.Run()
}

func PopulateOptions(options *provider.Options) {
	proxy := new(dotnetProxy)
	options.Factories[name] = proxy
	options.Executors[name] = proxy
}
