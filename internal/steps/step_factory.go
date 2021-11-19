package steps

type StepFactory interface {
	Name() string
	Construct(definition interface{}) (Step, error)
}
