package provider

import (
	"manifold/internal/step"
)

type Options struct {
	Factories map[string]step.Factory
	Executors map[string]step.Executor
}

type Interface interface {
	FactoryFor(name string) step.Factory
	ExecutorFor(name string) step.Executor
}

type provider struct {
	options Options
}

func NewProvider(options Options) Interface {
	return &provider{options: options}
}

func (p *provider) FactoryFor(name string) step.Factory {
	return p.options.Factories[name]
}

func (p *provider) ExecutorFor(name string) step.Executor {
	return p.options.Executors[name]
}
