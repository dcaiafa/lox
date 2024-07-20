package main

import (
	"fmt"
	"strconv"

	gotoken "go/token"

	"github.com/dcaiafa/loxlex/simplelexer"
)

func Parse(filename string, fileData []byte) (*Program, error) {
	fset := gotoken.NewFileSet()
	file := fset.AddFile(filename, -1, len(fileData))

	l := simplelexer.New(simplelexer.Config{
		StateMachine: new(_LexerStateMachine),
		File:         file,
		Input:        fileData,
	})

	p := new(parser)
	ok := p.parse(l)

	if !ok {
		return nil, fmt.Errorf("failed to parse")
	}

	return p.program, nil
}

type Token = simplelexer.Token

type parser struct {
	lox

	program *Program
}

func (p *parser) on_program(stmts []Statement) any {
	p.program = &Program{
		Statements: stmts,
	}
	return nil
}

func (p *parser) on_stmts__1(stmts []Statement, _ Token, stmt Statement) []Statement {
	if stmt != nil {
		stmts = append(stmts, stmt)
	}
	return stmts
}

func (p *parser) on_stmts__2(stmt Statement) []Statement {
	if stmt == nil {
		return nil
	}
	return []Statement{stmt}
}

func (p *parser) on_stmt(stmt Statement) Statement {
	return stmt
}

func (p *parser) on_stmt__empty() Statement {
	return nil
}

func (p *parser) on_func_call_stmt(fc *FuncCall) Statement {
	return &FuncCallStatement{
		FuncCall: fc,
	}
}

func (p *parser) on_var_assign(n Token, _ Token, v Expr) *VarAssign {
	return &VarAssign{
		VarName: string(n.Str),
		Value:   v,
	}
}

func (p *parser) on_func_call(
	n Token,
	_ Token,
	args []Expr,
	_ Token,
) *FuncCall {
	return &FuncCall{
		FuncName: string(n.Str),
		Args:     args,
	}
}

func (p *parser) on_expr__bin(l Expr, optok Token, r Expr) Expr {
	var op Op
	switch optok.Type {
	case PLUS:
		op = OpPlus
	case MINUS:
		op = OpMinus
	case TIMES:
		op = OpTimes
	case DIV:
		op = OpDiv
	default:
		panic("unreachable")
	}
	return &BinaryExpr{
		Left:  l,
		Right: r,
		Op:    op,
	}
}

func (p *parser) on_expr__paren(_ Token, e Expr, _ Token) Expr {
	return e
}

func (p *parser) on_expr__simple(e Expr) Expr {
	return e
}

func (p *parser) on_simple_expr(e Expr) Expr {
	return e
}

func (p *parser) on_var_ref(n Token) *VarRef {
	return &VarRef{
		VarName: string(n.Str),
	}
}

func (p *parser) on_literal(l Token) *Literal {
	switch l.Type {
	case INT:
		n, err := strconv.Atoi(string(l.Str))
		if err != nil {
			panic(err)
		}
		return &Literal{Val: n}

	case STR:
		return &Literal{Val: string(l.Str[1 : len(l.Str)-1])}

	default:
		panic("invalid token type")
	}
}
