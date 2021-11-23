package steps

type ProviderOptions struct {
	Factories map[string]Factory
	Executors map[string]Executor
}

func NewProviderOptions() *ProviderOptions {
	options := new(ProviderOptions)
	options.Factories = make(map[string]Factory)
	options.Executors = make(map[string]Executor)

	return options
}
