package config

type TargetKind int8

const (
	ProjectTargetKind TargetKind = iota
	WorkspaceTargetKind
)

type Target interface {
	Validatable

	Name() string
	Kind() TargetKind
	Dependencies() []Dependency
}
