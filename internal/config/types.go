package config

type Configuration struct {
	ProjectTarget   *ProjectTarget   `yaml:"project"`
	WorkspaceTarget *WorkspaceTarget `yaml:"workspace"`
}

type ProjectTarget struct {
	Name                string              `yaml:"name"`
	ProjectDependencies []ProjectDependency `yaml:"dependencies"`
	Steps               []StepDefinition    `yaml:"steps"`
}

type ProjectDependency struct {
	Path    string `yaml:"path"`
	Project string `yaml:"project"`
}

type StepDefinition map[string]interface{}

type WorkspaceTarget struct {
	Name     string              `yaml:"name"`
	Includes []IncludeDefinition `yaml:"includes"`
}

type IncludeDefinition string
