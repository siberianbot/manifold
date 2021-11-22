package config

type DependencyKind int8

const (
	UnknownDependencyKind DependencyKind = iota
	ProjectDependencyKind
	PathDependencyKind
)

type Dependency struct {
	kind  DependencyKind
	value string
}

func (d *Dependency) Kind() DependencyKind {
	return d.kind
}

func (d *Dependency) Value() string {
	return d.value
}
