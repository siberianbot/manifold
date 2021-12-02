package builder

import (
	"manifold/internal/config"
	"manifold/internal/errors"
	"manifold/internal/step"
	stepProvider "manifold/internal/step/provider"
)

const (
	emptyStep         = "step is empty"
	ambiguousStep     = "step is ambiguous"
	unknownStep       = "unknown step \"%s\""
	stepFactoryFailed = "failed to process step \"%s\": %v"
)

type Interface interface {
	FromConfig(definition config.ProjectStep) (step.Step, error)
}

type builder struct {
	provider stepProvider.Interface
}

func NewBuilder(provider stepProvider.Interface) Interface {
	return &builder{provider: provider}
}

func (b *builder) FromConfig(definition config.ProjectStep) (step.Step, error) {
	if len(definition) == 0 {
		return nil, errors.NewError(emptyStep)
	}

	var matchedFactory step.Factory = nil
	matchedKey := ""
	var matchedValue interface{} = nil

	for key, value := range definition {
		factory := b.provider.FactoryFor(key)

		if factory == nil {
			return nil, errors.NewError(unknownStep, key)
		}

		if matchedFactory != nil {
			return nil, errors.NewError(ambiguousStep)
		}

		matchedFactory = factory
		matchedKey = key
		matchedValue = value
	}

	if matchedFactory == nil {
		panic("matchedFactory == nil")
	}

	result, err := matchedFactory.CreateFrom(matchedValue)

	if err != nil {
		return nil, errors.NewError(stepFactoryFailed, matchedKey, err)
	}

	return result, nil
}
