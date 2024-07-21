package main

var _lexerMode0 = []uint32{
	33, 116, 129, 143, 147, 163, 179, 201, 217, 239, 261, 277, 299, 321,
	343, 347, 369, 376, 380, 387, 391, 398, 402, 406, 410, 414, 418, 422,
	426, 433, 437, 441, 445, 82, 27, 9, 9, 1, 10, 10, 3, 13,
	13, 1, 32, 32, 1, 34, 34, 2, 40, 40, 32, 41, 41, 31,
	42, 42, 23, 43, 43, 25, 44, 44, 30, 45, 45, 24, 47, 47,
	22, 48, 48, 17, 49, 57, 16, 60, 60, 20, 61, 61, 28, 62,
	62, 18, 65, 90, 4, 95, 95, 4, 97, 97, 8, 98, 110, 4,
	111, 111, 9, 112, 118, 4, 119, 119, 15, 120, 122, 4, 123, 123,
	27, 125, 125, 26, 12, 3, 9, 9, 1, 13, 13, 1, 32, 32,
	1, 4, 0, 13, 4, 0, 9, 2, 11, 33, 2, 34, 34, 14,
	35, 1114111, 2, 3, 0, 3, 23, 15, 4, 48, 57, 4, 65, 90,
	4, 95, 95, 4, 97, 122, 4, 3, 22, 15, 4, 48, 57, 4,
	65, 90, 4, 95, 95, 4, 97, 122, 4, 3, 21, 21, 6, 48,
	57, 4, 65, 90, 4, 95, 95, 4, 97, 99, 4, 100, 100, 5,
	101, 122, 4, 3, 22, 15, 4, 48, 57, 4, 65, 90, 4, 95,
	95, 4, 97, 122, 4, 3, 20, 21, 6, 48, 57, 4, 65, 90,
	4, 95, 95, 4, 97, 109, 4, 110, 110, 6, 111, 122, 4, 3,
	22, 21, 6, 48, 57, 4, 65, 90, 4, 95, 95, 4, 97, 113,
	4, 114, 114, 7, 115, 122, 4, 3, 22, 15, 4, 48, 57, 4,
	65, 90, 4, 95, 95, 4, 97, 122, 4, 3, 19, 21, 6, 48,
	57, 4, 65, 90, 4, 95, 95, 4, 97, 104, 4, 105, 105, 12,
	106, 122, 4, 3, 22, 21, 6, 48, 57, 4, 65, 90, 4, 95,
	95, 4, 97, 107, 4, 108, 108, 13, 109, 122, 4, 3, 22, 21,
	6, 48, 57, 4, 65, 90, 4, 95, 95, 4, 97, 100, 4, 101,
	101, 10, 102, 122, 4, 3, 22, 3, 0, 3, 18, 21, 6, 48,
	57, 4, 65, 90, 4, 95, 95, 4, 97, 103, 4, 104, 104, 11,
	105, 122, 4, 3, 22, 6, 1, 48, 57, 16, 3, 17, 3, 0,
	3, 17, 6, 1, 61, 61, 19, 3, 14, 3, 0, 3, 15, 6,
	1, 61, 61, 21, 3, 12, 3, 0, 3, 13, 3, 0, 3, 11,
	3, 0, 3, 10, 3, 0, 3, 9, 3, 0, 3, 8, 3, 0,
	3, 7, 3, 0, 3, 6, 6, 1, 61, 61, 29, 3, 5, 3,
	0, 3, 16, 3, 0, 3, 4, 3, 0, 3, 3, 3, 0, 3,
	2,
}

var _lexerModes = [][]uint32{

	_lexerMode0,
}

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

	// The format of the row is as follows:
	//
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

	gotoN := int(mode[i])
	i++

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

	// Move 'i' to the beginning of the actions section.
	i += gotoN * 3

	for ; i < end; i += 2 {
		switch mode[i] {
		case 1: // PushMode
			modeIndex := int(mode[i+1])
			l.modeStack.Push(mode)
			l.mode = _lexerModes[modeIndex]
		case 2: // PopMode
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

	if l.state == 0 && r == 0 {
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
