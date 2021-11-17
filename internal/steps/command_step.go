package steps

import (
	"github.com/google/shlex"
	"io"
	"log"
	"os/exec"
)

type CommandStep struct {
	Command string
}

func (CommandStep) Kind() StepKind {
	return CommandStepKind
}

func (step CommandStep) Execute() error {
	// TODO:
	// 1. remove usage of Windows' cmd - it is not cross-platform
	// 2. interpret command as command block or script
	// 3. introduce a better way of output redirection

	cmdSplit, _ := shlex.Split(step.Command)

	cmd := exec.Command("cmd", "/c")
	cmd.Args = append(cmd.Args, cmdSplit...)

	mw := io.MultiWriter(log.Writer())
	cmd.Stdout = mw
	cmd.Stderr = mw

	cmdErr := cmd.Run()

	return cmdErr
}
