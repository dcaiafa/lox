package main

import (
	"fmt"
	"os"
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

	p := &parser{
		file: file,
	}
	ok := p.parse(l)

	if !ok || p.program == nil {
		return nil, fmt.Errorf("failed to parse")
	}

	return p.program, nil
}

type Token = simplelexer.Token

type parser struct {
	lox
	file    *gotoken.File
	program *Program
}

func (p *parser) on_program(block *Block) any {
	p.program = &Program{
		Block: block,
	}
	return nil
}

func (p *parser) on_program__error(e Error) any {
	if e.Token.Type == ERROR {
		fmt.Fprintf(
			os.Stderr, "%v: %v\n",
			p.file.Position(e.Token.Pos), e.Token.Err)
	} else {
		fmt.Fprintf(
			os.Stderr, "%v: unexpected %v\n",
			p.file.Position(e.Token.Pos), _TokenToString(e.Token.Type))
	}
	return nil
}

func (p *parser) on_block(stmts []Statement) *Block {
	return &Block{
		Statements: stmts,
	}
}

func (p *parser) on_stmt(stmt Statement, _ Token) Statement {
	return stmt
}

func (p *parser) on_stmt__kw(kw Token, _ Token) Statement {
	switch kw.Type {
	case CONTINUE:
		return &ContinueStatement{}
	default:
		panic("unreachable")
	}
}

func (p *parser) on_stmt__nl(_ Token) Statement {
	return &Noop{}
}

func (p *parser) on_while_stmt(_ Token, e Expr, _ Token, b *Block, _ Token) *While {
	return &While{
		Pred:  e,
		Block: b,
	}
}

func (p *parser) on_func_call_stmt(fc *FuncCall) Statement {
	return &FuncCallStatement{
		FuncCall: fc,
	}
}

func (p *parser) on_if_stmt(_ Token, pred Expr, _ Token, b *Block, _ Token, elifs []*Elif, els *Else) *IfStatement {
	return &IfStatement{
		Pred:  pred,
		Block: b,
		Elifs: elifs,
		Else:  els,
	}
}

func (p *parser) on_elif(_ Token, pred Expr, _ Token, b *Block, _ Token) *Elif {
	return &Elif{
		Pred:  pred,
		Block: b,
	}
}

func (p *parser) on_else(_, _ Token, b *Block, _ Token) *Else {
	return &Else{
		Block: b,
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
	case LT:
		op = OpLT
	case LE:
		op = OpLE
	case GT:
		op = OpGT
	case GE:
		op = OpGE
	case EQ:
		op = OpEq
	case AND:
		op = OpAnd
	case OR:
		op = OpOr
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

func (p *parser) on_literal__tok(l Token) Expr {
	switch l.Type {
	case INT:
		n, err := strconv.Atoi(string(l.Str))
		if err != nil {
			panic(err)
		}
		return &Literal{Val: n}

	case TRUE:
		return &Literal{Val: true}

	case FALSE:
		return &Literal{Val: false}

	case NIL:
		return &Literal{Val: nil}

	default:
		panic("invalid token type")
	}
}

func (p *parser) on_literal__string(str Expr) Expr {
	return str
}

func (p *parser) on_string(_ Token, parts []Expr, _ Token) Expr {
	return &String{
		Parts: parts,
	}
}

func (p *parser) on_string_part__char_seq(cs Token) Expr {
	return &StringCharSeq{
		Seq: string(cs.Str),
	}
}

func (p *parser) on_string_part__expr(_ Token, e Expr, _ Token) Expr {
	return e
}
