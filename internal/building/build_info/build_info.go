package build_info

type BuildInfoKind int8

const (
	ProjectBuildInfoKind BuildInfoKind = iota
	WorkspaceBuildInfoKind
)

type BuildInfo interface {
	Name() string
	Path() string
	Kind() BuildInfoKind
}
