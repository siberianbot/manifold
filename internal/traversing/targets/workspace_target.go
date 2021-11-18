package targets

type WorkspaceTarget struct {
	name string
	path string
}

func (workspace WorkspaceTarget) Name() string {
	return workspace.name
}

func (workspace WorkspaceTarget) Path() string {
	return workspace.path
}

func (WorkspaceTarget) Kind() TargetKind {
	return WorkspaceTargetKind
}
