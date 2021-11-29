package steps

import (
	"fmt"
	"manifold/internal/config"
	"manifold/internal/errors"
)

type Provider struct {
	options *ProviderOptions
}

func NewProvider(options *ProviderOptions) *Provider {
	provider := new(Provider)
	provider.options = options

	return provider
}

func (provider *Provider) CreateFrom(configStep config.Step) (Step, error) {
	if len(configStep) == 0 {
		return nil, errors.NewError(EmptyStep)
	}

	for name, factory := range provider.options.Factories {
		data := configStep[name]

		if data == nil {
			continue
		}

		step, stepErr := factory(data)

		if stepErr != nil {
			return nil, errors.NewError(StepFailed, name, stepErr)
		}

		return step, nil
	}

	return nil, errors.NewError(StepNotMatched)
}

func (provider *Provider) Execute(step Step) error {
	executor := provider.options.Executors[step.Name()]

	if executor == nil {
		panic(fmt.Sprintf("no executor for %s", step.Name()))
	}

	return executor(step)
}
