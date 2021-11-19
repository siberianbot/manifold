package command_step

import (
	"github.com/google/shlex"
	"io"
	"log"
	"manifold/internal/steps"
	"os/exec"
)

type commandStepExecutor struct {
	//
}

func (commandStepExecutor) Execute(step steps.Step) error {
	cmdArgs, _ := shlex.Split(step.(commandStep).cmd)

	cmd := exec.Command(cmdArgs[0])

	if len(cmdArgs) > 1 {
		cmd.Args = append(cmd.Args, cmdArgs[1:]...)
	}

	mw := io.MultiWriter(log.Writer())
	cmd.Stdout = mw
	cmd.Stderr = mw

	return cmd.Run()
}

func newCommandStepExecutor() *commandStepExecutor {
	return new(commandStepExecutor)
}
