package step

type Step interface {
	Name() string
}

type Factory interface {
	CreateFrom(value interface{}) (Step, error)
}

type ExecutorContext interface {
	Step() Step
	Dir() string
}

type Executor interface {
	Execute(context ExecutorContext) error
}
