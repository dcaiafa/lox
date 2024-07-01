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

func (p *Program) Run() error {
	for _, stmt := range p.Statements {
		err := stmt.Run()
		if err != nil {
			return err
		}
	}
	return nil
}

type Statement interface {
	AST
	Run() error
}

type FuncCallStatement struct {
	AST

	FuncCall *FuncCall
}

func (s *FuncCallStatement) Run() error {
	_, err := s.FuncCall.Eval()
	return err
}

type Expr interface {
	AST
	Eval() (any, error)
}

type FuncCall struct {
	baseAST

	FuncName string
	Args     []Expr
}

func (c *FuncCall) Eval() (any, error) {
	var err error
	vals := make([]any, len(c.Args))
	for i, arg := range c.Args {
		vals[i], err = arg.Eval()
		if err != nil {
			return nil, err
		}
	}

	switch c.FuncName {
	case "print":
		return doPrint(vals)

	default:
		return nil, fmt.Errorf(
			"undefined: %v", c.FuncName)
	}
}

func doPrint(vals []any) (any, error) {
	fmt.Println(vals...)
	return nil, nil
}

type Literal struct {
	baseAST

	Val any
}

func (l *Literal) Eval() (any, error) {
	return l.Val, nil
}
