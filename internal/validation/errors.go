package validation

const (
	EmptyManifoldName   = "name is empty"
	InvalidManifoldName = "name \"%s\" does not matches regex pattern \"%s\""
	EmptyPath           = "path is empty"
	InvalidPath         = "invalid path \"%s\""

	EmptyProjectDependency           = "project dependency does not declare dependent project name nor path to dependent project"
	DependencyWithBothProjectAndPath = "project dependency declares both dependent project name and path to dependent project"
	InvalidDependentProject          = "invalid dependent project: %v"

	StepNotMatch = "step does not matches any known toolchain"
)
