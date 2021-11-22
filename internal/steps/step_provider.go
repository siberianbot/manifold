package steps

import (
	"manifold/internal/config"
)

type StepProvider interface {
	CreateFrom(configStep config.Step) (Step, error)
}
