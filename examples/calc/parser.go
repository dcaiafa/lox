package main

import (
	gotoken "go/token"
	"math"
	"strconv"
)

func Eval(expr string) (float64, error) {
	fset := gotoken.NewFileSet()
	file := fset.AddFile("expr", -1, len(expr))
	errLogger := &ErrLogger{
		Fset: fset,
	}

	var parser calcParser
	parser.errLogger = errLogger
	lex := newLex(file, []byte(expr), errLogger)
	_ = parser.parse(lex)
	return parser.result, errLogger.Err()
}

type calcParser struct {
	lox
	errLogger *ErrLogger
	result    float64
}

func (p *calcParser) on_S__foo(e float64) any {
	p.result = e
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
	v, err := strconv.ParseFloat(num.Str, 64)
	if err != nil {
		p.errLogger.Errorf(num.Pos, "invalid float: %v", err)
		return 0
	}
	return v
}

func (p *calcParser) on_num__minus(_ Token, num Token) float64 {
	v, err := strconv.ParseFloat(num.Str, 64)
	if err != nil {
		p.errLogger.Errorf(num.Pos, "invalid float: %v", err)
		return 0
	}
	return -v
}

func (p *calcParser) onError() {
	p.errLogger.Errorf(p.errorToken().Pos, "unexpected token %v", p.errorToken())
}
