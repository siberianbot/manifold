package config

type Configuration struct {
	ProjectTarget   *ProjectTarget   `yaml:"project"`
	WorkspaceTarget *WorkspaceTarget `yaml:"workspace"`
}
