package steps

import (
	"manifold/internal/config"
)

type StepProvider interface {
	CreateFor(configStep config.Step) (Step, error)
}
