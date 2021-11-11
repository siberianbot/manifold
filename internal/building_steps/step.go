package building_steps

type StepKind int32

const (
	UnknownStepKind StepKind = iota // TODO: Is UnknownStepKind required?
	CommandStepKind
)

type Step interface {
	Kind() StepKind
	Execute() (bool, error)
}
