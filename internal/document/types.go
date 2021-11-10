package document

type Document struct {
	Project   *Project   `yaml:"project"`
	Workspace *Workspace `yaml:"workspace"`
}

type Project struct {
	Name         string       `yaml:"name"`
	Dependencies []Dependency `yaml:"dependencies"`
	Steps        []Step       `yaml:"steps"`
}

type Workspace struct {
	Name     string   `yaml:"name"`
	Includes []string `yaml:"includes"`
}

type Dependency struct {
	Path    string `yaml:"path"`
	Project string `yaml:"project"`
}

type Step struct {
	Command string `yaml:"cmd"`
}
