package validation

type ValidationContext interface {
	AddError(message string) // TODO: AddError should accept error, introduce error types
	AddWarning(message string)

	IsValid() bool
}
