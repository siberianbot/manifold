package config

import (
	"bytes"
	"math/rand"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	t.Run("EmptyFile", func(t *testing.T) {
		reader := strings.NewReader("")
		configuration, err := Read(reader)

		if configuration != nil {
			t.Error("configuration is not nil")
		}

		if err == nil {
			t.Error("error is nil")
		}
	})

	t.Run("RandomFile", func(t *testing.T) {
		data := make([]byte, 1024)
		rand.Read(data)

		reader := bytes.NewReader(data)
		configuration, err := Read(reader)

		if configuration != nil {
			t.Error("configuration is not nil")
		}

		if err == nil {
			t.Error("error is nil")
		}
	})

	t.Run("Project", testReadProject)
	t.Run("Workspace", testReadWorkspace)
}

func testReadProject(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		content := `project:`

		reader := strings.NewReader(content)
		configuration, err := Read(reader)

		if err != nil {
			t.Errorf("error is %v, not nil", err)
		}

		if configuration == nil {
			t.Error("configuration is nil")
		} else {
			if configuration.ProjectTarget != nil {
				t.Error("configuration.ProjectTarget is not nil")
			}

			if configuration.WorkspaceTarget != nil {
				t.Error("configuration.WorkspaceTarget is not nil")
			}
		}
	})

	t.Run("Full", func(t *testing.T) {
		content := `
project:
  name: foo
  dependencies:
  - project: bar
  - path: baz
  steps:
  - foo: foo
  - bar: bar
  - baz: baz
`

		reader := strings.NewReader(content)
		configuration, err := Read(reader)

		if err != nil {
			t.Errorf("error is %v, not nil", err)
		}

		if configuration == nil {
			t.Error("configuration is nil")
		} else {
			if configuration.ProjectTarget == nil {
				t.Error("configuration.ProjectTarget is nil")
			} else {
				if configuration.ProjectTarget.Name != "foo" {
					t.Errorf("configuration.ProjectTarget.Name is %s, not %s", configuration.ProjectTarget.Name, "foo")
				}

				if configuration.ProjectTarget.Steps == nil {
					t.Error("configuration.ProjectTarget.Steps is nil")
				} else if len(configuration.ProjectTarget.Steps) != 3 {
					t.Errorf("len(configuration.ProjectTarget.Steps) is %v, not %v", len(configuration.ProjectTarget.Steps), 3)
				} else {
					if !containsNamedStep(configuration.ProjectTarget.Steps, "foo") {
						t.Error("configuration.ProjectTarget.Steps doesn't contains foo")
					}

					if !containsNamedStep(configuration.ProjectTarget.Steps, "bar") {
						t.Error("configuration.ProjectTarget.Steps doesn't contains bar")
					}

					if !containsNamedStep(configuration.ProjectTarget.Steps, "baz") {
						t.Error("configuration.ProjectTarget.Steps doesn't contains baz")
					}
				}

				if configuration.ProjectTarget.ProjectDependencies == nil {
					t.Error("configuration.ProjectTarget.ProjectDependencies is nil")
				} else if len(configuration.ProjectTarget.ProjectDependencies) != 2 {
					t.Errorf("len(configuration.ProjectTarget.ProjectDependencies) is %v, not %v", len(configuration.ProjectTarget.ProjectDependencies), 2)
				} else {
					if !containsProjectDependency(configuration.ProjectTarget.ProjectDependencies, "", "bar") {
						t.Error("configuration.ProjectTarget.ProjectDependencies doesn't contains project bar")
					}

					if !containsProjectDependency(configuration.ProjectTarget.ProjectDependencies, "baz", "") {
						t.Error("configuration.ProjectTarget.ProjectDependencies doesn't contains path baz")
					}
				}
			}
		}
	})
}

func testReadWorkspace(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		content := `workspace:`

		reader := strings.NewReader(content)
		configuration, err := Read(reader)

		if err != nil {
			t.Errorf("error is %v, not nil", err)
		}

		if configuration == nil {
			t.Error("configuration is nil")
		} else {
			if configuration.ProjectTarget != nil {
				t.Error("configuration.ProjectTarget is not nil")
			}

			if configuration.WorkspaceTarget != nil {
				t.Error("configuration.WorkspaceTarget is not nil")
			}
		}
	})

	t.Run("Full", func(t *testing.T) {
		content := `
workspace:
  name: foo
  includes:
  - bar
  - baz
`

		reader := strings.NewReader(content)
		configuration, err := Read(reader)

		if err != nil {
			t.Errorf("error is %v, not nil", err)
		}

		if configuration == nil {
			t.Error("configuration is nil")
		} else {
			if configuration.WorkspaceTarget == nil {
				t.Error("configuration.WorkspaceTarget is nil")
			} else {
				if configuration.WorkspaceTarget.Name != "foo" {
					t.Errorf("configuration.WorkspaceTarget.Name is %s, not %s", configuration.WorkspaceTarget.Name, "foo")
				}

				if configuration.WorkspaceTarget.Includes == nil {
					t.Error("configuration.WorkspaceTarget.Includes is nil")
				} else if len(configuration.WorkspaceTarget.Includes) != 2 {
					t.Errorf("len(configuration.WorkspaceTarget.Includes) is %v, not %v", len(configuration.WorkspaceTarget.Includes), 2)
				} else {
					if !containsInclude(configuration.WorkspaceTarget.Includes, "bar") {
						t.Error("configuration.WorkspaceTarget.Includes doesn't contains bar")
					}

					if !containsInclude(configuration.WorkspaceTarget.Includes, "baz") {
						t.Error("configuration.WorkspaceTarget.Includes doesn't contains baz")
					}
				}
			}
		}
	})
}

func containsNamedStep(steps []Step, name string) bool {
	for _, step := range steps {
		if step[name] == name {
			return true
		}
	}

	return false
}

func containsInclude(includes []string, include string) bool {
	for _, inc := range includes {
		if inc == include {
			return true
		}
	}

	return false
}

func containsProjectDependency(projectDependencies []ProjectDependency, path string, project string) bool {
	for _, dependency := range projectDependencies {
		if dependency.Path == path && dependency.Project == project {
			return true
		}
	}

	return false
}
