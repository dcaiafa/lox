package codegen

import (
	"cmp"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/CloudyKit/jet/v6"
	"github.com/dcaiafa/lox/internal/base/assert"
	"github.com/dcaiafa/lox/internal/lexergen/dfa"
	"github.com/dcaiafa/lox/internal/lexergen/mode"
	"github.com/dcaiafa/lox/internal/lexergen/rang3"
)

const lexerTemplate = `
{{ range _, mode := modes() }}

var _lexerMode{{mode.Index}} = []uint32 {
	{{ mode_table(mode) }}
}

{{ end }}


var _lexerModes = [][]uint32 {
{{ range _, mode := modes() }}
	_lexerMode{{mode.Index}},
{{ end }}
}

// Flag for the mode table that indicates that the state is non-greedy
// accepting. At this state, the state machine is expected to accept the current
// string without attempting to consume additional input.
const  _stateNonGreedyAccepting = 1

const (
	_lexerConsume  = 0
	_lexerAccept   = 1
	_lexerDiscard  = 2
	_lexerTryAgain = 3
	_lexerEOF      = 4
	_lexerError    = -1
)

type _LexerStateMachine struct {
	token int
	state int
	mode  []uint32
	modeStack _Stack[[]uint32]
}

func (l *_LexerStateMachine) PushRune(r rune) int {
	if l.mode == nil {
		l.mode = _lexerMode0
	}

	mode := l.mode

	// Find the table row corresponding to state.
	i := int(mode[int(l.state)])
	count := int(mode[i])
	i++
	end := i + count

	// The format of each row is as follows:
	//
	//   stateFlags uint32
	//   gotoCount uint32
	//   [gotoCount]struct{
	//     rangeBegin uint32
	//     rangeEnd   uint32
	//     gotoState  uint32
	//   }
	//   [actionCount]struct {
	//     actionType  uint32
	//     actionParam uint32
	//   }
	//
	// Where 'actionCount' is determined by the amount of uint32 left in the row.

	flags := mode[i]
	gotoN := int(mode[i+1])
	i += 2

	if flags & _stateNonGreedyAccepting == 0 {
		// Use binary-search to find the next state.
		b := 0
		e := gotoN
		for b < e {
			j := b + (e-b)/2
			k := i + j*3
			switch {
			case r >= rune(mode[k]) && r <= rune(mode[k+1]):
				l.state = int(mode[k+2])
				return _lexerConsume
			case r < rune(mode[k]):
				e = j
			case r > rune(mode[k+1]):
				b = j + 1
			default:
				panic("not reached")
			}
		}
	}

	// Move 'i' to the beginning of the actions section.
	i += gotoN * 3

	for ; i < end; i += 2 {
		switch mode[i] {
		case 1: // PushMode
			modeIndex := int(mode[i+1])
			l.modeStack.Push(mode)
			l.mode = _lexerModes[modeIndex]
		case 2: // PopMode
		  if len(l.modeStack) == 0 {
				return _lexerError
			}
			l.mode = l.modeStack.Peek(0)
			l.modeStack.Pop(1)
		case 3: // Accept
			l.token = int(mode[i+1])
			l.state = 0
			return _lexerAccept
		case 4: // Discard
			l.state = 0
			return _lexerDiscard
		case 5: // Accum
			l.state = 0
			return _lexerTryAgain
		}
	}

	if l.state == 0 && r == -1 {
		return _lexerEOF
	}

	return _lexerError}

func (l *_LexerStateMachine) Reset() {
	l.mode = nil
	l.state = 0
}

func (l *_LexerStateMachine) Token() int {
	return l.token
}
`

// Flag for the mode table that indicates that the state is non-greedy
// accepting. At this state, the state machine is expected to accept the current
// string without attempting to consume additional input.
const stateNonGreedyAcceptingFlag uint32 = 1

func (c *context) EmitLexer() bool {
	vars := make(jet.VarMap)

	vars.Set("array", func(arr []uint32) string {
		var str strings.Builder
		WriteArray(&str, arr)
		return str.String()
	})

	vars.Set("modes", func() []*mode.Mode {
		modes := make([]*mode.Mode, 0, len(c.LexerModes))
		for _, mode := range c.LexerModes {
			modes = append(modes, mode)
		}
		slices.SortFunc(modes, func(a, b *mode.Mode) int {
			return cmp.Compare(a.Index, b.Index)
		})
		return modes
	})

	vars.Set("mode_table", func(m *mode.Mode) string {
		table := newTable[uint32]()
		for _, state := range m.DFA.States {
			var row []uint32
			actions := state.Data.(*mode.Actions)

			inputs := make([]rang3.Range, 0, state.Transitions.Len())
			state.Transitions.ForEach(func(eventRaw any, toState *dfa.State) {
				inputs = append(inputs, eventRaw.(rang3.Range))
			})
			slices.SortFunc(inputs, rang3.Compare)

			var stateFlags uint32
			if state.Accept && state.NonGreedy {
				stateFlags = stateNonGreedyAcceptingFlag
			}

			row = append(row, stateFlags)
			row = append(row, uint32(len(inputs)))
			for _, input := range inputs {
				toState, ok := state.Transitions.Get(input)
				assert.True(ok)
				row = append(row, uint32(input.B), uint32(input.E), toState.ID)
			}

			if actions != nil {
				for _, action := range actions.Actions {
					switch action.Type {
					case mode.ActionPushMode:
						mode := c.LexerModes[action.Mode]
						assert.True(mode != nil)
						row = append(row, 1, uint32(mode.Index))
					case mode.ActionPopMode:
						row = append(row, 2, 0)
					case mode.ActionAccept:
						row = append(row, 3, uint32(action.Terminal))
					case mode.ActionDiscard:
						row = append(row, 4, 0)
					case mode.ActionAccum:
						row = append(row, 5, 0)
					default:
						panic("unreachable")
					}
				}
			}

			assert.True(len(row) > 0)
			table.AddRow(int(state.ID), row)
		}
		return table.String()
	})

	lexerGen := renderTemplate(
		c.GoPackageName, c.GoPackagePath, lexerTemplate, vars)

	err := os.WriteFile(
		filepath.Join(c.Dir, lexerGenGo), []byte(lexerGen), 0666)
	if err != nil {
		c.Errs.GeneralError(err)
		return false
	}

	return true
}
