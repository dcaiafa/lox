package dfa

import (
	"github.com/dcaiafa/lox/internal/base/assert"
	"github.com/dcaiafa/lox/internal/base/set"
	"github.com/dcaiafa/lox/internal/base/stablemap"
	"github.com/dcaiafa/lox/internal/lexergen/nfa"
)

type partitions struct {
	stateToGroup stablemap.Map[*State, int]
	groupToState stablemap.Map[int, *set.Set[*State]]
}

func (p *partitions) Add(s *State, group int) {
	p.stateToGroup.Put(s, group)
	states, ok := p.groupToState.Get(group)
	if !ok {
		states = new(set.Set[*State])
		p.groupToState.Put(group, states)
	}
	states.Add(s)
}

func (p *partitions) Remove(s *State) {
	group, ok := p.stateToGroup.Get(s)
	assert.True(ok)
	states, ok := p.groupToState.Get(group)
	assert.True(ok)
	states.Remove(s)
}

func (p *partitions) Move(s *State, group int) {
	p.Remove(s)
	p.Add(s, group)
}

func (p *partitions) Count() int {
	return p.groupToState.Len()
}

func (p *partitions) GetGroup(group int) *set.Set[*State] {
	states, ok := p.groupToState.Get(group)
	assert.True(ok)
	return states
}

func (p *partitions) GetStateGroup(s *State) int {
	group, ok := p.stateToGroup.Get(s)
	assert.True(ok)
	return group
}

func optimize(d *DFA) {
	p := &partitions{}

	// Divide states into two initial partitions:
	// #0: non-accepting
	// #1: accepting
	for _, s := range d.States {
		if s.Accept {
			p.Add(s, 1)
		} else {
			p.Add(s, 0)
		}
	}

	if p.Count() < 2 {
		// All states are accepting.
		// DFA can't be optimized any futher.
		return
	}

	// Sub-partition
	pcount := 0
	for pcount != p.Count() {
		pcount = p.Count()
		for i := 0; i < pcount; i++ {
			subPartition(p, i)
		}
	}

	newStates := make([]*State, p.Count())
	for i := range newStates {
		newStates[i] = new(State)
	}

	var startGroup int
	for group := 0; group < p.Count(); group++ {
		groupStates := p.GetGroup(group)
		groupStates.ForEach(func(s *State) {
			if s.ID == 0 {
				startGroup = group
			}
			newStates[group].Accept = newStates[group].Accept || s.Accept
			newStates[group].NonGreedy = newStates[group].NonGreedy || s.NonGreedy
			newStates[group].NFAStates = append(
				newStates[group].NFAStates,
				s.NFAStates...)
		})
	}

	for _, s := range d.States {
		fromGroup := p.GetStateGroup(s)
		fromState := newStates[fromGroup]

		s.Transitions.ForEach(func(input any, ts *State) {
			toGroup := p.GetStateGroup(ts)
			toState := newStates[toGroup]
			fromState.AddTransition(toState, input)
		})
	}

	// Place group with the start event at index 0 so that it becomes the starting
	// event of the new DFA.
	newStates[0], newStates[startGroup] = newStates[startGroup], newStates[0]

	for i := range newStates {
		newStates[i].ID = uint32(i)
	}

	d.States = newStates
}

func subPartition(p *partitions, group int) {
	transitionGroup := func(s *State, input any) int {
		toState, ok := s.Transitions.Get(input)
		if !ok {
			return -1
		}
		return p.GetStateGroup(toState)
	}

	newGroup := p.Count()
	states := p.GetGroup(group)

	var inputs set.Set[any]
	states.ForEach(func(s *State) {
		s.Transitions.ForEach(func(input any, toState *State) {
			inputs.Add(input)
		})
	})

	var first *State
	var move set.Set[*State]
	states.ForEach(func(s *State) {
		if first == nil {
			first = s
			return
		}

		// Distinguish states that have transitions to a different group/partition.
		inputs.ForEach(func(input any) {
			toGroupFirst := transitionGroup(first, input)
			toGroupS := transitionGroup(s, input)
			if toGroupFirst != toGroupS {
				move.Add(s)
			}
		})

		// Distinguish states that have different accepting NFA states.
		// This is not covered in theory, as far as I could determine.
		// Given the example below:
		//
		//           +
		//       +------> ((1))
		//      /
		//    (0)
		//      \    -
		//       +------> ((2))
		//
		// We don't want to combine 1 & 2. Even if the resulting DFA would recognize
		// the same inputs, we need to differentiate 1 from 2 at code generation.
		if first.Accept {
			assert.True(s.Accept)
			if !acceptingNFAStates(first).Equal(acceptingNFAStates(s)) {
				move.Add(s)
				return
			}
		}
	})
	move.ForEach(func(s *State) {
		p.Move(s, newGroup)
	})
}

func acceptingNFAStates(s *State) set.Set[*nfa.State] {
	var accepting set.Set[*nfa.State]
	for _, ns := range s.NFAStates {
		if ns.Accept {
			accepting.Add(ns)
		}
	}
	return accepting
}
