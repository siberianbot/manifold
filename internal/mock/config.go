package mock

import (
	"github.com/stretchr/testify/mock"
	"manifold/internal/config"
)

type ConfigLoader struct {
	mock.Mock
}

func (c *ConfigLoader) FromPath(path string) (*config.Configuration, error) {
	args := c.Called(path)

	var cfg *config.Configuration

	if args.Get(0) == nil {
		cfg = nil
	} else {
		cfg = args.Get(0).(*config.Configuration)
	}

	var err error

	if args.Get(1) == nil {
		err = nil
	} else {
		err = args.Get(1).(error)
	}

	return cfg, err
}
