package ast

import (
	"github.com/dcaiafa/lox/internal/lexergen/mode"
	"github.com/dcaiafa/lox/internal/lexergen/nfa"
	"github.com/dcaiafa/lox/internal/lexergen/rang3"
)

type LexerTermCharClass struct {
	baseAST

	Neg            bool
	CharClassItems []*CharClassItem
}

func (t *LexerTermCharClass) RunPass(ctx *Context, pass Pass) {}

func (t *LexerTermCharClass) NFACons(ctx *Context) *mode.NFAComposite {
	ranges := make([]rang3.Range, len(t.CharClassItems))
	for i, item := range t.CharClassItems {
		ranges[i] = rang3.Range{
			B: item.From,
			E: item.To,
		}
	}
	ranges = rang3.Flatten(ranges)
	if t.Neg {
		ranges = rang3.Negate(ranges)
	}

	nfaFactory := ctx.Mode().StateFactory
	nfaCons := &mode.NFAComposite{}
	nfaCons.B = nfaFactory.NewState()
	nfaCons.E = nfaFactory.NewState()
	for _, r := range ranges {
		rc := mode.NFAComposite{}
		rc.B = nfaFactory.NewState()
		rc.E = nfaFactory.NewState()
		rc.B.AddTransition(rc.E, r)

		nfaCons.B.AddTransition(rc.B, nfa.Epsilon)
		rc.E.AddTransition(nfaCons.E, nfa.Epsilon)
	}
	return nfaCons
}

type CharClassItem struct {
	baseAST
	From rune
	To   rune
}

func (i *CharClassItem) RunPass(ctx *Context, pass Pass) {}
