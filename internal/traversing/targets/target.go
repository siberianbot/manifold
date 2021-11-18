package targets

type TargetKind int8

const (
	ProjectTargetKind TargetKind = iota
	WorkspaceTargetKind
)

type Target interface {
	Name() string
	Path() string
	Kind() TargetKind
}
