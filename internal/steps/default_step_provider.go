package steps

import (
	"errors"
	"fmt"
	"manifold/internal/config"
	"manifold/internal/validation"
	"reflect"
	"strings"
)

type defaultStepProvider struct {
	factories []StepFactory
}

func (provider *defaultStepProvider) CreateFor(configStep config.Step) (Step, error) {
	factory, factoryErr := provider.getFactoryFor(configStep)

	if factoryErr != nil {
		return nil, validation.NewError(validation.StepResolveFailed, factoryErr)
	}

	stepName := factory.Name()
	step, constructErr := factory.Construct(configStep[stepName])

	if constructErr != nil {
		return nil, validation.NewError(validation.StepFailed, stepName, constructErr)
	}

	return step, nil
}

func (provider *defaultStepProvider) getFactoryFor(stepDefinition config.Step) (StepFactory, error) {
	if len(stepDefinition) == 0 {
		return nil, errors.New(validation.EmptyStep)
	}

	matchedFactories := make([]StepFactory, 0)
	notMatchedKeys := make([]string, 0)

	for _, keyValue := range reflect.ValueOf(stepDefinition).MapKeys() {
		key := keyValue.String()
		matched := false

		for _, factory := range provider.factories {
			if key != factory.Name() {
				continue
			}

			matchedFactories = append(matchedFactories, factory)
			matched = true
		}

		if !matched {
			notMatchedKeys = append(notMatchedKeys, key)
		}
	}

	switch {
	case len(matchedFactories) == 1:
		return matchedFactories[0], nil

	case len(matchedFactories) == 0:
		var name string

		if len(notMatchedKeys) == 1 {
			name = notMatchedKeys[0]
		} else {
			name = fmt.Sprintf("with keys %s", strings.Join(notMatchedKeys, ", "))
		}

		return nil, errors.New(fmt.Sprintf(validation.StepNotMatchedAnyToolchain, name))

	default:
		return nil, errors.New(validation.StepMatchesManyToolchains)
	}
}

func NewDefaultStepProvider(factories ...StepFactory) StepProvider {
	provider := new(defaultStepProvider)
	provider.factories = factories

	return provider
}
