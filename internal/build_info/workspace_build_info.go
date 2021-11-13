package build_info

type WorkspaceBuildInfo struct {
	name string
	path string
}

func (workspace WorkspaceBuildInfo) Name() string {
	return workspace.name
}

func (workspace WorkspaceBuildInfo) Path() string {
	return workspace.path
}

func (WorkspaceBuildInfo) Kind() BuildInfoKind {
	return WorkspaceBuildInfoKind
}
