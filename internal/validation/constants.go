package validation

const (
	NameRegexPattern = "^[a-zA-Z0-9\\-_.]+$"
)

const (
	EmptyManifoldName   = "name is empty"
	InvalidManifoldName = "name \"%s\" does not matches regex pattern \"%s\""
	EmptyPath           = "path is empty"
	InvalidPath         = "invalid path \"%s\""
)
