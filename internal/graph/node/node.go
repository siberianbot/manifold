package node

import (
	"manifold/internal/config"
	"manifold/internal/step"
)

type Node interface {
	Name() string
	Path() string
	Dependencies() []Dependency
	Steps() []step.Step
}

type Builder interface {
	FromConfig(path string, cfg *config.Configuration) (Node, error)
}

type DependencyKind uint8

const (
	NamedDependencyKind DependencyKind = iota
	PathedDependencyKind
)

type Dependency interface {
	Kind() DependencyKind
	Value() string
}

type fixedDependency struct {
	kind  DependencyKind
	value string
}

func NewDependency(kind DependencyKind, value string) Dependency {
	return &fixedDependency{kind: kind, value: value}
}

func (d *fixedDependency) Kind() DependencyKind {
	return d.kind
}

func (d *fixedDependency) Value() string {
	return d.value
}
