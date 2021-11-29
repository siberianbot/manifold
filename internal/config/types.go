package config

type Configuration struct {
	ProjectTarget   *ProjectTarget   `yaml:"project"`
	WorkspaceTarget *WorkspaceTarget `yaml:"workspace"`
}

type ProjectTarget struct {
	Name                string              `yaml:"name"`
	ProjectDependencies []ProjectDependency `yaml:"dependencies"`
	Steps               []Step              `yaml:"steps"`
}

type ProjectDependency struct {
	Path    string `yaml:"path"`
	Project string `yaml:"project"`
}

type Step map[string]interface{}

type WorkspaceTarget struct {
	Name     string   `yaml:"name"`
	Includes []string `yaml:"includes"`
}
