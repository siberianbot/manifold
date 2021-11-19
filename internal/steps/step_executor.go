package steps

type StepExecutor interface {
	Execute(step Step) error
}
