package mode

import (
	"math"
	"strings"
	"testing"

	"github.com/dcaiafa/lox/internal/lexergen/dfa"
	"github.com/dcaiafa/lox/internal/lexergen/nfa"
	"github.com/dcaiafa/lox/internal/lexergen/rang3"
	"github.com/dcaiafa/lox/internal/testutil"
)

func TestNormalizeInputs(t *testing.T) {
	n := nfa.NewStateFactory()
	s := make([]*nfa.State, 12)
	for i := range s {
		s[i] = n.NewState()
	}

	s[0].AddTransition(s[1], nfa.Epsilon)
	s[1].AddTransition(s[2], rang3.Range{B: '5', E: '5'})
	s[2].AddTransition(s[7], nfa.Epsilon)
	s[0].AddTransition(s[3], nfa.Epsilon)
	s[3].AddTransition(s[4], rang3.Range{B: '0', E: '9'})
	s[4].AddTransition(s[7], nfa.Epsilon)
	s[0].AddTransition(s[5], nfa.Epsilon)
	s[5].AddTransition(s[6], rang3.Range{B: 0, E: rune(math.MaxInt32)})
	s[6].AddTransition(s[7], nfa.Epsilon)

	s[2].Accept = true
	s[4].Accept = true
	s[6].Accept = true

	normalizeInputs(s[0])

	var str strings.Builder
	s[0].Print(&str)

	testutil.RequireEqualStr(t, str.String(), `
digraph G {
  rankdir="LR";
  0 -> 1 [label="ε"];
  0 -> 3 [label="ε"];
  0 -> 5 [label="ε"];
  1 -> 2 [label="5"];
  2 -> 7 [label="ε"];
  3 -> 4 [label="0-4"];
  3 -> 4 [label="5"];
  3 -> 4 [label="6-9"];
  4 -> 7 [label="ε"];
  5 -> 6 [label="\\u0000-/"];
  5 -> 6 [label=":-\\u7fffffff"];
  5 -> 6 [label="0-4"];
  5 -> 6 [label="5"];
  5 -> 6 [label="6-9"];
  6 -> 7 [label="ε"];
  0 [label="0", shape="circle"];
  1 [label="1", shape="circle"];
  2 [label="2", shape="doublecircle"];
  3 [label="3", shape="circle"];
  4 [label="4", shape="doublecircle"];
  5 [label="5", shape="circle"];
  6 [label="6", shape="doublecircle"];
  7 [label="7", shape="circle"];
}
`)
}

func TestNonGreedy(t *testing.T) {
	//	                      ε
	//	                   +-----+
	//	                   |     |
	//	   /     *  ?  ε   v .   |    ?
	//	s0 -> s1 -> s2 -> s3 -> s4 -> s5 -> s6 -> s7
	//	            |                 ^
	//	            |        ε        |
	//	            +-----------------+
	//
	// Recognizes: '/*' .*? '*/'

	n := nfa.NewStateFactory()

	s0 := n.NewState()
	s1 := n.NewState()
	s2 := n.NewState()
	s2.NonGreedy = true
	s3 := n.NewState()
	s4 := n.NewState()
	s5 := n.NewState()
	s5.NonGreedy = true
	s6 := n.NewState()
	s7 := n.NewState()

	s0.AddTransition(s1, rang3.Range{B: '/', E: '/'})
	s1.AddTransition(s2, rang3.Range{B: '*', E: '*'})
	s2.AddTransition(s3, nfa.Epsilon)
	s2.AddTransition(s5, nfa.Epsilon)
	s3.AddTransition(s4, rang3.Range{B: 0, E: 0x0010FFFF})
	s4.AddTransition(s3, nfa.Epsilon)
	s4.AddTransition(s5, nfa.Epsilon)
	s5.AddTransition(s6, rang3.Range{B: '*', E: '*'})
	s6.AddTransition(s7, rang3.Range{B: '/', E: '/'})
	s7.Accept = true

	normalizeInputs(s0)

	d := dfa.NFAToDFA(s0)
	var res strings.Builder
	d.Print(&res)

	testutil.RequireEqualStr(t, res.String(), `
digraph G {
  rankdir="LR";
  0 -> 2 [label="/"];
  1 -> 3 [label="\\u0000-)"];
  1 -> 3 [label="+-."];
  1 -> 3 [label="/"];
  1 -> 3 [label="0-\\u10ffff"];
  1 -> 4 [label="*"];
  2 -> 3 [label="*"];
  3 -> 3 [label="\\u0000-)"];
  3 -> 3 [label="+-."];
  3 -> 3 [label="/"];
  3 -> 3 [label="0-\\u10ffff"];
  3 -> 4 [label="*"];
  4 -> 1 [label="/"];
  4 -> 3 [label="\\u0000-)"];
  4 -> 3 [label="+-."];
  4 -> 3 [label="0-\\u10ffff"];
  4 -> 4 [label="*"];
  0 [label="0", shape="circle"];
  1 [label="1", shape="doubleoctagon"];
  2 [label="2", shape="circle"];
  3 [label="3", shape="octagon"];
  4 [label="4", shape="octagon"];
}
`)
}
