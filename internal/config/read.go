package config

import (
	"gopkg.in/yaml.v2"
	"io"
)

func Read(reader io.Reader) (*ConfigurationDefinition, error) {
	var document ConfigurationDefinition

	decoder := yaml.NewDecoder(reader)

	if yamlErr := decoder.Decode(&document); yamlErr != nil {
		return nil, yamlErr
	}

	return &document, nil
}
