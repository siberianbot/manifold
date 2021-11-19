package command_step

import "manifold/internal/steps"

type commandStep struct {
	executor steps.StepExecutor
	cmd      string
}

func (step commandStep) Execute() error {
	return step.executor.Execute(step)
}

func newCommandStep(cmd string, executor steps.StepExecutor) (step *commandStep) {
	step = new(commandStep)
	step.cmd = cmd
	step.executor = executor

	return step
}
