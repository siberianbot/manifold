package config

type ProjectDependency struct {
	Path    string `yaml:"path"`
	Project string `yaml:"project"`
}

func (p *ProjectDependency) ToDependency() Dependency {
	switch {
	case p.Project != "" && p.Path != "":
		return Dependency{kind: UnknownDependencyKind}

	case p.Project != "":
		return Dependency{kind: ProjectDependencyKind, value: p.Project}

	case p.Path != "":
		return Dependency{kind: PathDependencyKind, value: p.Path}

	default:
		return Dependency{kind: UnknownDependencyKind}
	}
}
