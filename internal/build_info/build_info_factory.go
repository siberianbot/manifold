package build_info

import (
	"manifold/internal/document_definition"
	steps2 "manifold/internal/steps"
	"manifold/internal/traversing"
)

func FromDocumentDefinition(document *document_definition.DocumentDefinition, ctx traversing.TraverseContext) BuildInfo {
	switch {
	case document.Project != nil:
		return fromProject(document.Project, ctx)

	case document.Workspace != nil:
		return fromWorkspace(document.Workspace, ctx)

	default:
		ctx.AddError("document is empty")
		return nil
	}
}

func fromProject(definition *document_definition.ProjectDefinition, ctx traversing.TraverseContext) BuildInfo {
	if definition.Name == "" {
		ctx.AddError("project does not contain name")
		return nil
	}

	project := ProjectBuildInfo{
		name:  definition.Name,
		path:  ctx.CurrentFile(),
		Steps: make([]steps2.Step, 0),
	}

	if len(definition.Steps) == 0 {
		ctx.AddWarning("project does not contain any steps")
	} else {
		for _, stepDefinition := range definition.Steps {
			step := steps2.FromStepDefinition(&stepDefinition, ctx)

			if step != nil {
				project.Steps = append(project.Steps, step)
			} else {

			}
		}
	}

	return project
}

func fromWorkspace(definition *document_definition.WorkspaceDefinition, ctx traversing.TraverseContext) BuildInfo {
	if definition.Name == "" {
		ctx.AddError("workspace does not contain name")
		return nil
	}

	workspace := WorkspaceBuildInfo{
		name: definition.Name,
		path: ctx.CurrentFile(),
	}

	return workspace
}