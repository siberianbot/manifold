package validator

import (
	"manifold/internal/errors"
	"manifold/internal/graph"
	"manifold/internal/graph/node"
	"manifold/internal/graph/validator/cycle_detection"
	"strings"
)

const (
	cyclesDetected = "cycles detected:"
	selfReference  = "\tproject \"%s\" self references"
	cycleDetected  = "\tcycle: %s"
)

type Interface interface {
	Validate(graph graph.Interface) error
}

type validator struct {
	cycleDetectionAlgorithm cycle_detection.CycleDetectionAlgorithm
}

func NewValidator(cycleDetectionAlgorithm cycle_detection.CycleDetectionAlgorithm) Interface {
	return &validator{cycleDetectionAlgorithm: cycleDetectionAlgorithm}
}

func (v *validator) Validate(graph graph.Interface) error {
	cycles := v.cycleDetectionAlgorithm(graph)

	if len(cycles) == 0 {
		return nil
	}

	errs := make([]error, len(cycles))

	for idx, cycle := range cycles {
		errs[idx] = newCycleError(cycle)
	}

	return errors.NewAggregateError(cyclesDetected, errs...)
}

func newCycleError(nodes []node.Node) error {
	if len(nodes) == 1 {
		return errors.NewError(selfReference, nodes[0].Name())
	} else {
		return errors.NewError(cycleDetected, strings.Join(getNames(nodes), " -> "))
	}
}

func getNames(nodes []node.Node) []string {
	cycleNames := make([]string, len(nodes))

	for idx, n := range nodes {
		cycleNames[idx] = n.Name()
	}

	return cycleNames
}
