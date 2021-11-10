package document

import (
	"gopkg.in/yaml.v2"
	"io"
)

func Read(reader io.Reader) (*Document, error) {
	var document Document

	decoder := yaml.NewDecoder(reader)

	if yamlErr := decoder.Decode(&document); yamlErr != nil {
		return nil, yamlErr
	}

	return &document, nil
}
