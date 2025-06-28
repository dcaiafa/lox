package ast

import (
	gotoken "go/token"

	"github.com/dcaiafa/lox/internal/lexergen/mode"
)

type Bounds struct {
	Begin gotoken.Pos
	End   gotoken.Pos
}

type AST interface {
	RunPass(ctx *Context, pass Pass)
	SetBounds(b Bounds)
	Bounds() Bounds
}

type baseAST struct {
	bounds Bounds
}

func (a *baseAST) SetBounds(b Bounds) {
	a.bounds = b
}

func (a *baseAST) Bounds() Bounds {
	return a.bounds
}

type Statement interface {
	AST
	isStatement()
	Discard() bool
}

type baseStatement struct {
	baseAST
}

func (s *baseStatement) isStatement() {}

func (s *baseStatement) Discard() bool { return false }

// DiscardStatement is a dummy Statement that will be discarded by the parser.
// It is used with the *! cardinality which makes the parser call the Discard()
// method.
type DiscardStatement struct{}

func (s *DiscardStatement) SetBounds(b Bounds)              {}
func (s *DiscardStatement) Bounds() Bounds                  { return Bounds{} }
func (n *DiscardStatement) RunPass(ctx *Context, pass Pass) {}
func (n *DiscardStatement) isStatement()                    {}
func (n *DiscardStatement) Discard() bool                   { return true } // Discard me!

// DiscardStatementSingleton is a singleton for DiscardStatement. It works
// because DiscardStatement has no data and will be discarded by the parser.
var DiscardStatementSingleton = &DiscardStatement{}

type LexerTerm interface {
	AST
	NFACons(ctx *Context) *mode.NFAComposite
}

type Card int

const (
	One Card = iota
	ZeroOrOne
	ZeroOrMore
	ZeroOrMoreNG
	OneOrMore
	OneOrMoreNG
)

type Associativity int

const (
	Left  Associativity = 0
	Right Associativity = 1
)

type ProdQualifier struct {
	baseAST
	Precedence    int
	Associativity Associativity
}
