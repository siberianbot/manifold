package steps

type Step interface {
	Execute() error
}
