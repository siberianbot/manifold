package targets

import (
	"manifold/internal/steps"
)

type ProjectTarget struct {
	name string
	path string

	Steps []steps.Step
}

func (project ProjectTarget) Name() string {
	return project.name
}

func (project ProjectTarget) Path() string {
	return project.path
}

func (ProjectTarget) Kind() TargetKind {
	return ProjectTargetKind
}
