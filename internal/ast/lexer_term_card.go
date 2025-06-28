package ast

import (
	"github.com/dcaiafa/lox/internal/lexergen/mode"
	"github.com/dcaiafa/lox/internal/lexergen/nfa"
)

type LexerTermCard struct {
	baseAST

	Term LexerTerm
	Card Card
}

func (t *LexerTermCard) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Print:
		printer := ctx.CurrentPrinter.Peek()
		printer.Printf("LexerTermCard:")
		ctx.CurrentPrinter.Push(printer.WithIndent(2))
		defer ctx.CurrentPrinter.Pop()
	}

	t.Term.RunPass(ctx, pass)
}

func (t *LexerTermCard) NFACons(ctx *Context) *mode.NFAComposite {
	nfaFactory := ctx.Mode().StateFactory

	switch t.Card {
	case One:
		return t.Term.NFACons(ctx)

	case ZeroOrOne:
		//    ε               ε
		//  b -> tb -...-> te -> e
		//  |                    ^
		//  +--------------------+
		//            ε
		termCons := t.Term.NFACons(ctx)
		nfaCons := &mode.NFAComposite{
			B: nfaFactory.NewState(),
			E: nfaFactory.NewState(),
		}
		nfaCons.B.AddTransition(nfaCons.E, nfa.Epsilon)
		nfaCons.B.AddTransition(termCons.B, nfa.Epsilon)
		termCons.E.AddTransition(nfaCons.E, nfa.Epsilon)
		return nfaCons

	case ZeroOrMore, ZeroOrMoreNG:
		//             ε
		//        +--------+
		//    ε   v        |  ε
		//  b -> tb -...-> te -> e
		//  |                    ^
		//  +--------------------+
		//            ε
		termCons := t.Term.NFACons(ctx)
		nfaCons := &mode.NFAComposite{
			B: nfaFactory.NewState(),
			E: nfaFactory.NewState(),
		}
		nfaCons.E.NonGreedy = t.Card == ZeroOrMoreNG
		nfaCons.B.AddTransition(nfaCons.E, nfa.Epsilon)
		nfaCons.B.AddTransition(termCons.B, nfa.Epsilon)
		termCons.E.AddTransition(termCons.B, nfa.Epsilon)
		termCons.E.AddTransition(nfaCons.E, nfa.Epsilon)
		return nfaCons

	case OneOrMore, OneOrMoreNG:
		//             ε
		//        +--------+
		//    ε   v        |  ε
		//  b -> tb -...-> te -> e
		//
		termCons := t.Term.NFACons(ctx)
		nfaCons := &mode.NFAComposite{
			B: nfaFactory.NewState(),
			E: nfaFactory.NewState(),
		}
		nfaCons.E.NonGreedy = t.Card == ZeroOrMoreNG
		nfaCons.B.AddTransition(termCons.B, nfa.Epsilon)
		termCons.E.AddTransition(termCons.B, nfa.Epsilon)
		termCons.E.AddTransition(nfaCons.E, nfa.Epsilon)
		return nfaCons

	default:
		panic("not reached")
	}
}
