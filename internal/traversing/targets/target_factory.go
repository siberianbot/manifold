package targets

import (
	"manifold/internal/document_definition"
	"manifold/internal/steps"
	"manifold/internal/traversing"
	"manifold/internal/traversing/dependents"
	"manifold/internal/validation"
)

func FromDocumentDefinition(document *document_definition.DocumentDefinition, ctx traversing.Context) (Target, []dependents.DependentInfo) {
	switch {
	case document.Project != nil && document.Workspace != nil:
		ctx.AddError(validation.DocumentWithBothProjectAndWorkspace)
		return nil, make([]dependents.DependentInfo, 0)

	case document.Project != nil:
		return fromProject(document.Project, ctx)

	case document.Workspace != nil:
		return fromWorkspace(document.Workspace, ctx)

	default:
		ctx.AddError(validation.EmptyDocument)
		return nil, make([]dependents.DependentInfo, 0)
	}
}

func fromProject(definition *document_definition.ProjectDefinition, ctx traversing.Context) (Target, []dependents.DependentInfo) {
	if err := validation.ValidateManifoldName(definition.Name); err != nil {
		ctx.AddError(validation.InvalidProject, err)
		return nil, make([]dependents.DependentInfo, 0)
	}

	project := ProjectTarget{
		name:  definition.Name,
		path:  ctx.CurrentFile(),
		Steps: make([]steps.Step, 0),
	}

	if len(definition.Steps) == 0 {
		ctx.AddWarning(validation.ProjectWithoutSteps)
	} else {
		for _, stepDefinition := range definition.Steps {
			step := steps.FromStepDefinition(stepDefinition, ctx)

			if step != nil {
				project.Steps = append(project.Steps, step)
			}
		}
	}

	dependencies := make([]dependents.DependentInfo, 0)

	for _, dependencyDefinition := range definition.Dependencies {
		dependency := dependents.FromDependencyDefinition(dependencyDefinition, ctx)
		dependencies = append(dependencies, dependency)
	}

	return project, dependencies
}

func fromWorkspace(definition *document_definition.WorkspaceDefinition, ctx traversing.Context) (Target, []dependents.DependentInfo) {
	if err := validation.ValidateManifoldName(definition.Name); err != nil {
		ctx.AddError(validation.InvalidWorkspace, err)
		return nil, make([]dependents.DependentInfo, 0)
	}

	workspace := WorkspaceTarget{
		name: definition.Name,
		path: ctx.CurrentFile(),
	}

	dependencies := make([]dependents.DependentInfo, 0)

	for _, includeDefinition := range definition.Includes {
		dependency := dependents.FromIncludeDefinition(includeDefinition, ctx)

		dependencies = append(dependencies, dependency)
	}

	return workspace, dependencies
}
