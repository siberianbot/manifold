package config

type WorkspaceTarget struct {
	Name     string   `yaml:"name"`
	Includes []string `yaml:"includes"`
}

func (w *WorkspaceTarget) Dependencies() []Dependency {
	includes := make([]Dependency, len(w.Includes))

	for idx, include := range w.Includes {
		includes[idx] = Dependency{kind: PathDependencyKind, value: include}
	}

	return includes
}
