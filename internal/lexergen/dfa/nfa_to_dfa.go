package dfa

import (
	"sort"

	"github.com/dcaiafa/lox/internal/base/array"
	"github.com/dcaiafa/lox/internal/base/set"
	"github.com/dcaiafa/lox/internal/base/stack"
	"github.com/dcaiafa/lox/internal/lexergen/nfa"
)

func NFAToDFA(n *nfa.State) *DFA {
	// DFA states already created, indexed by their signature:
	// the unique combination of NFA states that define a DFA state.
	states := make(map[string]*State)

	// Create the start state.
	start := eClosure(set.New[*nfa.State](n))
	states[start.sig()] = start

	// Starting from the newly created 'start' state, build all DFA states, one by
	// one.
	var stack stack.Stack[*State]
	stack.Push(start)
	for !stack.Empty() {
		from := stack.Pop()

		inputs := getInputs(from.NFAStates)
		inputs.ForEach(func(input any) {
			// Create a subset of all NFA states reachable from the the NFA states
			// composing 'from', using 'input'.
			var subset set.Set[*nfa.State]
			for _, fromNFA := range from.NFAStates {
				for _, toNFA := range fromNFA.Transitions.GetOrZero(input).Elements() {
					subset.Add(toNFA)
				}
			}

			// Expand the subset to include states reachable via ε.
			to := eClosure(subset)
			toSig := to.sig()

			// Look for an existing DFA state with the same set of NFA states
			// (uniquely represented by its signature).
			if existing := states[toSig]; existing != nil {
				// Reuse state.
				from.AddTransition(existing, input)
			} else {
				// Create a new state.
				states[toSig] = to
				from.AddTransition(to, input)

				// Add it to the stack for processing.
				stack.Push(to)
			}
		})
	}

	dfa := &DFA{
		States: transitiveClosure(start),
	}

	optimize(dfa)

	return dfa
}

func eClosure(nfaStates set.Set[*nfa.State]) *State {
	dfaState := new(State)

	closure := make(map[uint32]*nfa.State)
	var stack stack.Stack[*nfa.State]

	nfaStates.ForEach(func(s *nfa.State) {
		closure[s.ID] = s
		stack.Push(s)
		dfaState.Accept = dfaState.Accept || s.Accept
		dfaState.NonGreedy = dfaState.NonGreedy || s.NonGreedy
	})

	for !stack.Empty() {
		state := stack.Pop()
		eTransitions := state.Transitions.GetOrZero(nfa.Epsilon)
		for _, to := range eTransitions.Elements() {
			if _, ok := closure[to.ID]; !ok {
				closure[to.ID] = to
				stack.Push(to)
				dfaState.Accept = dfaState.Accept || to.Accept
				dfaState.NonGreedy = dfaState.NonGreedy || to.NonGreedy
			}
		}
	}

	dfaState.NFAStates = make([]*nfa.State, 0, len(closure))
	for _, nfaState := range closure {
		dfaState.NFAStates = append(dfaState.NFAStates, nfaState)
	}
	sort.Slice(dfaState.NFAStates, func(i, j int) bool {
		return dfaState.NFAStates[i].ID < dfaState.NFAStates[j].ID
	})

	return dfaState
}

func getInputs(nfaStates []*nfa.State) set.Set[any] {
	var inputs set.Set[any]
	for _, nfaState := range nfaStates {
		nfaState.Transitions.ForEach(func(input any, _ *array.Array[*nfa.State]) {
			if input != nfa.Epsilon {
				inputs.Add(input)
			}
		})
	}
	return inputs
}

func transitiveClosure(s *State) []*State {
	var visited set.Set[*State]
	var pending stack.Stack[*State]
	var states []*State

	visited.Add(s)
	pending.Push(s)
	for !pending.Empty() {
		s = pending.Pop()
		s.ID = uint32(len(states))
		states = append(states, s)
		dests := s.Transitions.Values()
		sort.Slice(dests, func(i, j int) bool {
			return dests[i].ID < dests[j].ID
		})
		for _, dest := range dests {
			if !visited.Has(dest) {
				visited.Add(dest)
				pending.Push(dest)
			}
		}
	}
	return states
}
