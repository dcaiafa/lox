package ast

import (
	"fmt"

	"github.com/dcaiafa/lox/internal/base/assert"
	"github.com/dcaiafa/lox/internal/parsergen/lr1"
)

type ParserTermType int

const (
	ParserTermSimple      ParserTermType = iota
	ParserTermZeroOrMore                 // *
	ParserTermZeroOrMoreF                // *!
	ParserTermOneOrMore                  // +
	ParserTermOneOrMoreF                 // +!
	ParserTermZeroOrOne                  // ?
	ParserTermList                       // @list
	ParserTermListOpt                    // @list?
	ParserTermError                      // @error
)

func (t ParserTermType) String() string {
	switch t {
	case ParserTermSimple:
		return "Simple"
	case ParserTermZeroOrMore:
		return "ZeroOrMore"
	case ParserTermOneOrMore:
		return "OneOrMore"
	case ParserTermZeroOrOne: // ?
		return "ZeroOrOne"
	case ParserTermList: // @list
		return "List"
	case ParserTermListOpt: // @list?
		return "ListOpt"
	case ParserTermError: // @error
		return "Error"
	default:
		return "???"
	}
}

type ParserTerm struct {
	baseAST
	Type  ParserTermType
	Name  string
	Alias string
	Child *ParserTerm
	Sep   *ParserTerm

	Symbol lr1.Term
}

func (t *ParserTerm) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		if !t.preCheck(ctx) {
			return
		}
	case Normalize:
		t.normalize(ctx)

	case Print:
		printer := ctx.CurrentPrinter.Peek()
		alias := ""
		if t.Alias != "" {
			alias = fmt.Sprintf("Alias: %v ", t.Alias)
		}
		printer.Printf(
			"Term: Name: %v%v Type: %v",
			t.Name, alias, t.Type)
		ctx.CurrentPrinter.Push(printer.WithIndent(2))
		defer ctx.CurrentPrinter.Pop()
	}

	if t.Child != nil {
		t.Child.RunPass(ctx, pass)
	}
	if t.Sep != nil {
		t.Sep.RunPass(ctx, pass)
	}

	switch pass {
	case Check:
		if !t.postCheck(ctx) {
			return
		}
	}
}

func (t *ParserTerm) preCheck(ctx *Context) bool {
	switch {
	case t.Name != "":
		ast := ctx.Lookup(t.Name)
		if ast == nil {
			ctx.Errs.Errorf(ctx.Position(t), "undefined: %v", t.Name)
			return false
		}
		switch ast := ast.(type) {
		case *ParserRule:
			t.Symbol = ast.Rule
		case *TokenRule:
			t.Symbol = ast.Terminal
		default:
			ctx.Errs.Errorf(ctx.Position(t), "%v is not a parser or token rule", t.Name)
			return false
		}

	case t.Alias != "":
		ast := ctx.LookupAlias(t.Alias)
		switch ast {
		case nil:
			ctx.Errs.Errorf(ctx.Position(t), "unknown token literal: '%v'", t.Alias)
			return false
		case AmbiguousAlias:
			ctx.Errs.Errorf(ctx.Position(t), "ambiguous token literal: '%v'", t.Alias)
			return false
		}
		t.Symbol = ast.Terminal

	case t.Type == ParserTermError:
		t.Symbol = ctx.Grammar.ErrorTerminal
	}

	return true
}

func (t *ParserTerm) postCheck(ctx *Context) bool {
	if t.Type == ParserTermList || t.Type == ParserTermListOpt {
		if t.Child.Type != ParserTermSimple {
			ctx.Errs.Errorf(
				ctx.Position(t.Child),
				"@list entry param must be a simple token or rule")
			return false
		}

		if t.Sep.Type != ParserTermSimple {
			ctx.Errs.Errorf(
				ctx.Position(t.Child),
				"@list separator param must be a simple token or rule")
			return false
		}
	}
	return true
}

func (t *ParserTerm) normalize(ctx *Context) {
	generate := func(name string, f func(r *ParserRule)) {
		var r *ParserRule
		existingRule := ctx.Lookup(name)
		if existingRule == nil {
			r = &ParserRule{Name: name}

			f(r)
			unit := ctx.CurrentUnit.Peek()
			unit.Statements = append(unit.Statements, r)

			r.RunPass(ctx, CreateNames)
			r.RunPass(ctx, Check)
			r.RunPass(ctx, Normalize)
			t.RunPass(ctx, Check)
			assert.False(ctx.Errs.HasError())
		} else {
			r = existingRule.(*ParserRule)
		}

		t.Name = r.Name
		t.Type = ParserTermSimple
		t.Child = nil
		t.Sep = nil
		t.Symbol = r.Rule
	}

	switch t.Type {
	case ParserTermSimple, ParserTermError:
		// No changes required.

	case ParserTermZeroOrMore:
		// a = b c*
		//   =>
		// a = b a'
		// a' = c+ | ε
		generate(t.Child.Symbol.TermName()+"*", func(r *ParserRule) {
			r.Prods = []*ParserProd{
				{
					Terms: []*ParserTerm{
						{Type: ParserTermOneOrMore, Child: t.Child},
					},
				},
				{},
			}
		})

	case ParserTermZeroOrMoreF:
		// a = b c*
		//   =>
		// a = b a'
		// a' = c+ | ε
		generate(t.Child.Symbol.TermName()+"*!", func(r *ParserRule) {
			r.Prods = []*ParserProd{
				{
					Terms: []*ParserTerm{
						{Type: ParserTermOneOrMoreF, Child: t.Child},
					},
				},
				{},
			}
		})

	case ParserTermOneOrMore:
		// a = b c+
		//   =>
		// a = b a'
		// a' = a' c
		//    | c
		generate(t.Child.Symbol.TermName()+"+", func(r *ParserRule) {
			r.Prods = []*ParserProd{
				{
					Terms: []*ParserTerm{
						{Name: r.Name},
						t.Child,
					},
				},
				{
					Terms: []*ParserTerm{
						t.Child,
					},
				},
			}
		})

	case ParserTermOneOrMoreF:
		// a = b c+
		//   =>
		// a = b a'
		// a' = a' c
		//    | c
		generate(t.Child.Symbol.TermName()+"+!", func(r *ParserRule) {
			r.Prods = []*ParserProd{
				{
					Terms: []*ParserTerm{
						{Name: r.Name},
						t.Child,
					},
				},
				{
					Terms: []*ParserTerm{
						t.Child,
					},
				},
			}
		})

	case ParserTermZeroOrOne:
		// a = b c?
		//   =>
		// a = b a'
		// a' = c | ε
		generate(t.Child.Symbol.TermName()+"?", func(r *ParserRule) {
			r.Prods = []*ParserProd{
				{
					Terms: []*ParserTerm{
						t.Child,
					},
				},
				{},
			}
		})

	case ParserTermListOpt:
		// a = b @list(x, sep)?
		//   =>
		// a = b a'
		// a' = @list(x, sep) | ε
		name := fmt.Sprintf(
			"@list(%v,%v)?",
			t.Child.Symbol.TermName(),
			t.Sep.Symbol.TermName())
		generate(name, func(r *ParserRule) {
			r.Prods = []*ParserProd{
				{
					Terms: []*ParserTerm{
						{Type: ParserTermList, Child: t.Child, Sep: t.Sep},
					},
				},
				{ /* ε */ },
			}
		})

	case ParserTermList:
		// a = b @list(c, sep)
		//   =>
		// a = b a'
		// a' = a' sep c
		//    | c
		name := fmt.Sprintf(
			"@list(%v,%v)",
			t.Child.Symbol.TermName(),
			t.Sep.Symbol.TermName())
		generate(name, func(r *ParserRule) {
			r.Prods = []*ParserProd{
				{
					Terms: []*ParserTerm{
						{Name: r.Name},
						t.Sep,
						t.Child,
					},
				},
				{
					Terms: []*ParserTerm{
						t.Child,
					},
				},
			}
		})
	}
}
