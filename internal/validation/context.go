package validation

import "path/filepath"

type Context interface {
	Dir() string
}

type defaultContext struct {
	path string
}

func (d *defaultContext) Dir() string {
	return filepath.Dir(d.path)
}
