package document_definition

import (
	"gopkg.in/yaml.v2"
	"io"
)

func Read(reader io.Reader) (*DocumentDefinition, error) {
	var document DocumentDefinition

	decoder := yaml.NewDecoder(reader)

	if yamlErr := decoder.Decode(&document); yamlErr != nil {
		return nil, yamlErr
	}

	return &document, nil
}
