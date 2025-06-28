package dfa

import (
	"reflect"
	"strings"
	"testing"

	"github.com/dcaiafa/lox/internal/base/set"
	"github.com/dcaiafa/lox/internal/lexergen/nfa"
	"github.com/dcaiafa/lox/internal/lexergen/rang3"
	"github.com/dcaiafa/lox/internal/testutil"
)

func requireEqual[T any](t *testing.T, actual, expected T) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("not equal:\nexpected:\n%v\nactual:\n%v", expected, actual)
	}
}

// nfa1 produces the following NFA:
//
//	                  ε
//	   +---------------------------------+
//	   |                                 |
//		 | ε     ε     a         ε         |
//		s0 -> s1 -> s2 -> s3 --------+     |
//		      ^ \                    |     |
//		      |  \   ε      b     ε  v  ε  v  a     b     b
//		      |   \----> s4 -> s5 -> s6 -> s7 -> s8 -> s9 -> ((s10))
//	        |          ε           |
//	        +----------------------+
//
// Recognizes (a|b)*abb
// Based on Dragon Book Fig. 3.27.
func nfa1() *nfa.State {
	n := nfa.NewStateFactory()

	s0 := n.NewState()
	s1 := n.NewState()
	s2 := n.NewState()
	s3 := n.NewState()
	s4 := n.NewState()
	s5 := n.NewState()
	s6 := n.NewState()
	s7 := n.NewState()
	s8 := n.NewState()
	s9 := n.NewState()
	s10 := n.NewState()

	s0.AddTransition(s1, nfa.Epsilon)
	s1.AddTransition(s2, nfa.Epsilon)
	s2.AddTransition(s3, "a")
	s3.AddTransition(s6, nfa.Epsilon)
	s1.AddTransition(s4, nfa.Epsilon)
	s4.AddTransition(s5, "b")
	s5.AddTransition(s6, nfa.Epsilon)
	s6.AddTransition(s7, nfa.Epsilon)
	s6.AddTransition(s1, nfa.Epsilon)
	s7.AddTransition(s8, "a")
	s8.AddTransition(s9, "b")
	s9.AddTransition(s10, "b")
	s0.AddTransition(s7, nfa.Epsilon)
	s10.Accept = true

	return s0
}

func TestEClosure(t *testing.T) {
	ids := func(nfaStates []*nfa.State) []uint32 {
		ids := make([]uint32, len(nfaStates))
		for i, s := range nfaStates {
			ids[i] = s.ID
		}
		return ids
	}

	c := eClosure(set.New[*nfa.State](nfa1()))
	requireEqual(t, []uint32{0, 1, 2, 4, 7}, ids(c.NFAStates))
}

func TestNFAToDFA(t *testing.T) {
	n := nfa1()
	d := NFAToDFA(n)

	str := &strings.Builder{}
	d.Print(str)

	requireEqual(t, strings.TrimSpace(str.String()), strings.TrimSpace(`
digraph G {
  rankdir="LR";
  0 -> 0 [label="b"];
  0 -> 3 [label="a"];
  1 -> 0 [label="b"];
  1 -> 3 [label="a"];
  2 -> 1 [label="b"];
  2 -> 3 [label="a"];
  3 -> 2 [label="b"];
  3 -> 3 [label="a"];
  0 [label="0", shape="circle"];
  1 [label="1", shape="doublecircle"];
  2 [label="2", shape="circle"];
  3 [label="3", shape="circle"];
}
`))
}

func TestNFAToDFA2(t *testing.T) {
	n := nfa.NewStateFactory()

	s := make([]*nfa.State, 13)
	for i := range s {
		s[i] = n.NewState()
	}

	transitions := []struct {
		From  *nfa.State
		Input string
		To    *nfa.State
	}{
		{s[0], "ε", s[1]},
		{s[1], "ε", s[2]},
		{s[2], "s", s[3]},
		{s[3], "ε", s[10]},
		{s[1], "ε", s[4]},
		{s[4], "r", s[5]},
		{s[5], "ε", s[10]},
		{s[1], "ε", s[6]},
		{s[6], "n", s[7]},
		{s[7], "ε", s[10]},
		{s[1], "ε", s[8]},
		{s[8], "t", s[9]},
		{s[9], "ε", s[10]},
		{s[10], "ε", s[11]},
		{s[11], "ε", s[0]},
		{s[11], "ε", s[12]},
	}

	for _, tr := range transitions {
		var input any = tr.Input
		if tr.Input == "ε" {
			input = nfa.Epsilon
		}
		tr.From.AddTransition(tr.To, input)
	}

	s[12].Accept = true

	d := NFAToDFA(s[0])

	var res strings.Builder
	d.Print(&res)

	testutil.RequireEqualStr(t, res.String(), `
digraph G {
  rankdir="LR";
  0 -> 1 [label="s"];
  0 -> 1 [label="r"];
  0 -> 1 [label="n"];
  0 -> 1 [label="t"];
  1 -> 1 [label="s"];
  1 -> 1 [label="r"];
  1 -> 1 [label="n"];
  1 -> 1 [label="t"];
  0 [label="0", shape="circle"];
  1 [label="1", shape="doublecircle"];
}
`)

}

func TestNFAToDFA3(t *testing.T) {
	n := nfa.NewStateFactory()

	s := make([]*nfa.State, 17)
	for i := range s {
		s[i] = n.NewState()
	}

	transitions := []struct {
		From  int
		Input string
		To    int
	}{
		{0, "ε", 1},
		{1, "09", 2},
		{2, "09", 2},
		{2, "ε", 3},

		{0, "ε", 4},
		{4, "+", 5},
		{5, "ε", 6},

		{0, "ε", 7},
		{7, "-", 8},
		{8, "ε", 9},

		{0, "ε", 10},
		{10, "ε", 11},
		{11, "r", 12},
		{12, "ε", 15},
		{10, "ε", 13},
		{13, "n", 14},
		{14, "ε", 15},
		{15, "ε", 10},
		{15, "ε", 16},
	}

	s[3].Accept = true
	s[6].Accept = true
	s[9].Accept = true
	s[16].Accept = true

	for _, tr := range transitions {
		var input any
		switch {
		case tr.Input == "ε":
			input = nfa.Epsilon
		case len(tr.Input) == 2:
			input = rang3.Range{B: rune(tr.Input[0]), E: rune(tr.Input[1])}
		case len(tr.Input) == 1:
			input = rang3.Range{B: rune(tr.Input[0]), E: rune(tr.Input[0])}
		default:
			panic("not reached")
		}
		s[tr.From].AddTransition(s[tr.To], input)
	}

	d := NFAToDFA(s[0])

	var res strings.Builder
	d.Print(&res)

	testutil.RequireEqualStr(t, res.String(), `
digraph G {
  rankdir="LR";
  0 -> 1 [label="r"];
  0 -> 1 [label="n"];
  0 -> 2 [label="\\-"];
  0 -> 3 [label="+"];
  0 -> 4 [label="0-9"];
  1 -> 1 [label="r"];
  1 -> 1 [label="n"];
  4 -> 4 [label="0-9"];
  0 [label="0", shape="circle"];
  1 [label="1", shape="doublecircle"];
  2 [label="2", shape="doublecircle"];
  3 [label="3", shape="doublecircle"];
  4 [label="4", shape="doublecircle"];
}
`)
}
