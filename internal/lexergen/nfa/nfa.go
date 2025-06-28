package nfa

import (
	"cmp"
	"fmt"
	"io"
	"slices"

	"github.com/dcaiafa/lox/internal/base/array"
	"github.com/dcaiafa/lox/internal/base/set"
	"github.com/dcaiafa/lox/internal/base/stablemap"
	"github.com/dcaiafa/lox/internal/base/stack"
)

type epsilon struct {
}

func (e epsilon) String() string {
	return "ε"
}

// Epsilon is the ε (empty string) input.
var Epsilon = epsilon{}

// State is a state in a NFA.
// NFAStates should only created using NFA.NewState().
type State struct {
	// ID of the state.
	// It is assigned by the NFA and should be read-only.
	ID uint32

	// Transitions of the state.
	// Transitions should only be added using NFA.AddTransitions.
	Transitions stablemap.Map[any, *array.Array[*State]]

	// Accept indicates that the state machine should accept/recognize the input.
	// You set this yourself.
	Accept bool

	// NonGreedy, in combination with Accept, indicate that the state machine
	// should accept the current string without consuming additional input.
	NonGreedy bool

	// Data is some user-data associated with this state.
	// Your set this yourself (or don't, I don't care).
	Data any
}

// AddTransition adds a transition between two states on a given input.
// The input can be any comparable data, including Epsilon.
func (s *State) AddTransition(to *State, input any) {
	states, ok := s.Transitions.Get(input)
	if !ok {
		states = new(array.Array[*State])
		s.Transitions.Put(input, states)
	}
	states.Add(to)
}

// NFA represents a Nondeterministic Finite Automaton.
// https://en.wikipedia.org/wiki/Nondeterministic_finite_automaton
// The NFA is state machine where:
// - A state is allowed to have multiple transitions for the same input.
// - The input ε (empty string) is allowed.
type StateFactory struct {
	nextID uint32
}

func NewStateFactory() *StateFactory {
	return &StateFactory{}
}

// NewState creates a new state in the NFA.
func (n *StateFactory) NewState() *State {
	s := &State{
		ID: n.nextID,
	}
	n.nextID++
	return s
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
		state.Transitions.ForEach(func(input any, destStates *array.Array[*State]) {
			for _, destState := range destStates.Elements() {
				edges = append(edges, Edge{
					from:  state,
					to:    destState,
					input: input,
				})
				stack.Push(destState)
			}
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
	slices.SortFunc(states, func(a, b *State) int {
		return cmp.Compare(a.ID, b.ID)

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
