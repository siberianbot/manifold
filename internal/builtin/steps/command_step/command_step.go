package command_step

import (
	"github.com/google/shlex"
	"io"
	"log"
	"manifold/internal/errors"
	"manifold/internal/step"
	"manifold/internal/step/provider"
	"os/exec"
)

const (
	name          = "cmd"
	stepIsInvalid = "definition should be a non-empty string"
)

type commandStep struct {
	cmd string
}

func (commandStep) Name() string {
	return name
}

type commandProxy struct {
	//
}

func (commandProxy) CreateFrom(value interface{}) (step.Step, error) {
	cmd, ok := value.(string)

	if !ok || cmd == "" {
		return nil, errors.NewError(stepIsInvalid)
	}

	return &commandStep{cmd: cmd}, nil
}

func (commandProxy) Execute(context step.ExecutorContext) error {
	cmdArgs, _ := shlex.Split(context.Step().(*commandStep).cmd)

	cmd := exec.Command(cmdArgs[0])

	cmd.Dir = context.Dir()

	if len(cmdArgs) > 1 {
		cmd.Args = append(cmd.Args, cmdArgs[1:]...)
	}

	mw := io.MultiWriter(log.Writer())
	cmd.Stdout = mw
	cmd.Stderr = mw

	return cmd.Run()
}

func PopulateOptions(options *provider.Options) {
	proxy := new(commandProxy)
	options.Factories[name] = proxy
	options.Executors[name] = proxy
}
