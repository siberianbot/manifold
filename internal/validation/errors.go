package validation

const (
	InvalidManifoldName = "name \"%s\" does not matches regex pattern \"%s\""

	EmptyProjectDependency  = "project dependency does not declare dependent project name nor path to dependent project"
	InvalidDependentProject = "invalid dependent project: %v"
)
