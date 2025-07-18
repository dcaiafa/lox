package main

var _lexerMode0 = []uint32{
	14, 50, 58, 64, 75, 83, 91, 96, 101, 106, 111, 116, 121, 126,
	35, 0, 11, 32, 32, 1, 37, 37, 9, 40, 40, 7, 41, 41,
	6, 42, 42, 11, 43, 43, 13, 45, 45, 12, 47, 47, 10, 48,
	48, 5, 49, 57, 3, 94, 94, 8, 7, 0, 1, 32, 32, 1,
	4, 0, 5, 0, 1, 48, 57, 4, 10, 0, 2, 46, 46, 2,
	48, 57, 3, 3, 10, 7, 0, 1, 48, 57, 4, 3, 10, 7,
	0, 1, 46, 46, 2, 3, 10, 4, 0, 0, 3, 9, 4, 0,
	0, 3, 8, 4, 0, 0, 3, 7, 4, 0, 0, 3, 6, 4,
	0, 0, 3, 5, 4, 0, 0, 3, 4, 4, 0, 0, 3, 3,
	4, 0, 0, 3, 2,
}

var _lexerModes = [][]uint32{

	_lexerMode0,
}

// Flag for the mode table that indicates that the state is non-greedy
// accepting. At this state, the state machine is expected to accept the current
// string without attempting to consume additional input.
const _stateNonGreedyAccepting = 1

const (
	_lexerConsume  = 0
	_lexerAccept   = 1
	_lexerDiscard  = 2
	_lexerTryAgain = 3
	_lexerEOF      = 4
	_lexerError    = -1
)

type _LexerStateMachine struct {
	token     int
	state     int
	mode      []uint32
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

	if flags&_stateNonGreedyAccepting == 0 {
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

	return _lexerError
}

func (l *_LexerStateMachine) Reset() {
	l.mode = nil
	l.state = 0
}

func (l *_LexerStateMachine) Token() int {
	return l.token
}
