package config

import "manifold/internal/validation"

type ProjectTarget interface {
	Target

	Dependencies() []Dependency
	Steps() []Step
}

type projectTarget struct {
	ProjectName     string       `yaml:"name"`
	RawDependencies []dependency `yaml:"dependencies"`
	RawSteps        []Step       `yaml:"steps"`
}

func (p *projectTarget) Validate(ctx ValidationContext) error {
	if err := validation.ValidateManifoldName(p.ProjectName); err != nil {
		return err
	}

	for _, dependency := range p.RawDependencies {
		if err := dependency.Validate(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (p *projectTarget) Name() string {
	return p.ProjectName
}

func (projectTarget) Kind() TargetKind {
	return ProjectTargetKind
}

func (p *projectTarget) Dependencies() []Dependency {
	dependencies := make([]Dependency, len(p.RawDependencies))

	for idx := 0; idx < len(p.RawDependencies); idx++ {
		dependencies[idx] = &(p.RawDependencies[idx])
	}

	return dependencies
}

func (p *projectTarget) Steps() []Step {
	return p.RawSteps
}
