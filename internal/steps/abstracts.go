package steps

type Step interface {
	Name() string
}

type Factory func(definition interface{}) (Step, error)

type Executor func(step Step) error
