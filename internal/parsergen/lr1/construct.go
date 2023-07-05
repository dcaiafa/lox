package lr1

import (
	"github.com/dcaiafa/lox/internal/parsergen/grammar"
	"github.com/dcaiafa/lox/internal/util/logger"
)

func ConstructLR(
	g *grammar.AugmentedGrammar,
	logger *logger.Logger,
) *ParserTable {
	pt := NewParserTable(g)

	start := NewItemSet()
	start.Add(NewItem(g, g.Sprime.Prods[0], 0, g.EOF))
	start.Closure(g)
	pt.States.Add(start.LR1Key(), start)

	changed := true
	for changed {
		changed = false
		pt.States.ForEach(func(from *ItemSet) {
			for _, sym := range from.Follow(g) {
				to := from.Goto(g, sym)
				toKey := to.LR1Key()
				existing := pt.States.Get(toKey)
				if existing != nil {
					to = existing
				} else {
					pt.States.Add(toKey, to)
					changed = true
				}
				pt.Transitions.Add(from, to, sym)
			}
		})
	}

	createActions(pt, logger)

	return pt
}

func ConstructLALR(
	g *grammar.AugmentedGrammar,
	log *logger.Logger,
) *ParserTable {
	pt := NewParserTable(g)

	start := NewItemSet()
	start.Add(NewItem(g, g.Sprime.Prods[0], 0, g.EOF))
	start.Closure(g)
	pt.States.Add(start.LR0Key(), start)

	changed := true
	for changed {
		changed = false
		pt.States.ForEach(func(from *ItemSet) {
			for _, sym := range from.Follow(g) {
				to := from.Goto(g, sym)
				toKey := to.LR0Key()
				existing := pt.States.Get(toKey)
				if existing != nil {
					for _, item := range to.GetItems() {
						changed = existing.Add(item) || changed
					}
					to = existing
				} else {
					pt.States.Add(toKey, to)
					changed = true
				}
				pt.Transitions.Add(from, to, sym)
			}
		})
	}

	createActions(pt, log)

	return pt
}

func createActions(pt *ParserTable, logger *logger.Logger) {
	g := pt.Grammar
	pt.States.ForEach(func(s *ItemSet) {
		logger := logger
		if s.Index > 0 {
			logger.Logf("")
		}
		logger.Logf("I%d:", s.Index)
		logger = logger.WithIndent()
		logger.Logf("%v", s.ToString(g))
		logger.Logf("")

		for _, item := range s.GetItems() {
			prod := g.Prods[item.Prod]
			if item.Dot == uint32(len(prod.Terms)) {
				rule := g.ProdRule(prod)
				act := Action{
					Type:   ActionReduce,
					Reduce: rule,
				}
				if rule == g.Sprime {
					act = Action{Type: ActionAccept}
				}
				terminal := g.Terminals[item.Lookahead]
				pt.Actions.Add(s, terminal, act, logger)
				continue
			}
			terminal, ok := g.TermSymbol(prod.Terms[item.Dot]).(*grammar.Terminal)
			if !ok {
				continue
			}
			shiftState := pt.Transitions.Get(s, terminal)
			shiftAction := Action{
				Type:  ActionShift,
				Shift: shiftState,
			}
			pt.Actions.Add(s, terminal, shiftAction, logger)
		}
	})
}
