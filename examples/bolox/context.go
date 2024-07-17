package main

import (
	"fmt"
	"reflect"
)

type Func func(args []any) (any, error)

type Context struct {
	Globals map[string]any
}

func NewContext() *Context {
	c := &Context{
		Globals: make(map[string]any),
	}
	c.RegisterFunc("print", builtinPrint)
	return c
}

func (c *Context) RegisterFunc(funcName string, f Func) {
	c.SetGlobal(funcName, f)
}

func (c *Context) SetGlobal(n string, v any) {
	c.Globals[n] = v
}

func (c *Context) GetGlobal(n string) (any, bool) {
	v, ok := c.Globals[n]
	return v, ok
}

func (c *Context) Call(funcName string, args []any) (any, error) {
	v, ok := c.GetGlobal(funcName)
	if !ok {
		return nil, fmt.Errorf(
			"undefined: %v", funcName)
	}
	f, ok := v.(Func)
	if !ok {
		return nil, fmt.Errorf(
			"cannot call %v", reflect.TypeOf(v))
	}
	return f(args)
}

func builtinPrint(args []any) (any, error) {
	fmt.Println(args...)
	return nil, nil
}
