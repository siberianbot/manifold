package building

type ValidationContext interface {
	AddError(message string)
	AddWarning(message string)

	IsValid() bool
}

type TraverseContext interface {
	ValidationContext

	CurrentFile() string
}
