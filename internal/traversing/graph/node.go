package graph

import (
	"manifold/internal/traversing/build_info"
)

type Node struct {
	BuildInfo    build_info.BuildInfo
	Dependencies []*Node
}
