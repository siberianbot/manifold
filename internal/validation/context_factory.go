package validation

type ContextFactory interface {
	New(path string) Context
}

type defaultContextFactory struct {
	//
}

func (d *defaultContextFactory) New(path string) Context {
	ctx := new(defaultContext)
	ctx.path = path

	return ctx
}

func NewDefaultContextFactory() ContextFactory {
	return new(defaultContextFactory)
}
