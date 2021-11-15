package dependents

type DependentInfoKind int8

const (
	DependentProjectInfoKind DependentInfoKind = iota
	DependentPathInfoKind
)

type DependentInfo interface {
	Kind() DependentInfoKind
}
