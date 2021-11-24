package graph

import (
	"manifold/internal/config"
	"manifold/internal/utils"
)

type NodeDependencyKind uint8

const (
	ByNameDependencyKind NodeDependencyKind = iota
	ByPathDependencyKind
)

type NodeDependency interface {
	Kind() NodeDependencyKind
	Value() string
}

type byNameNodeDependency struct {
	Name string
}

func (byNameNodeDependency) Kind() NodeDependencyKind {
	return ByNameDependencyKind
}

func (dependency *byNameNodeDependency) Value() string {
	return dependency.Name
}

type byPathNodeDependency struct {
	Path string
}

func (byPathNodeDependency) Kind() NodeDependencyKind {
	return ByPathDependencyKind
}

func (dependency *byPathNodeDependency) Value() string {
	return dependency.Path
}

func FromProjectDependency(dependency config.ProjectDependency, dir string) NodeDependency {
	switch {
	case dependency.Project != "":
		return &byNameNodeDependency{Name: dependency.Project}

	case dependency.Path != "":
		return &byPathNodeDependency{Path: utils.BuildPath(dir, dependency.Path)}

	default:
		panic("dependency is empty")
	}
}

func FromWorkspaceInclude(include string, dir string) NodeDependency {
	return &byPathNodeDependency{Path: utils.BuildPath(dir, include)}
}
