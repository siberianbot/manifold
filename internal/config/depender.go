package config

type Depender interface {
	Dependencies() []Dependency
}
