package building

type ValidationContext interface {
	AddError(message string) // TODO: AddError should accept error, introduce error types
	AddWarning(message string)

	IsValid() bool
}

type TraverseContext interface {
	ValidationContext

	CurrentFile() string
}
