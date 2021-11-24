package validation

const (
	EmptyConfiguration     = "configuration is empty"
	AmbiguousConfiguration = "configuration is ambiguous: it contains both definitions of project and workspace"

	InvalidProject   = "project is invalid: %v"
	InvalidWorkspace = "workspace is invalid: %v"

	EmptyWorkspaceInclude   = "workspace include is empty"
	InvalidWorkspaceInclude = "workspace include is invalid: %v"

	EmptyProjectDependency     = "project dependency is empty"
	AmbiguousProjectDependency = "project dependency is ambiguous: it declares both dependent project name and path to dependent project"
	InvalidProjectDependency   = "project dependency is invalid: %v"
)
