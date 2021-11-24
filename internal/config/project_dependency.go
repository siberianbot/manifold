package config

type ProjectDependency struct {
	Path    string `yaml:"path"`
	Project string `yaml:"project"`
}
