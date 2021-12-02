package config

type Configuration struct {
	Project   *Project   `yaml:"project"`
	Workspace *Workspace `yaml:"workspace"`
}

type Project struct {
	Name         string              `yaml:"name"`
	Dependencies []ProjectDependency `yaml:"dependencies"`
	Steps        []ProjectStep       `yaml:"steps"`
}

type ProjectDependency struct {
	Path    string `yaml:"path"`
	Project string `yaml:"project"`
}

type ProjectStep map[string]interface{}

type Workspace struct {
	Name     string             `yaml:"name"`
	Includes []WorkspaceInclude `yaml:"includes"`
}

type WorkspaceInclude string
