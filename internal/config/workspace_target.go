package config

import "manifold/internal/validation"

type WorkspaceTarget interface {
	Target

	Includes() []Include
}

type workspaceTarget struct {
	WorkspaceName string   `yaml:"name"`
	RawIncludes   []string `yaml:"includes"`
}

func (w *workspaceTarget) Validate(ctx validation.Context) error {
	if err := validation.ValidateManifoldName(w.WorkspaceName); err != nil {
		return err
	}

	for _, inc := range w.Includes() {
		if err := inc.Validate(ctx); err != nil {
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

func (w *workspaceTarget) Includes() []Include {
	includes := make([]Include, len(w.RawIncludes))

	for idx := 0; idx < len(w.RawIncludes); idx++ {
		includes[idx] = newInclude(w.RawIncludes[idx])
	}

	return includes
}
