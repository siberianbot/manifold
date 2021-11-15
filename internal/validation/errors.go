package validation

const (
	InvalidManifoldName = "name \"%s\" does not matches regex pattern \"%s\""
	EmptyPath           = "path is empty"
	InvalidPath         = "invalid path \"%s\""

	EmptyProjectDependency  = "project dependency does not declare dependent project name nor path to dependent project"
	InvalidDependentProject = "invalid dependent project: %v"
)
