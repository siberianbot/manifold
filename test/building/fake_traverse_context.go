package building

type FakeTraverseContext struct {
	Errors   []string
	Warnings []string
	File     string
}

func (f *FakeTraverseContext) AddError(message string) {
	f.Errors = append(f.Errors, message)
}

func (f *FakeTraverseContext) AddWarning(message string) {
	f.Warnings = append(f.Warnings, message)
}

func (f FakeTraverseContext) IsValid() bool {
	return len(f.Errors) > 0
}

func (f FakeTraverseContext) CurrentFile() string {
	return f.File
}

func NewFakeTraverseContext() FakeTraverseContext {
	return FakeTraverseContext{
		Errors:   make([]string, 0),
		Warnings: make([]string, 0),
		File:     "",
	}
}
