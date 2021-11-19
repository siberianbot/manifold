package config

type ConfigurationDefinition struct {
	Project   *ProjectDefinition   `yaml:"project"`
	Workspace *WorkspaceDefinition `yaml:"workspace"`
}

type ProjectDefinition struct {
	Name         string                 `yaml:"name"`
	Dependencies []DependencyDefinition `yaml:"dependencies"`
	Steps        []StepDefinition       `yaml:"steps"`
}

type WorkspaceDefinition struct {
	Name     string              `yaml:"name"`
	Includes []IncludeDefinition `yaml:"includes"`
}

type DependencyDefinition struct {
	Path    string `yaml:"path"`
	Project string `yaml:"project"`
}

type StepDefinition map[string]interface{}

type IncludeDefinition string
