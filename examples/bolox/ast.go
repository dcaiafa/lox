package main

import (
	"fmt"
	gotoken "go/token"
	"reflect"
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

type Op int

const (
	OpPlus Op = iota
	OpMinus
	OpTimes
	OpDiv
)

type BinaryExpr struct {
	baseAST

	Left  Expr
	Right Expr
	Op    Op
}

func (e *BinaryExpr) Eval(ctx *Context) (any, error) {
	va, err := e.Left.Eval(ctx)
	if err != nil {
		return nil, err
	}
	vb, err := e.Right.Eval(ctx)
	if err != nil {
		return nil, err
	}

	if a, b, ok := castBinaryExpr[int](va, vb); ok {
		switch e.Op {
		case OpPlus:
			return a + b, nil
		case OpMinus:
			return a - b, nil
		case OpTimes:
			return a * b, nil
		case OpDiv:
			if b == 0 {
				return nil, fmt.Errorf("division by zero")
			}
			return a / b, nil
		default:
			panic("unreachable")
		}
	} else if a, b, ok := castBinaryExpr[string](va, vb); ok {
		if e.Op != OpPlus {
			return nil, fmt.Errorf("operation not supported by string")
		}
		return a + b, nil
	} else {
		return nil, fmt.Errorf(
			"operation not supported between %v and %v",
			reflect.TypeOf(va), reflect.TypeOf(vb))
	}
}

func castBinaryExpr[T any](a, b any) (ca, cb T, ok bool) {
	ca, ok = a.(T)
	if !ok {
		return ca, cb, false
	}
	cb, ok = b.(T)
	if !ok {
		return ca, cb, false
	}
	return ca, cb, true
}
