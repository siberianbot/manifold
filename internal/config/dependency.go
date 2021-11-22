package config

type DependencyKind int8

const (
	UnknownDependencyKind DependencyKind = iota
	ProjectDependencyKind
	PathDependencyKind
)

type Dependency interface {
	Validatable

	Kind() DependencyKind
	Value() string
}
