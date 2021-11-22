package config

type ProjectTarget struct {
	Name                string              `yaml:"name"`
	ProjectDependencies []ProjectDependency `yaml:"dependencies"`
	Steps               []Step              `yaml:"steps"`
}

func (p *ProjectTarget) Dependencies() []Dependency {
	dependencies := make([]Dependency, len(p.ProjectDependencies))

	for idx, projectDependency := range p.ProjectDependencies {
		dependencies[idx] = projectDependency.ToDependency()
	}

	return dependencies
}
