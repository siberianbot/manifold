package config

type ProjectTarget struct {
	Name                string              `yaml:"name"`
	ProjectDependencies []ProjectDependency `yaml:"dependencies"`
	Steps               []Step              `yaml:"steps"`
}

type WorkspaceTarget struct {
	Name     string   `yaml:"name"`
	Includes []string `yaml:"includes"`
}
