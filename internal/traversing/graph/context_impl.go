package graph

import (
	"fmt"
	"manifold/internal/steps"
	"manifold/internal/steps/command_step"
)

type contextImpl struct {
	errors   []string
	warnings []string

	path string
	dir  string

	stepProvider steps.StepProvider
}

func (ctx *contextImpl) GetStepProvider() steps.StepProvider {
	return ctx.stepProvider
}

func (ctx *contextImpl) AddError(message string, params ...interface{}) {
	if len(params) > 0 {
		message = fmt.Sprintf(message, params...)
	}

	ctx.errors = append(ctx.errors, message)
}

func (ctx *contextImpl) AddWarning(message string, params ...interface{}) {
	if len(params) > 0 {
		message = fmt.Sprintf(message, params...)
	}

	ctx.warnings = append(ctx.warnings, message)
}

func (ctx contextImpl) IsValid() bool {
	return len(ctx.errors) == 0
}

func (ctx contextImpl) CurrentFile() string {
	return ctx.path
}

func newContext(path string) (ctx *contextImpl, err error) {
	absPath, absDir, err := processPath(path)

	if err != nil {
		return nil, err
	}

	ctx = new(contextImpl)
	ctx.errors = make([]string, 0)
	ctx.warnings = make([]string, 0)
	ctx.path = absPath
	ctx.dir = absDir

	ctx.stepProvider = steps.NewDefaultStepProvider(
		command_step.NewStepFactory())

	return ctx, nil
}
