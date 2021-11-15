package validation

type Context interface {
	AddError(message string, params ...interface{})
	AddWarning(message string, params ...interface{})

	IsValid() bool
}
