package main



var _lexerMode0 = []uint32 {
	11, 43, 50, 54, 58, 62, 66, 70, 74, 78, 82, 31, 10, 32, 
32, 1, 37, 37, 5, 40, 40, 3, 41, 41, 2, 42, 42, 7, 
43, 43, 9, 45, 45, 8, 47, 47, 6, 48, 57, 10, 94, 94, 
4, 6, 1, 32, 32, 1, 4, 0, 3, 0, 3, 10, 3, 0, 
3, 9, 3, 0, 3, 8, 3, 0, 3, 7, 3, 0, 3, 6, 
3, 0, 3, 5, 3, 0, 3, 4, 3, 0, 3, 3, 6, 1, 
48, 57, 10, 3, 2, 
}




var _lexerModes = [][]uint32 {

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

	return _lexerError}

func (l *_LexerStateMachine) Reset() {
	l.mode = nil
	l.state = 0
}

func (l *_LexerStateMachine) Token() int {
	return l.token
}
