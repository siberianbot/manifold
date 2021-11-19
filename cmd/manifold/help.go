package main

import "fmt"

func help() string {
	return fmt.Sprintf(`
Usage: manifold <subcommand> [args]

build <path>    - Build project or workspace. Path should be valid Manifold 
                  configuration. Current directory is used if path is not 
                  defined.

Manifold %s
GitHub: %s
`, Version, Link)
}
