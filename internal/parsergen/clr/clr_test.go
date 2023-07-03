package clr

import (
	"os"
	"strings"
	"testing"

	"github.com/dcaiafa/lox/internal/parsergen/grammar"
	"github.com/dcaiafa/lox/internal/util/logger"
)

func TestCLR(t *testing.T) {
	sg := &grammar.Grammar{
		Terminals: []*grammar.Terminal{
			{Name: "c"},
			{Name: "d"},
		},
		Rules: []*grammar.Rule{
			{
				Name: "S",
				Prods: []*grammar.Prod{
					{Terms: []*grammar.Term{{Name: "C"}, {Name: "C"}}},
				},
			},
			{
				Name: "C",
				Prods: []*grammar.Prod{
					{Terms: []*grammar.Term{{Name: "c"}, {Name: "C"}}},
					{Terms: []*grammar.Term{{Name: "d"}}},
				},
			},
		},
	}

	g, err := sg.ToAugmentedGrammar()
	if err != nil {
		t.Fatalf("ToAugmentedGrammar failed: %v", err)
	}

	parserTable := constructCLR(g, logger.New(os.Stdout))

	var graph strings.Builder
	parserTable.PrintStateGraph(&graph)

	expected := `
digraph G {
  I0 [label="I0\nS' -> .S, $\nS -> .C C, $\nC -> .c C, c\nC -> .c C, d\nC -> .d, c\nC -> .d, d"];
  I1 [label="I1\nS -> C .C, $\nC -> .c C, $\nC -> .d, $"];
  I2 [label="I2\nS' -> S., $"];
  I3 [label="I3\nC -> .c C, c\nC -> .c C, d\nC -> c .C, c\nC -> c .C, d\nC -> .d, c\nC -> .d, d"];
  I4 [label="I4\nC -> d., c\nC -> d., d"];
  I5 [label="I5\nS -> C C., $"];
  I6 [label="I6\nC -> .c C, $\nC -> c .C, $\nC -> .d, $"];
  I7 [label="I7\nC -> d., $"];
  I8 [label="I8\nC -> c C., c\nC -> c C., d"];
  I9 [label="I9\nC -> c C., $"];
  I0 -> I1 [label="C"];
  I0 -> I2 [label="S"];
  I0 -> I3 [label="c"];
  I0 -> I4 [label="d"];
  I1 -> I5 [label="C"];
  I1 -> I6 [label="c"];
  I1 -> I7 [label="d"];
  I3 -> I8 [label="C"];
  I3 -> I3 [label="c"];
  I3 -> I4 [label="d"];
  I6 -> I9 [label="C"];
  I6 -> I6 [label="c"];
  I6 -> I7 [label="d"];
}
`
	expected = strings.TrimSpace(expected)
	actual := strings.TrimSpace(graph.String())

	if expected != actual {
		t.Log("Expected:\n", expected)
		t.Log("Actual:\n", actual)
		t.Fatal("State graph does not match expectation")
	}
}