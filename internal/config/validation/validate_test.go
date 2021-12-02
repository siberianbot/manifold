package validation

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"manifold/internal/config"
	"manifold/internal/validation"
	"os"
	"path/filepath"
	"testing"
)

func TestValidateConfiguration(t *testing.T) {
	t.Run("EmptyConfig", func(t *testing.T) {
		dir := ""
		cfg := config.Configuration{}

		err := validateConfiguration(&cfg, dir)

		assert.EqualError(t, err, EmptyConfiguration)
	})

	t.Run("AmbiguousConfig", func(t *testing.T) {
		dir := ""
		cfg := config.Configuration{
			Project:   &config.Project{},
			Workspace: &config.Workspace{},
		}

		err := validateConfiguration(&cfg, dir)

		assert.EqualError(t, err, AmbiguousConfiguration)
	})

	t.Run("ValidProjectConfig", func(t *testing.T) {
		dir := ""
		cfg := config.Configuration{
			Project: &config.Project{Name: "foo"},
		}

		err := validateConfiguration(&cfg, dir)

		assert.NoError(t, err)
	})

	t.Run("ValidWorkspaceConfig", func(t *testing.T) {
		dir := ""
		cfg := config.Configuration{
			Workspace: &config.Workspace{Name: "foo"},
		}

		err := validateConfiguration(&cfg, dir)

		assert.NoError(t, err)
	})
}

func TestValidateProject(t *testing.T) {
	t.Run("EmptyName", func(t *testing.T) {
		name := ""
		dir := ""
		project := config.Project{Name: name}

		err := validateProject(&project, dir)

		assert.EqualError(t, err, fmt.Sprintf(InvalidProject, validation.EmptyManifoldName))
	})

	t.Run("InvalidName", func(t *testing.T) {
		name := "foo!"
		dir := ""
		project := config.Project{Name: name}

		err := validateProject(&project, dir)

		assert.EqualError(t, err, fmt.Sprintf(InvalidProject, fmt.Sprintf(validation.InvalidManifoldName, name, validation.NameRegexPattern)))
	})
}

func TestValidateWorkspace(t *testing.T) {
	t.Run("EmptyName", func(t *testing.T) {
		name := ""
		dir := ""
		workspace := config.Workspace{Name: name}

		err := validateWorkspace(&workspace, dir)

		assert.EqualError(t, err, fmt.Sprintf(InvalidWorkspace, validation.EmptyManifoldName))
	})

	t.Run("InvalidName", func(t *testing.T) {
		name := "foo!"
		dir := ""
		workspace := config.Workspace{Name: name}

		err := validateWorkspace(&workspace, dir)

		assert.EqualError(t, err, fmt.Sprintf(InvalidWorkspace, fmt.Sprintf(validation.InvalidManifoldName, name, validation.NameRegexPattern)))
	})
}

func TestValidateProjectDependency(t *testing.T) {
	t.Run("EmptyDependency", func(t *testing.T) {
		dir := ""
		dependency := config.ProjectDependency{}

		err := validateProjectDependency(dependency, dir)

		assert.EqualError(t, err, EmptyProjectDependency)
	})

	t.Run("AmbiguousDependency", func(t *testing.T) {
		dir := ""
		dependency := config.ProjectDependency{
			Path:    "foo",
			Project: "bar",
		}

		err := validateProjectDependency(dependency, dir)

		assert.EqualError(t, err, AmbiguousProjectDependency)
	})

	t.Run("InvalidProject", func(t *testing.T) {
		dir := ""
		project := "foo!"
		dependency := config.ProjectDependency{
			Project: project,
		}

		err := validateProjectDependency(dependency, dir)

		assert.EqualError(t, err, fmt.Sprintf(InvalidProjectDependency, fmt.Sprintf(validation.InvalidManifoldName, project, validation.NameRegexPattern)))
	})

	t.Run("ValidProject", func(t *testing.T) {
		dir := ""
		project := "foo"
		dependency := config.ProjectDependency{
			Project: project,
		}

		err := validateProjectDependency(dependency, dir)

		assert.NoError(t, err)
	})

	t.Run("InvalidPath", func(t *testing.T) {
		dir := ""
		path := "foo"
		dependency := config.ProjectDependency{
			Path: path,
		}

		err := validateProjectDependency(dependency, dir)

		assert.EqualError(t, err, fmt.Sprintf(InvalidProjectDependency, fmt.Sprintf(validation.InvalidPath, path)))
	})

	t.Run("ValidPath", func(t *testing.T) {
		dir := t.TempDir()
		dependency := config.ProjectDependency{
			Path: "foo",
		}

		_ = os.Mkdir(filepath.Join(dir, dependency.Path), os.ModeDir)

		err := validateProjectDependency(dependency, dir)

		assert.NoError(t, err)
	})
}

func TestValidateWorkspaceInclude(t *testing.T) {
	t.Run("EmptyInclude", func(t *testing.T) {
		dir := ""
		include := ""

		err := validateInclude(config.WorkspaceInclude(include), dir)

		assert.EqualError(t, err, EmptyWorkspaceInclude)
	})

	t.Run("InvalidInclude", func(t *testing.T) {
		dir := ""
		include := "foo"

		err := validateInclude(config.WorkspaceInclude(include), dir)

		assert.EqualError(t, err, fmt.Sprintf(InvalidWorkspaceInclude, fmt.Sprintf(validation.InvalidPath, include)))
	})

	t.Run("ValidInclude", func(t *testing.T) {
		dir := t.TempDir()
		include := "foo"

		_ = os.Mkdir(filepath.Join(dir, include), os.ModeDir)

		err := validateInclude(config.WorkspaceInclude(include), dir)

		assert.NoError(t, err)
	})
}
