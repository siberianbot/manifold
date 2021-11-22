package validation

import (
	"errors"
	"fmt"
)

const (
	EmptyManifoldName   = "name is empty"
	InvalidManifoldName = "name \"%s\" does not matches regex pattern \"%s\""
	EmptyPath           = "path is empty"
	InvalidPath         = "invalid path \"%s\""

	StepResolveFailed          = "step resolution failed: %v"
	EmptyStep                  = "step is empty"
	StepFailed                 = "step %s failed: %v"
	StepNotMatchedAnyToolchain = "step %s does not matches any known toolchain"
	StepMatchesManyToolchains  = "step matches many known toolchain"

	CmdStepIsInvalid = "definition should be a non-empty string"

	EmptyConfiguration                   = "configuration is empty"
	ConfigurationWithProjectAndWorkspace = "configuration contains both definitions of project and workspace"

	InvalidProject           = "project is invalid: %v"
	InvalidProjectDependency = "project dependency is invalid: %v"
	InvalidWorkspace         = "workspace is invalid: %v"
	InvalidWorkspaceInclude  = "workspace include is invalid: %v"

	EmptyProjectDependency                  = "project dependency does not declare dependent project name nor path to dependent project"
	ProjectDependencyWithBothProjectAndPath = "project dependency declares both dependent project name and path to dependent project"
	EmptyWorkspaceInclude                   = "workspace include is empty"

	NotManifoldPath = "path \"%s\" does not contain manifold configuration"
)

func NewError(msg string, args ...interface{}) error {
	if len(args) > 0 {
		return errors.New(fmt.Sprintf(msg, args...))
	} else {
		return errors.New(msg)
	}
}
