package config

type ValidationContext interface {
	Dir() string
}

type Validatable interface {
	Validate(ctx ValidationContext) error
}
