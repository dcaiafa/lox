package main

import (
	"fmt"
	gotoken "go/token"
	"math/rand"
	"reflect"
	"strconv"
)

type Func func(args []any) (any, error)

type Context struct {
	FileSet *gotoken.FileSet
	Globals map[string]any
}

func NewContext(fs *gotoken.FileSet) *Context {
	c := &Context{
		FileSet: fs,
		Globals: make(map[string]any),
	}

	c.RegisterFunc("print", builtinPrint)
	c.RegisterFunc("rand", builtinRand)
	c.RegisterFunc("prompt", builtinPrompt)
	c.RegisterFunc("parse_int", builtinParseInt)

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

func builtinRand(args []any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("rand expects one argument")
	}
	arg0, ok := args[0].(int)
	if !ok {
		return nil, fmt.Errorf("argument #1 must be an int")
	}
	n := rand.Intn(arg0)
	return int(n), nil
}

func builtinParseInt(args []any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("parse_int expects one argument")
	}
	arg0, ok := args[0].(string)
	if !ok {
		return nil, fmt.Errorf("argument #1 must be an int")
	}
	n, err := strconv.Atoi(arg0)
	if err != nil {
		return nil, nil
	}
	return n, nil
}

func builtinPrompt(args []any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("parse_int expects one argument")
	}
	arg0, ok := args[0].(string)
	if !ok {
		return nil, fmt.Errorf("argument #1 must be an int")
	}
	fmt.Print(arg0 + " ")
	var v string
	fmt.Scanf("%s", &v)
	return v, nil
}
