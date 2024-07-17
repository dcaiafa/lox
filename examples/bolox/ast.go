package main

import (
	"fmt"
	gotoken "go/token"
)

type Bounds struct {
	Begin gotoken.Pos
	End   gotoken.Pos
}

type AST interface {
	Bounds() Bounds
	SetBounds(b Bounds)
}

type baseAST struct {
	bounds Bounds
}

func (b *baseAST) Bounds() Bounds {
	return b.bounds
}

func (b *baseAST) SetBounds(v Bounds) {
	b.bounds = v
}

type Program struct {
	AST
	Statements []Statement
}

func (p *Program) Run(ctx *Context) error {
	for _, stmt := range p.Statements {
		err := stmt.Run(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

type Statement interface {
	AST
	Run(ctx *Context) error
}

type FuncCallStatement struct {
	AST

	FuncCall *FuncCall
}

func (s *FuncCallStatement) Run(ctx *Context) error {
	_, err := s.FuncCall.Eval(ctx)
	return err
}

type VarAssign struct {
	AST

	VarName string
	Value   Expr
}

func (a *VarAssign) Run(ctx *Context) error {
	v, err := a.Value.Eval(ctx)
	if err != nil {
		return err
	}
	ctx.SetGlobal(a.VarName, v)
	return nil
}

type Expr interface {
	AST
	Eval(ctx *Context) (any, error)
}

type FuncCall struct {
	baseAST

	FuncName string
	Args     []Expr
}

func (c *FuncCall) Eval(ctx *Context) (any, error) {
	var err error
	vals := make([]any, len(c.Args))
	for i, arg := range c.Args {
		vals[i], err = arg.Eval(ctx)
		if err != nil {
			return nil, err
		}
	}
	return ctx.Call(c.FuncName, vals)
}

type VarRef struct {
	baseAST

	VarName string
}

func (r *VarRef) Eval(ctx *Context) (any, error) {
	v, ok := ctx.GetGlobal(r.VarName)
	if !ok {
		return nil, fmt.Errorf("undefined: %v", r.VarName)
	}
	return v, nil
}

func doPrint(vals []any) (any, error) {
	fmt.Println(vals...)
	return nil, nil
}

type Literal struct {
	baseAST

	Val any
}

func (l *Literal) Eval(ctx *Context) (any, error) {
	return l.Val, nil
}
