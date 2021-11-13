package steps

import (
	"github.com/google/shlex"
	"os/exec"
)

type CommandStep struct {
	Command string
}

func (step CommandStep) Kind() StepKind {
	return CommandStepKind
}

func (step CommandStep) Execute() error {
	// TODO:
	// 1. remove usage of Windows' cmd - it is not cross-platform
	// 2. interpret command as command block or script

	cmdSplit, _ := shlex.Split(step.Command)

	cmd := exec.Command("cmd", "/c")
	cmd.Args = append(cmd.Args, cmdSplit...)
	cmdErr := cmd.Run()

	return cmdErr
}
