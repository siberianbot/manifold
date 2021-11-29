package command_step

import (
	"github.com/google/shlex"
	"io"
	"log"
	"manifold/internal/errors"
	"manifold/internal/steps"
	"os/exec"
)

const (
	Name          = "cmd"
	StepIsInvalid = "definition should be a non-empty string"
)

type commandStep struct {
	cmd string
}

func (commandStep) Name() string {
	return Name
}

func executeStep(step steps.Step, _ *steps.ExecutorContext) error {
	cmdArgs, _ := shlex.Split(step.(*commandStep).cmd)

	cmd := exec.Command(cmdArgs[0])

	if len(cmdArgs) > 1 {
		cmd.Args = append(cmd.Args, cmdArgs[1:]...)
	}

	mw := io.MultiWriter(log.Writer())
	cmd.Stdout = mw
	cmd.Stderr = mw

	return cmd.Run()
}

func newStep(definition interface{}) (steps.Step, error) {
	cmd, ok := definition.(string)

	if !ok || cmd == "" {
		return nil, errors.NewError(StepIsInvalid)
	}

	step := new(commandStep)
	step.cmd = cmd

	return step, nil
}

func PopulateOptions(options *steps.ProviderOptions) {
	options.Factories[Name] = newStep
	options.Executors[Name] = executeStep
}
