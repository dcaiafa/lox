package main

import (
	gotoken "go/token"
	"math"
	"strconv"

	"github.com/dcaiafa/loxlex/simplelexer"
)

func Eval(expr string) (float64, error) {
	fset := gotoken.NewFileSet()
	file := fset.AddFile("expr", -1, len(expr))
	errs := &ErrLogger{
		Fset: fset,
	}

	var parser calcParser
	parser.errLogger = errs

	lex := simplelexer.New(simplelexer.Config{
		StateMachine: new(_LexerStateMachine),
		File:         file,
		Input:        []byte(expr),
	})

	_ = parser.parse(lex)
	return parser.result, errs.Err()
}

type Token = simplelexer.Token

type calcParser struct {
	lox
	errLogger *ErrLogger
	result    float64
}

func (p *calcParser) on_S(v float64) any {
	p.result = v
	return nil
}

func (p *calcParser) on_S__error(err Error) any {
	if err.Token.Type == ERROR {
		p.errLogger.Errorf(
			err.Token.Pos,
			"Error: %v",
			err.Token.Err)
	} else {
		p.errLogger.Errorf(
			err.Token.Pos,
			"Error: unexpected token %v",
			_TokenToString(err.Token.Type))
	}
	return nil
}

func (p *calcParser) on_expr__binary(left float64, op Token, right float64) float64 {
	switch op.Type {
	case ADD:
		return left + right
	case SUB:
		return left - right
	case MUL:
		return left * right
	case DIV:
		return left / right
	case REM:
		return math.Mod(left, right)
	case POW:
		return math.Pow(left, right)
	default:
		panic("not reached")
	}
}

func (p *calcParser) on_expr__paren(_ Token, e float64, _ Token) float64 {
	return e
}

func (p *calcParser) on_expr__num(e float64) float64 {
	return e
}

func (p *calcParser) on_num(num Token) float64 {
	v, err := strconv.ParseFloat(string(num.Str), 64)
	if err != nil {
		p.errLogger.Errorf(num.Pos, "invalid float: %v", err)
		return 0
	}
	return v
}

func (p *calcParser) on_num__minus(_ Token, num Token) float64 {
	v, err := strconv.ParseFloat(string(num.Str), 64)
	if err != nil {
		p.errLogger.Errorf(num.Pos, "invalid float: %v", err)
		return 0
	}
	return -v
}
