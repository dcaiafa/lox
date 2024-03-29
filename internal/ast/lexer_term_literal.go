package ast

import (
	"unicode/utf8"

	"github.com/dcaiafa/lox/internal/lexergen/mode"
	"github.com/dcaiafa/lox/internal/lexergen/rang3"
)

type LexerTermLiteral struct {
	baseAST

	Literal string
}

func (t *LexerTermLiteral) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		if len(t.Literal) == 0 {
			ctx.Errs.Errorf(ctx.Position(t), "literal cannot be empty")
			return
		}

	case Print:
		printer := ctx.CurrentPrinter.Peek()
		printer.Printf("LexerTermLiteral: %q", t.Literal)
	}
}

func (t *LexerTermLiteral) NFACons(ctx *Context) *mode.NFAComposite {
	// For a literal "foo", build the NFACons:
	//     f      o      o
	//  B --> s1 --> s2 --> E
	nfa := ctx.Mode().StateFactory
	nfaCons := new(mode.NFAComposite)
	nfaCons.B = nfa.NewState()
	nfaCons.E = nfaCons.B
	str := t.Literal
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		str = str[size:]
		s := nfa.NewState()
		nfaCons.E.AddTransition(s, rang3.Range{B: r, E: r})
		nfaCons.E = s
	}
	return nfaCons
}
