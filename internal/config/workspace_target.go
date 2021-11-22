package config

import "manifold/internal/validation"

type WorkspaceTarget interface {
	Target
}

type workspaceTarget struct {
	WorkspaceName string   `yaml:"name"`
	RawIncludes   []string `yaml:"includes"`
}

func (w *workspaceTarget) Validate(ctx validation.Context) error {
	if err := validation.ValidateManifoldName(w.WorkspaceName); err != nil {
		return err
	}

	for _, include := range w.Dependencies() {
		if err := include.Validate(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (w *workspaceTarget) Name() string {
	return w.WorkspaceName
}

func (workspaceTarget) Kind() TargetKind {
	return WorkspaceTargetKind
}

func (w *workspaceTarget) Dependencies() []Dependency {
	includes := make([]Dependency, len(w.RawIncludes))

	for idx, include := range w.RawIncludes {
		includes[idx] = newInclude(include)
	}

	return includes
}
