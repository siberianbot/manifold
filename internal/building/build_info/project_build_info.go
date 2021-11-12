package build_info

import "manifold/internal/building/steps"

type ProjectBuildInfo struct {
	name string
	path string

	Steps []steps.Step
}

func (project ProjectBuildInfo) Name() string {
	return project.name
}

func (project ProjectBuildInfo) Path() string {
	return project.path
}

func (ProjectBuildInfo) Kind() BuildInfoKind {
	return ProjectBuildInfoKind
}
