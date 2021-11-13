package traversing

import "manifold/internal/validation"

type TraverseContext interface {
	validation.ValidationContext

	CurrentFile() string
}
