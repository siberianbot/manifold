package config

import (
	"gopkg.in/yaml.v2"
	"io"
)

func Read(reader io.Reader) (Configuration, error) {
	config := new(configuration)

	decoder := yaml.NewDecoder(reader)

	if yamlErr := decoder.Decode(config); yamlErr != nil {
		return nil, yamlErr
	}

	return config, nil
}
