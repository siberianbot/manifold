package steps

type testStep struct {
	name string
}

func newTestStep(name string) Step {
	return &testStep{name: name}
}

func (t testStep) Name() string {
	return t.name
}
