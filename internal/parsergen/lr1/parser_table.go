package lr1

import (
	"fmt"
	"io"

	"github.com/dcaiafa/lox/internal/parsergen/grammar"
	"github.com/dcaiafa/lox/internal/util/logger"
)

type ParserTable struct {
	Grammar     *grammar.AugmentedGrammar
	States      *StateSet
	Transitions *TransitionMap
	Actions     *ActionMap
}

func NewParserTable(g *grammar.AugmentedGrammar) *ParserTable {
	return &ParserTable{
		Grammar:     g,
		States:      NewStateSet(),
		Transitions: NewTransitionMap(),
		Actions:     NewActionMap(),
	}
}

func (t *ParserTable) Print(w io.Writer) {
	l := logger.New(w)
	t.States.ForEach(func(s *ItemSet) {
		l := l
		l.Logf("I%d:", s.Index)
		l = l.WithIndent()
		l.Logf("%v", s.ToString(t.Grammar))
		t.Actions.ForEachActionSet(
			t.Grammar, s,
			func(sym grammar.Symbol, actions []Action) {
				l := l.WithIndent()
				conflict := ""
				if len(actions) > 1 {
					conflict = " <== CONFLICT"
				}
				for _, action := range actions {
					l.Logf(
						"on %v: %v%v",
						sym.SymName(), action.ToString(t.Grammar), conflict)
				}
			})
	})
}

func (t *ParserTable) PrintStateGraph(w io.Writer) {
	fmt.Fprintln(w, `digraph G {`)
	fmt.Fprintln(w, `  rankdir="LR";`)
	t.States.ForEach(func(s *ItemSet) {
		fmt.Fprintf(w, "  I%d [label=%q];\n",
			s.Index,
			fmt.Sprintf("I%d\n%v", s.Index, s.ToString(t.Grammar)),
		)
	})
	t.States.ForEach(func(from *ItemSet) {
		t.Transitions.ForEach(from, func(sym grammar.Symbol, to *ItemSet) {
			fmt.Fprintf(w, "  I%d -> I%d [label=%q];\n",
				from.Index,
				to.Index,
				sym.SymName())
		})
	})
	fmt.Fprintln(w, `}`)
}