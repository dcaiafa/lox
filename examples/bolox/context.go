package main

type Context struct {
	Globals map[string]any
}

func NewContext() *Context {
	return &Context{
		Globals: make(map[string]any),
	}
}

func (c *Context) SetGlobal(n string, v any) {
	c.Globals[n] = v
}

func (c *Context) GetGlobal(n string) (any, bool) {
	v, ok := c.Globals[n]
	return v, ok
}
