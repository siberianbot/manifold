package validation

import "path/filepath"

type Context struct {
	Path string
}

func NewContext(path string) (ctx *Context) {
	ctx = new(Context)
	ctx.Path = path

	return
}

func (ctx *Context) Dir() string {
	return filepath.Dir(ctx.Path)
}
