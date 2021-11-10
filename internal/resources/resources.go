package resources

import (
	"embed"
	"fmt"
)

//go:embed manifold.schema.json
var resourcesFS embed.FS

func getFile(name string) []byte {
	file, err := resourcesFS.ReadFile(name)

	if err != nil {
		panic(fmt.Sprintf("failed to load embedded resource %s: %v", name, err))
	}

	return file
}

func GetSchema() []byte {
	return getFile("manifold.schema.json")
}
