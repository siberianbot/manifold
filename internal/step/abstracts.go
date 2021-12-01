package step

type Step interface {
	Name() string
}

type Factory func(definition interface{}) (Step, error)

type Executor func(step Step, context *ExecutorContext) error

type ExecutorContext struct {
	Dir string
}
