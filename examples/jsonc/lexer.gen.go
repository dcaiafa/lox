package main

var _lexerMode0 = []uint32{
	36, 90, 104, 113, 127, 139, 153, 165, 182, 200, 208, 217, 231, 243,
	257, 263, 270, 276, 281, 287, 292, 298, 303, 309, 314, 320, 326, 331,
	337, 343, 348, 354, 359, 365, 370, 376, 53, 0, 17, 9, 10, 1,
	13, 13, 1, 32, 32, 1, 34, 34, 15, 44, 44, 26, 45, 45,
	10, 47, 47, 2, 48, 48, 13, 49, 57, 7, 58, 58, 23, 91,
	91, 31, 93, 93, 29, 102, 102, 25, 110, 110, 18, 116, 116, 32,
	123, 123, 35, 125, 125, 33, 13, 0, 3, 9, 10, 1, 13, 13,
	1, 32, 32, 1, 4, 0, 8, 0, 2, 42, 42, 6, 47, 47,
	4, 13, 1, 3, 0, 9, 4, 10, 10, 3, 11, 1114111, 4, 4,
	0, 11, 0, 3, 0, 9, 4, 10, 10, 3, 11, 1114111, 4, 13,
	1, 3, 0, 41, 6, 42, 42, 8, 43, 1114111, 6, 4, 0, 11,
	0, 3, 0, 41, 6, 42, 42, 8, 43, 1114111, 6, 16, 0, 4,
	46, 46, 16, 48, 57, 7, 69, 69, 12, 101, 101, 12, 3, 12,
	17, 0, 5, 0, 41, 6, 42, 42, 8, 43, 46, 6, 47, 47,
	5, 48, 1114111, 6, 7, 0, 1, 48, 57, 9, 3, 12, 8, 0,
	2, 48, 48, 13, 49, 57, 7, 13, 0, 3, 48, 57, 11, 69,
	69, 12, 101, 101, 12, 3, 12, 11, 0, 3, 43, 43, 14, 45,
	45, 14, 49, 57, 9, 13, 0, 3, 46, 46, 16, 69, 69, 12,
	101, 101, 12, 3, 12, 5, 0, 1, 49, 57, 9, 6, 0, 0,
	1, 1, 5, 0, 5, 0, 1, 48, 57, 11, 4, 0, 0, 3,
	10, 5, 0, 1, 117, 117, 20, 4, 0, 0, 3, 9, 5, 0,
	1, 108, 108, 22, 4, 0, 0, 3, 8, 5, 0, 1, 108, 108,
	17, 4, 0, 0, 3, 7, 5, 0, 1, 117, 117, 34, 5, 0,
	1, 97, 97, 27, 4, 0, 0, 3, 6, 5, 0, 1, 108, 108,
	28, 5, 0, 1, 115, 115, 30, 4, 0, 0, 3, 5, 5, 0,
	1, 101, 101, 19, 4, 0, 0, 3, 4, 5, 0, 1, 114, 114,
	24, 4, 0, 0, 3, 3, 5, 0, 1, 101, 101, 21, 4, 0,
	0, 3, 2,
}

var _lexerMode1 = []uint32{
	10, 28, 33, 28, 63, 28, 75, 87, 94, 106, 17, 0, 5, 32,
	33, 5, 34, 34, 7, 35, 91, 5, 92, 92, 2, 93, 1114111, 5,
	4, 0, 0, 5, 0, 29, 0, 9, 34, 34, 3, 47, 47, 3,
	92, 92, 3, 98, 98, 3, 102, 102, 3, 110, 110, 3, 114, 114,
	3, 116, 116, 3, 117, 117, 4, 11, 0, 3, 48, 57, 9, 65,
	70, 9, 97, 102, 9, 11, 0, 3, 48, 57, 1, 65, 70, 1,
	97, 102, 1, 6, 0, 0, 2, 0, 3, 11, 11, 0, 3, 48,
	57, 6, 65, 70, 6, 97, 102, 6, 11, 0, 3, 48, 57, 8,
	65, 70, 8, 97, 102, 8,
}

var _lexerModes = [][]uint32{

	_lexerMode0,

	_lexerMode1,
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
