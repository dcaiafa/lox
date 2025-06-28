package dfa

import (
	"cmp"
	"encoding/binary"
	"fmt"
	"io"
	"slices"
	"sort"

	"github.com/dcaiafa/lox/internal/base/set"
	"github.com/dcaiafa/lox/internal/base/stablemap"
	"github.com/dcaiafa/lox/internal/base/stack"
	"github.com/dcaiafa/lox/internal/lexergen/nfa"
)

type DFA struct {
	States []*State
}

func (d *DFA) Print(out io.Writer) {
	d.States[0].Print(out)
}

type State struct {
	ID          uint32
	Transitions stablemap.Map[any, *State]
	Accept      bool
	NonGreedy   bool
	NFAStates   []*nfa.State
	Data        any
}

func (s *State) AddTransition(toState *State, input any) {
	s.Transitions.Put(input, toState)
}

func (s *State) sig() string {
	sig := make([]byte, len(s.NFAStates)*4)
	for i, nfaState := range s.NFAStates {
		binary.BigEndian.PutUint32(sig[i*4:], nfaState.ID)
	}
	return string(sig)
}

func (n *State) Print(out io.Writer) {
	fmt.Fprintf(out, "digraph G {\n")
	fmt.Fprintf(out, "  rankdir=\"LR\";\n")

	type Edge struct {
		from  *State
		to    *State
		input any
	}

	var edges []Edge
	var visited set.Set[*State]
	var stack stack.Stack[*State]
	stack.Push(n)
	for !stack.Empty() {
		state := stack.Pop()
		if visited.Has(state) {
			continue
		}
		visited.Add(state)
		state.Transitions.ForEach(func(input any, destState *State) {
			edges = append(edges, Edge{
				from:  state,
				to:    destState,
				input: input,
			})
			stack.Push(destState)
		})
	}

	slices.SortStableFunc(edges, func(a, b Edge) int {
		c := cmp.Compare(a.from.ID, b.from.ID)
		if c == 0 {
			c = cmp.Compare(a.to.ID, b.to.ID)
		}
		return c
	})

	for _, e := range edges {
		inputStr := fmt.Sprintf("%v", e.input)
		fmt.Fprintf(out, "  %v -> %v [label=%q];\n", e.from.ID, e.to.ID, inputStr)
	}

	states := visited.Elements()
	sort.Slice(states, func(i, j int) bool {
		return states[i].ID < states[j].ID
	})

	for _, state := range states {
		var shape string
		switch {
		case state.Accept && state.NonGreedy:
			shape = "doubleoctagon"
		case state.NonGreedy:
			shape = "octagon"
		case state.Accept:
			shape = "doublecircle"
		default:
			shape = "circle"
		}
		fmt.Fprintf(out, "  %v [label=\"%v\", shape=%q];\n", state.ID, state.ID, shape)
	}
	fmt.Fprintf(out, "}\n")
}
