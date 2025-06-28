package main

import (
	"errors"
	gotoken "go/token"
	"strconv"
	"strings"

	"github.com/dcaiafa/loxlex/simplelexer"

	"github.com/dcaiafa/lox/internal/base/errlogger"
)

func Parse(s string) (any, error) {
	fset := gotoken.NewFileSet()
	file := fset.AddFile("expr", -1, len(s))

	var errStr strings.Builder
	errs := errlogger.New(fset, &errStr)

	var parser jsoncParser
	parser.file = file
	parser.errs = errs

	lex := simplelexer.New(simplelexer.Config{
		StateMachine: new(_LexerStateMachine),
		File:         file,
		Input:        []byte(s),
	})

	_ = parser.parse(lex)
	if errs.HasError() {
		return nil, errors.New(strings.Trim(errStr.String(), "\n"))
	}
	return parser.res, nil
}

type Token = simplelexer.Token

type jsoncParser struct {
	lox
	file *gotoken.File
	errs *errlogger.ErrLogger
	res  any
}

func (p *jsoncParser) on_json(v any) any {
	p.res = v
	return v
}

func (p *jsoncParser) on_value__object(v map[string]any) any {
	return v
}

func (p *jsoncParser) on_value__array(v []any) any {
	return v
}

func (p *jsoncParser) on_value__tok(t Token) any {
	switch t.Type {
	case STRING:
		return unescape(t.Str[1 : len(t.Str)-1])
	case NUMBER:
		v, err := strconv.ParseFloat(string(t.Str), 64)
		if err != nil {
			p.errs.Errorf(t.Pos, "invalid number")
		}
		return v
	case TRUE:
		return true
	case FALSE:
		return false
	case NULL:
		return nil
	default:
		panic("unreachable")
	}
}

type member struct {
	K string
	V any
}

func (p *jsoncParser) on_object(_ Token, members []member, _ Token) map[string]any {
	m := make(map[string]any, len(members))

	for _, member := range members {
		m[member.K] = member.V
	}

	return m
}

func (p *jsoncParser) on_member(k Token, _ Token, v any) member {
	return member{K: unescape(k.Str[1 : len(k.Str)-1]), V: v}
}

func (p *jsoncParser) on_array(_ Token, entries []any, _ Token) []any {
	return entries
}

func unescape(lit []byte) string {
	var str strings.Builder

	for i := 0; i < len(lit); i++ {
		if lit[i] == '\\' {
			switch lit[i+1] {
			case '"', '\\', '/':
				str.WriteByte(lit[i+1])
				i++
			case 'b':
				str.WriteByte('\b')
				i++
			case 'f':
				str.WriteByte('\f')
				i++
			case 'n':
				str.WriteByte('\n')
				i++
			case 'r':
				str.WriteByte('\r')
				i++
			case 't':
				str.WriteByte('\t')
				i++
			case 'u':
				str.WriteRune(hexToRune(string(lit[i+2 : i+6])))
				i += 5
			default:
				panic("unreachable")
			}
		} else {
			str.WriteByte(lit[i])
		}
	}

	return str.String()
}

func hexToRune(str string) rune {
	v, err := strconv.ParseUint(string(str), 16, 32)
	if err != nil {
		panic(err)
	}
	return rune(v)
}
