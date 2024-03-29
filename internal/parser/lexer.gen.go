package parser

var _lexerMode0 = []uint32{
	93, 164, 177, 182, 192, 236, 242, 248, 255, 271, 281, 291, 307, 317,
	333, 349, 359, 375, 391, 407, 417, 433, 455, 471, 487, 497, 513, 529,
	545, 558, 574, 584, 600, 625, 641, 657, 673, 689, 699, 715, 731, 741,
	757, 776, 792, 808, 818, 834, 850, 866, 882, 898, 908, 921, 937, 953,
	963, 976, 992, 1008, 1018, 1040, 1062, 1078, 1088, 1104, 1120, 1136, 1146, 1162,
	1178, 1194, 1204, 1220, 1236, 1243, 1259, 1275, 1279, 1295, 1302, 1312, 1316, 1320,
	1324, 1328, 1332, 1336, 1340, 1344, 1348, 1352, 1356, 70, 23, 9, 10, 1,
	13, 13, 1, 32, 32, 1, 39, 39, 6, 40, 40, 85, 41, 41,
	84, 42, 42, 79, 43, 43, 74, 44, 44, 91, 45, 45, 83, 47,
	47, 2, 48, 57, 7, 59, 59, 92, 61, 61, 90, 63, 63, 82,
	64, 64, 4, 65, 90, 8, 91, 91, 5, 97, 122, 8, 123, 123,
	88, 124, 124, 89, 125, 125, 87, 126, 126, 86, 12, 3, 9, 10,
	1, 13, 13, 1, 32, 32, 1, 4, 0, 4, 1, 47, 47, 3,
	9, 2, 0, 9, 3, 11, 1114111, 3, 4, 0, 43, 14, 95, 95,
	9, 97, 99, 9, 100, 100, 31, 101, 101, 21, 102, 102, 34, 103,
	107, 9, 108, 108, 60, 109, 109, 42, 110, 111, 9, 112, 112, 32,
	113, 113, 9, 114, 114, 20, 115, 115, 33, 116, 122, 9, 5, 0,
	1, 1, 3, 35, 5, 0, 1, 2, 5, 0, 6, 1, 48, 57,
	7, 3, 33, 15, 4, 48, 57, 8, 65, 90, 8, 95, 95, 8,
	97, 122, 8, 3, 32, 9, 2, 95, 95, 9, 97, 122, 9, 3,
	31, 9, 2, 95, 95, 9, 97, 122, 9, 3, 29, 15, 4, 95,
	95, 9, 97, 115, 9, 116, 116, 10, 117, 122, 9, 3, 31, 9,
	2, 95, 95, 9, 97, 122, 9, 3, 30, 15, 4, 95, 95, 9,
	97, 103, 9, 104, 104, 11, 105, 122, 9, 3, 31, 15, 4, 95,
	95, 9, 97, 115, 9, 116, 116, 12, 117, 122, 9, 3, 31, 9,
	2, 95, 95, 9, 97, 122, 9, 3, 26, 15, 4, 95, 95, 9,
	97, 102, 9, 103, 103, 13, 104, 122, 9, 3, 31, 15, 4, 95,
	95, 9, 97, 104, 9, 105, 105, 14, 106, 122, 9, 3, 31, 15,
	4, 95, 95, 9, 97, 113, 9, 114, 114, 15, 115, 122, 9, 3,
	31, 9, 2, 95, 95, 9, 97, 122, 9, 3, 22, 15, 4, 95,
	95, 9, 97, 104, 9, 105, 105, 16, 106, 122, 9, 3, 31, 21,
	6, 95, 95, 9, 97, 108, 9, 109, 109, 17, 110, 113, 9, 114,
	114, 27, 115, 122, 9, 3, 31, 15, 4, 95, 95, 9, 97, 110,
	9, 111, 111, 18, 112, 122, 9, 3, 31, 15, 4, 95, 95, 9,
	97, 102, 9, 103, 103, 19, 104, 122, 9, 3, 31, 9, 2, 95,
	95, 9, 97, 122, 9, 3, 23, 15, 4, 95, 95, 9, 97, 114,
	9, 115, 115, 26, 116, 122, 9, 3, 31, 15, 4, 95, 95, 9,
	97, 98, 9, 99, 99, 52, 100, 122, 9, 3, 31, 15, 4, 95,
	95, 9, 97, 113, 9, 114, 114, 22, 115, 122, 9, 3, 31, 12,
	3, 95, 95, 9, 97, 97, 23, 98, 122, 9, 3, 31, 15, 4,
	95, 95, 9, 97, 100, 9, 101, 101, 24, 102, 122, 9, 3, 31,
	9, 2, 95, 95, 9, 97, 122, 9, 3, 21, 15, 4, 95, 95,
	9, 97, 104, 9, 105, 105, 25, 106, 122, 9, 3, 31, 24, 7,
	95, 95, 9, 97, 97, 47, 98, 110, 9, 111, 111, 39, 112, 116,
	9, 117, 117, 38, 118, 122, 9, 3, 31, 15, 4, 95, 95, 9,
	97, 115, 9, 116, 116, 56, 117, 122, 9, 3, 31, 15, 4, 95,
	95, 9, 97, 113, 9, 114, 114, 28, 115, 122, 9, 3, 31, 15,
	4, 95, 95, 9, 97, 99, 9, 100, 100, 29, 101, 122, 9, 3,
	31, 15, 4, 95, 95, 9, 97, 110, 9, 111, 111, 30, 112, 122,
	9, 3, 31, 9, 2, 95, 95, 9, 97, 122, 9, 3, 20, 15,
	4, 95, 95, 9, 97, 114, 9, 115, 115, 46, 116, 122, 9, 3,
	31, 15, 4, 95, 95, 9, 97, 111, 9, 112, 112, 40, 113, 122,
	9, 3, 31, 9, 2, 95, 95, 41, 97, 122, 9, 3, 31, 15,
	4, 95, 95, 9, 97, 108, 9, 109, 109, 64, 110, 122, 9, 3,
	31, 18, 5, 95, 95, 9, 97, 97, 48, 98, 110, 9, 111, 111,
	35, 112, 122, 9, 3, 31, 15, 4, 95, 95, 9, 97, 113, 9,
	114, 114, 36, 115, 122, 9, 3, 31, 15, 4, 95, 95, 9, 97,
	99, 9, 100, 100, 37, 101, 122, 9, 3, 31, 9, 2, 95, 95,
	9, 97, 122, 9, 3, 19, 15, 4, 95, 95, 9, 97, 103, 9,
	104, 104, 80, 105, 122, 9, 3, 31, 15, 4, 95, 95, 9, 97,
	113, 9, 114, 114, 68, 115, 122, 9, 3, 31, 15, 4, 95, 95,
	9, 97, 98, 9, 99, 99, 43, 100, 122, 9, 3, 31, 15, 4,
	95, 95, 9, 97, 113, 9, 114, 114, 44, 115, 122, 9, 3, 31,
	15, 4, 95, 95, 9, 97, 115, 9, 116, 116, 45, 117, 122, 9,
	3, 31, 9, 2, 95, 95, 9, 97, 122, 9, 3, 28, 12, 3,
	95, 95, 9, 97, 97, 49, 98, 122, 9, 3, 31, 15, 4, 95,
	95, 9, 97, 113, 9, 114, 114, 50, 115, 122, 9, 3, 31, 15,
	4, 95, 95, 9, 97, 115, 9, 116, 116, 51, 117, 122, 9, 3,
	31, 9, 2, 95, 95, 9, 97, 122, 9, 3, 27, 12, 3, 95,
	95, 9, 97, 97, 53, 98, 122, 9, 3, 31, 15, 4, 95, 95,
	9, 97, 114, 9, 115, 115, 54, 116, 122, 9, 3, 31, 15, 4,
	95, 95, 9, 97, 115, 9, 116, 116, 55, 117, 122, 9, 3, 31,
	9, 2, 95, 95, 9, 97, 122, 9, 3, 18, 21, 6, 95, 95,
	9, 97, 100, 9, 101, 101, 61, 102, 104, 9, 105, 105, 57, 106,
	122, 9, 3, 31, 21, 6, 95, 95, 9, 97, 101, 9, 102, 102,
	58, 103, 119, 9, 120, 120, 65, 121, 122, 9, 3, 31, 15, 4,
	95, 95, 9, 97, 113, 9, 114, 114, 59, 115, 122, 9, 3, 31,
	9, 2, 95, 95, 9, 97, 122, 9, 3, 25, 15, 4, 95, 95,
	9, 97, 110, 9, 111, 111, 69, 112, 122, 9, 3, 31, 15, 4,
	95, 95, 9, 97, 100, 9, 101, 101, 62, 102, 122, 9, 3, 31,
	15, 4, 95, 95, 9, 97, 100, 9, 101, 101, 63, 102, 122, 9,
	3, 31, 9, 2, 95, 95, 9, 97, 122, 9, 3, 24, 15, 4,
	95, 95, 9, 97, 114, 9, 115, 115, 76, 116, 122, 9, 3, 31,
	15, 4, 95, 95, 9, 97, 99, 9, 100, 100, 66, 101, 122, 9,
	3, 31, 15, 4, 95, 95, 9, 97, 100, 9, 101, 101, 67, 102,
	122, 9, 3, 31, 9, 2, 95, 95, 9, 97, 122, 9, 3, 17,
	15, 4, 95, 95, 9, 97, 99, 9, 100, 100, 70, 101, 122, 9,
	3, 31, 15, 4, 95, 95, 9, 97, 113, 9, 114, 114, 71, 115,
	122, 9, 3, 31, 6, 1, 63, 63, 77, 3, 14, 15, 4, 95,
	95, 9, 97, 110, 9, 111, 111, 72, 112, 122, 9, 3, 31, 15,
	4, 95, 95, 9, 97, 100, 9, 101, 101, 73, 102, 122, 9, 3,
	31, 3, 0, 3, 16, 15, 4, 95, 95, 9, 97, 108, 9, 109,
	109, 75, 110, 122, 9, 3, 31, 6, 1, 63, 63, 81, 3, 13,
	9, 2, 95, 95, 78, 97, 122, 9, 3, 31, 3, 0, 3, 15,
	3, 0, 3, 12, 3, 0, 3, 11, 3, 0, 3, 10, 3, 0,
	3, 9, 3, 0, 3, 8, 3, 0, 3, 7, 3, 0, 3, 6,
	3, 0, 3, 5, 3, 0, 3, 4, 3, 0, 3, 3, 3, 0,
	3, 2,
}

var _lexerMode1 = []uint32{
	13, 36, 40, 51, 79, 90, 94, 105, 111, 122, 133, 144, 155, 22,
	7, 0, 9, 1, 11, 44, 1, 45, 45, 5, 46, 91, 1, 92,
	92, 3, 93, 93, 7, 94, 1114111, 1, 3, 0, 3, 38, 10, 3,
	48, 57, 12, 65, 70, 12, 97, 102, 12, 27, 8, 45, 45, 1,
	85, 85, 2, 92, 92, 1, 110, 110, 1, 114, 114, 1, 116, 116,
	1, 117, 117, 9, 120, 120, 6, 3, 38, 10, 3, 48, 57, 1,
	65, 70, 1, 97, 102, 1, 3, 0, 3, 37, 10, 3, 48, 57,
	4, 65, 70, 4, 97, 102, 4, 5, 0, 2, 0, 3, 36, 10,
	3, 48, 57, 6, 65, 70, 6, 97, 102, 6, 10, 3, 48, 57,
	8, 65, 70, 8, 97, 102, 8, 10, 3, 48, 57, 9, 65, 70,
	9, 97, 102, 9, 10, 3, 48, 57, 10, 65, 70, 10, 97, 102,
	10, 10, 3, 48, 57, 11, 65, 70, 11, 97, 102, 11,
}

var _lexerMode2 = []uint32{
	22, 42, 46, 57, 85, 42, 96, 42, 107, 118, 129, 42, 140, 151,
	162, 173, 42, 184, 195, 206, 212, 223, 19, 6, 0, 9, 1, 11,
	38, 1, 39, 39, 19, 40, 91, 1, 92, 92, 3, 93, 1114111, 1,
	3, 0, 5, 0, 10, 3, 48, 57, 21, 65, 70, 21, 97, 102,
	21, 27, 8, 39, 39, 16, 85, 85, 2, 92, 92, 16, 110, 110,
	16, 114, 114, 16, 116, 116, 16, 117, 117, 18, 120, 120, 15, 5,
	0, 10, 3, 48, 57, 5, 65, 70, 5, 97, 102, 5, 10, 3,
	48, 57, 4, 65, 70, 4, 97, 102, 4, 10, 3, 48, 57, 6,
	65, 70, 6, 97, 102, 6, 10, 3, 48, 57, 7, 65, 70, 7,
	97, 102, 7, 10, 3, 48, 57, 9, 65, 70, 9, 97, 102, 9,
	10, 3, 48, 57, 8, 65, 70, 8, 97, 102, 8, 10, 3, 48,
	57, 10, 65, 70, 10, 97, 102, 10, 10, 3, 48, 57, 11, 65,
	70, 11, 97, 102, 11, 10, 3, 48, 57, 14, 65, 70, 14, 97,
	102, 14, 10, 3, 48, 57, 12, 65, 70, 12, 97, 102, 12, 10,
	3, 48, 57, 13, 65, 70, 13, 97, 102, 13, 5, 0, 2, 0,
	3, 34, 10, 3, 48, 57, 17, 65, 70, 17, 97, 102, 17, 10,
	3, 48, 57, 20, 65, 70, 20, 97, 102, 20,
}

var _lexerModes = [][]uint32{

	_lexerMode0,

	_lexerMode1,

	_lexerMode2,
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
