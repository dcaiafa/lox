package parser

import (
	_i0 "errors"
	_i1 "github.com/dcaiafa/lox/internal/ast"
	_i2 "github.com/dcaiafa/lox/internal/base/baselexer"
)

var _LHS = []int32{
	0, 1, 2, 2, 3, 4, 5, 6, 7, 8, 8, 8, 8, 9,
	10, 10, 10, 11, 11, 12, 13, 13, 14, 14, 14, 15, 16, 17,
	18, 19, 20, 21, 22, 22, 22, 23, 23, 23, 23, 24, 24, 25,
	26, 26, 27, 27, 27, 27, 28, 29, 30, 31, 32, 32, 33, 33,
	34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40,
	41, 41, 42, 42, 43, 43, 44, 44, 45, 45, 46, 46, 47, 47,
	48, 48, 49, 49, 50, 50, 51, 51, 52, 52,
}

var _TermCounts = []int32{
	1, 1, 1, 1, 2, 1, 5, 2, 2, 1, 1, 1, 1, 6,
	1, 1, 1, 4, 4, 2, 1, 1, 1, 1, 1, 5, 5, 4,
	5, 1, 1, 2, 1, 1, 1, 1, 1, 1, 3, 3, 1, 4,
	1, 1, 1, 1, 1, 1, 1, 4, 1, 4, 1, 0, 2, 1,
	1, 0, 2, 1, 1, 0, 3, 1, 2, 1, 1, 0, 1, 0,
	1, 0, 2, 1, 1, 0, 2, 1, 1, 0, 2, 1, 3, 1,
	2, 1, 1, 0, 1, 0, 2, 1, 1, 0,
}

var _Actions = []int32{
	133, 140, 155, 166, 173, 180, 187, 190, 197, 200, 203, 206, 217, 228,
	235, 246, 257, 260, 263, 266, 283, 298, 313, 320, 335, 352, 367, 384,
	391, 394, 246, 397, 412, 443, 246, 474, 477, 480, 513, 544, 555, 570,
	601, 626, 651, 654, 665, 680, 689, 544, 698, 709, 712, 723, 726, 737,
	740, 751, 762, 773, 784, 246, 246, 795, 820, 845, 870, 895, 920, 945,
	948, 953, 958, 963, 990, 1017, 1020, 1047, 1074, 1079, 1102, 1119, 1136, 1145,
	1148, 1157, 1160, 1177, 1180, 1211, 1242, 1249, 1256, 1263, 1270, 1275, 1278, 1289,
	680, 1304, 1315, 1318, 1321, 1326, 1331, 1348, 1365, 1382, 1399, 1416, 680, 1433,
	1448, 1465, 1482, 1515, 1522, 1525, 1528, 1531, 1540, 1543, 1546, 1549, 1560, 1571,
	680, 1576, 1579, 1582, 1585, 1590, 1595, 6, 0, -53, 18, 1, 17, 2,
	14, 0, -71, 22, 15, 32, 16, 18, -71, 21, 17, 23, 18, 17,
	-71, 10, 0, -57, 32, -61, 18, -57, 17, -57, 19, 9, 6, 0,
	-3, 18, -3, 17, -3, 6, 0, -2, 18, -2, 17, -2, 6, 0,
	-55, 18, -55, 17, -55, 2, 0, -1, 6, 0, -52, 18, 1, 17,
	2, 2, 0, 2147483647, 2, 32, -60, 2, 32, 28, 10, 0, -5, 32,
	-5, 18, -5, 17, -5, 19, -5, 10, 0, -59, 32, -59, 18, -59,
	17, -59, 19, -59, 6, 0, -4, 18, -4, 17, -4, 10, 0, -56,
	32, -61, 18, -56, 17, -56, 19, 9, 10, 32, 32, 34, 33, 35,
	-89, 9, 34, 8, 35, 2, 4, 30, 2, 32, 44, 2, 32, 29,
	16, 7, -23, 0, -23, 22, -23, 32, -23, 18, -23, 21, -23, 23,
	-23, 17, -23, 14, 0, -21, 22, -21, 32, -21, 18, -21, 21, -21,
	23, -21, 17, -21, 14, 0, -73, 22, -73, 32, -73, 18, -73, 21,
	-73, 23, -73, 17, -73, 6, 0, -19, 18, -19, 17, -19, 14, 0,
	-70, 22, 15, 32, 16, 18, -70, 21, 17, 23, 18, 17, -70, 16,
	7, -24, 0, -24, 22, -24, 32, -24, 18, -24, 21, -24, 23, -24,
	17, -24, 14, 0, -20, 22, -20, 32, -20, 18, -20, 21, -20, 23,
	-20, 17, -20, 16, 7, -22, 0, -22, 22, -22, 32, -22, 18, -22,
	21, -22, 23, -22, 17, -22, 6, 0, -54, 18, -54, 17, -54, 2,
	4, 47, 2, 6, 48, 14, 10, -29, 20, -29, 30, -29, 5, 62,
	25, -29, 24, -29, 2, -29, 30, 10, -36, 20, -36, 30, -36, 32,
	-36, 34, -36, 35, -36, 14, -36, 9, -36, 5, -36, 25, -36, 24,
	-36, 2, -36, 8, -36, 13, -36, 12, -36, 30, 10, -35, 20, -35,
	30, -35, 32, -35, 34, -35, 35, -35, 14, -35, 9, -35, 5, -35,
	25, -35, 24, -35, 2, -35, 8, -35, 13, -35, 12, -35, 2, 35,
	-88, 2, 35, 71, 32, 10, -40, 20, -40, 30, -40, 32, -40, 34,
	-40, 35, -40, 14, -40, 9, -40, 5, -40, 25, -40, 24, -40, 2,
	-40, 11, 70, 8, -40, 13, -40, 12, -40, 30, 10, -37, 20, -37,
	30, -37, 32, -37, 34, -37, 35, -37, 14, -37, 9, -37, 5, -37,
	25, -37, 24, -37, 2, -37, 8, -37, 13, -37, 12, -37, 10, 20,
	50, 30, 51, 25, 52, 24, 53, 2, -79, 14, 10, -83, 20, -83,
	30, -83, 5, -83, 25, -83, 24, -83, 2, -83, 30, 10, -87, 20,
	-87, 30, -87, 32, -87, 34, -87, 35, -87, 14, 64, 9, -87, 5,
	-87, 25, -87, 24, -87, 2, -87, 8, -87, 13, 65, 12, 66, 24,
	10, -85, 20, -85, 30, -85, 32, -85, 34, -85, 35, -85, 9, -85,
	5, -85, 25, -85, 24, -85, 2, -85, 8, -85, 24, 10, -30, 20,
	-30, 30, -30, 32, 32, 34, 33, 35, -89, 9, 34, 5, -30, 25,
	-30, 24, -30, 2, -30, 8, 35, 2, 4, 61, 10, 0, -58, 32,
	-58, 18, -58, 17, -58, 19, -58, 14, 0, -72, 22, -72, 32, -72,
	18, -72, 21, -72, 23, -72, 17, -72, 8, 26, 73, 32, 74, 28,
	75, 34, 76, 8, 7, -75, 22, 15, 32, 16, 21, 17, 10, 20,
	-48, 30, -48, 25, -48, 24, -48, 2, -48, 2, 9, 95, 10, 20,
	-50, 30, -50, 25, -50, 24, -50, 2, -50, 2, 9, 94, 10, 20,
	-81, 30, -81, 25, -81, 24, -81, 2, -81, 2, 2, 86, 10, 20,
	50, 30, 51, 25, 52, 24, 53, 2, -78, 10, 20, -44, 30, -44,
	25, -44, 24, -44, 2, -44, 10, 20, -47, 30, -47, 25, -47, 24,
	-47, 2, -47, 10, 20, -46, 30, -46, 25, -46, 24, -46, 2, -46,
	10, 20, -45, 30, -45, 25, -45, 24, -45, 2, -45, 24, 10, -84,
	20, -84, 30, -84, 32, -84, 34, -84, 35, -84, 9, -84, 5, -84,
	25, -84, 24, -84, 2, -84, 8, -84, 24, 10, -34, 20, -34, 30,
	-34, 32, -34, 34, -34, 35, -34, 9, -34, 5, -34, 25, -34, 24,
	-34, 2, -34, 8, -34, 24, 10, -33, 20, -33, 30, -33, 32, -33,
	34, -33, 35, -33, 9, -33, 5, -33, 25, -33, 24, -33, 2, -33,
	8, -33, 24, 10, -32, 20, -32, 30, -32, 32, -32, 34, -32, 35,
	-32, 9, -32, 5, -32, 25, -32, 24, -32, 2, -32, 8, -32, 24,
	10, -86, 20, -86, 30, -86, 32, -86, 34, -86, 35, -86, 9, -86,
	5, -86, 25, -86, 24, -86, 2, -86, 8, -86, 24, 10, -31, 20,
	-31, 30, -31, 32, -31, 34, -31, 35, -31, 9, -31, 5, -31, 25,
	-31, 24, -31, 2, -31, 8, -31, 2, 10, 88, 4, 35, -89, 8,
	35, 4, 38, 90, 37, 91, 4, 5, 98, 2, 99, 26, 3, -11,
	10, -11, 26, -11, 32, -11, 27, -11, 28, -11, 34, -11, 14, -11,
	5, -11, 29, -11, 2, -11, 13, -11, 12, -11, 26, 3, -9, 10,
	-9, 26, -9, 32, -9, 27, -9, 28, -9, 34, -9, 14, -9, 5,
	-9, 29, -9, 2, -9, 13, -9, 12, -9, 2, 9, 110, 26, 3,
	-10, 10, -10, 26, -10, 32, -10, 27, -10, 28, -10, 34, -10, 14,
	-10, 5, -10, 29, -10, 2, -10, 13, -10, 12, -10, 26, 3, -12,
	10, -12, 26, -12, 32, -12, 27, -12, 28, -12, 34, -12, 14, -12,
	5, -12, 29, -12, 2, -12, 13, -12, 12, -12, 4, 5, -63, 2,
	-63, 22, 26, -69, 32, -69, 27, -69, 28, -69, 34, -69, 14, 105,
	5, -69, 29, -69, 2, -69, 13, 106, 12, 107, 16, 26, -65, 32,
	-65, 27, -65, 28, -65, 34, -65, 5, -65, 29, -65, 2, -65, 16,
	26, 73, 32, 74, 27, 100, 28, 75, 34, 76, 5, -67, 29, 101,
	2, -67, 8, 7, -77, 22, -77, 32, -77, 21, -77, 2, 7, 111,
	8, 7, -74, 22, 15, 32, 16, 21, 17, 2, 2, 112, 16, 7,
	-27, 0, -27, 22, -27, 32, -27, 18, -27, 21, -27, 23, -27, 17,
	-27, 2, 2, 113, 30, 10, -38, 20, -38, 30, -38, 32, -38, 34,
	-38, 35, -38, 14, -38, 9, -38, 5, -38, 25, -38, 24, -38, 2,
	-38, 8, -38, 13, -38, 12, -38, 30, 10, -39, 20, -39, 30, -39,
	32, -39, 34, -39, 35, -39, 14, -39, 9, -39, 5, -39, 25, -39,
	24, -39, 2, -39, 8, -39, 13, -39, 12, -39, 6, 36, -42, 38,
	-42, 37, -42, 6, 36, -43, 38, -43, 37, -43, 6, 36, -91, 38,
	-91, 37, -91, 6, 36, 114, 38, 90, 37, 91, 4, 10, -93, 32,
	116, 2, 32, 118, 10, 20, -80, 30, -80, 25, -80, 24, -80, 2,
	-80, 14, 10, -82, 20, -82, 30, -82, 5, -82, 25, -82, 24, -82,
	2, -82, 10, 0, -6, 32, -6, 18, -6, 17, -6, 19, -6, 2,
	9, 121, 2, 9, 122, 4, 5, -66, 2, -66, 4, 5, -7, 2,
	-7, 16, 26, -64, 32, -64, 27, -64, 28, -64, 34, -64, 5, -64,
	29, -64, 2, -64, 16, 26, -15, 32, -15, 27, -15, 28, -15, 34,
	-15, 5, -15, 29, -15, 2, -15, 16, 26, -14, 32, -14, 27, -14,
	28, -14, 34, -14, 5, -14, 29, -14, 2, -14, 16, 26, -16, 32,
	-16, 27, -16, 28, -16, 34, -16, 5, -16, 29, -16, 2, -16, 16,
	26, -68, 32, -68, 27, -68, 28, -68, 34, -68, 5, -68, 29, -68,
	2, -68, 16, 26, -8, 32, -8, 27, -8, 28, -8, 34, -8, 5,
	-8, 29, -8, 2, -8, 14, 0, -25, 22, -25, 32, -25, 18, -25,
	21, -25, 23, -25, 17, -25, 16, 7, -26, 0, -26, 22, -26, 32,
	-26, 18, -26, 21, -26, 23, -26, 17, -26, 16, 7, -28, 0, -28,
	22, -28, 32, -28, 18, -28, 21, -28, 23, -28, 17, -28, 32, 10,
	-41, 20, -41, 30, -41, 32, -41, 34, -41, 35, -41, 14, -41, 9,
	-41, 5, -41, 25, -41, 24, -41, 2, -41, 11, -41, 8, -41, 13,
	-41, 12, -41, 6, 36, -90, 38, -90, 37, -90, 2, 10, -92, 2,
	10, 123, 2, 10, 124, 8, 7, -76, 22, -76, 32, -76, 21, -76,
	2, 3, 126, 2, 33, 127, 2, 33, 128, 10, 20, -49, 30, -49,
	25, -49, 24, -49, 2, -49, 10, 20, -51, 30, -51, 25, -51, 24,
	-51, 2, -51, 4, 5, -62, 2, -62, 2, 10, 130, 2, 10, 131,
	2, 10, 132, 4, 5, -17, 2, -17, 4, 5, -18, 2, -18, 26,
	3, -13, 10, -13, 26, -13, 32, -13, 27, -13, 28, -13, 34, -13,
	14, -13, 5, -13, 29, -13, 2, -13, 13, -13, 12, -13,
}

var _Goto = []int32{
	133, 146, 163, 174, 174, 174, 174, 175, 174, 174, 174, 174, 174, 174,
	182, 189, 174, 174, 174, 174, 174, 174, 174, 208, 174, 174, 174, 174,
	174, 174, 221, 174, 174, 174, 240, 174, 174, 174, 174, 259, 174, 274,
	174, 279, 174, 174, 174, 290, 303, 316, 174, 174, 174, 174, 174, 174,
	331, 174, 174, 174, 174, 342, 361, 174, 174, 174, 174, 174, 174, 174,
	376, 381, 174, 174, 174, 174, 174, 174, 174, 386, 174, 391, 174, 174,
	402, 174, 174, 174, 174, 174, 174, 174, 174, 411, 414, 174, 174, 174,
	417, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 428, 174,
	174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174,
	433, 174, 174, 174, 174, 174, 174, 12, 12, 3, 3, 4, 2, 5,
	32, 6, 33, 7, 1, 8, 16, 17, 19, 14, 20, 13, 21, 41,
	22, 42, 23, 18, 24, 15, 25, 16, 26, 10, 36, 10, 5, 11,
	4, 12, 34, 13, 35, 14, 0, 6, 12, 3, 3, 4, 2, 27,
	6, 36, 10, 5, 11, 4, 45, 18, 47, 31, 50, 36, 25, 37,
	24, 38, 19, 39, 20, 40, 23, 41, 21, 42, 48, 43, 12, 17,
	19, 14, 20, 13, 46, 18, 24, 15, 25, 16, 26, 18, 47, 31,
	50, 36, 25, 37, 24, 38, 19, 49, 20, 40, 23, 41, 21, 42,
	48, 43, 18, 47, 31, 50, 36, 25, 37, 24, 38, 19, 69, 20,
	40, 23, 41, 21, 42, 48, 43, 14, 27, 54, 45, 55, 46, 56,
	28, 57, 31, 58, 30, 59, 29, 60, 4, 22, 67, 49, 68, 10,
	50, 36, 25, 37, 24, 38, 23, 41, 21, 63, 12, 37, 72, 9,
	77, 6, 78, 8, 79, 7, 80, 38, 81, 12, 17, 19, 14, 82,
	43, 83, 44, 84, 18, 24, 16, 26, 14, 27, 54, 45, 85, 46,
	56, 28, 57, 31, 58, 30, 59, 29, 60, 10, 27, 96, 28, 57,
	31, 58, 30, 59, 29, 60, 18, 47, 31, 50, 36, 25, 37, 24,
	38, 19, 87, 20, 40, 23, 41, 21, 42, 48, 43, 14, 50, 36,
	25, 37, 24, 38, 20, 97, 23, 41, 21, 42, 48, 43, 4, 50,
	36, 25, 89, 4, 26, 92, 51, 93, 4, 10, 108, 40, 109, 10,
	9, 77, 11, 102, 39, 103, 8, 79, 7, 104, 8, 17, 19, 14,
	119, 18, 24, 16, 26, 2, 26, 115, 2, 52, 117, 10, 9, 77,
	6, 125, 8, 79, 7, 80, 38, 81, 4, 9, 77, 8, 120, 4,
	9, 77, 8, 129,
}

type _Bounds struct {
	Begin Token
	End   Token
	Empty bool
}

func _cast[T any](v any) T {
	cv, _ := v.(T)
	return cv
}

var _errorPlaceholder = _i0.New("error placeholder")

func _Find(table []int32, y, x int32) (int32, bool) {
	i := int(table[int(y)])
	count := int(table[i])
	i++
	end := i + count
	for ; i < end; i += 2 {
		if table[i] == x {
			return table[i+1], true
		}
	}
	return 0, false
}

type _Lexer interface {
	ReadToken() (Token, int)
}

type lox struct {
	_lex    _Lexer
	_state  _Stack[int32]
	_sym    _Stack[any]
	_bounds _Stack[_Bounds]

	_lookahead     Token
	_lookaheadType int
	_errorToken    Token
}

func (p *parser) parse(lex _Lexer) bool {
	const accept = 2147483647

	p._lex = lex

	p._state.Push(0)
	p._ReadToken()

	for {
		if p._lookaheadType == ERROR {
			_, ok := p._Recover()
			if !ok {
				return false
			}
		}
		topState := p._state.Peek(0)
		action, ok := _Find(
			_Actions, topState, int32(p._lookaheadType))
		if !ok {
			action, ok = p._Recover()
			if !ok {
				return false
			}
		}
		if action == accept {
			break
		} else if action >= 0 { // shift
			p._state.Push(action)
			p._sym.Push(p._lookahead)
			p._bounds.Push(
				_Bounds{Begin: p._lookahead,
					End: p._lookahead})
			p._ReadToken()
		} else { // reduce
			prod := -action
			termCount := _TermCounts[int(prod)]
			rule := _LHS[int(prod)]
			res := p._Act(prod)

			// Compute reduction token bounds.
			// Trim leading and trailing empty bounds.
			boundSlice := p._bounds.PeekSlice(int(termCount))
			for len(boundSlice) > 0 && boundSlice[0].Empty {
				boundSlice = boundSlice[1:]
			}
			for len(boundSlice) > 0 && boundSlice[len(boundSlice)-1].Empty {
				boundSlice = boundSlice[:len(boundSlice)-1]
			}
			var bounds _Bounds
			if len(boundSlice) > 0 {
				bounds.Begin = boundSlice[0].Begin
				bounds.End = boundSlice[len(boundSlice)-1].End
			} else {
				bounds.Empty = true
			}
			if !bounds.Empty {
				p._onBounds(res, bounds.Begin, bounds.End)
			}
			p._bounds.Pop(int(termCount))
			p._bounds.Push(bounds)
			p._state.Pop(int(termCount))
			p._sym.Pop(int(termCount))
			topState = p._state.Peek(0)
			nextState, _ := _Find(_Goto, topState, rule)
			p._state.Push(nextState)
			p._sym.Push(res)
		}
	}

	return true
}

func (p *parser) errorToken() Token {
	return p._errorToken
}

func (p *parser) _ReadToken() {
	p._lookahead, p._lookaheadType = p._lex.ReadToken()
}

func (p *parser) _Recover() (int32, bool) {
	p._errorToken = p._lookahead
	p._onError()

	for {
		for p._lookaheadType == ERROR {
			p._ReadToken()
		}

		saveState := p._state
		saveSym := p._sym
		saveBounds := p._bounds

		for len(p._state) > 1 {
			topState := p._state.Peek(0)
			action, ok := _Find(_Actions, topState, int32(ERROR))
			if ok {
				action2, ok := _Find(
					_Actions, action, int32(p._lookaheadType))
				if ok {
					p._state.Push(action)
					p._sym.Push(_errorPlaceholder)
					p._bounds.Push(_Bounds{})
					return action2, true
				}
			}
			p._state.Pop(1)
			p._sym.Pop(1)
			p._bounds.Pop(1)
		}

		if p._lookaheadType == EOF {
			return 0, false
		}

		p._ReadToken()
		p._state = saveState
		p._sym = saveSym
		p._bounds = saveBounds
	}
}

func (p *parser) _Act(prod int32) any {
	switch prod {
	case 1:
		return p.on_spec(
			_cast[[][]_i1.Statement](p._sym.Peek(0)),
		)
	case 2:
		return p.on_section(
			_cast[[]_i1.Statement](p._sym.Peek(0)),
		)
	case 3:
		return p.on_section(
			_cast[[]_i1.Statement](p._sym.Peek(0)),
		)
	case 4:
		return p.on_parser_section(
			_cast[_i2.Token](p._sym.Peek(1)),
			_cast[[]_i1.Statement](p._sym.Peek(0)),
		)
	case 5:
		return p.on_parser_statement(
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 6:
		return p.on_parser_rule(
			_cast[_i2.Token](p._sym.Peek(4)),
			_cast[_i2.Token](p._sym.Peek(3)),
			_cast[_i2.Token](p._sym.Peek(2)),
			_cast[[]*_i1.ParserProd](p._sym.Peek(1)),
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 7:
		return p.on_parser_prod(
			_cast[[]*_i1.ParserTerm](p._sym.Peek(1)),
			_cast[*_i1.ProdQualifier](p._sym.Peek(0)),
		)
	case 8:
		return p.on_parser_term_card(
			_cast[*_i1.ParserTerm](p._sym.Peek(1)),
			_cast[_i1.ParserTermType](p._sym.Peek(0)),
		)
	case 9:
		return p.on_parser_term__token(
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 10:
		return p.on_parser_term__token(
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 11:
		return p.on_parser_term__token(
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 12:
		return p.on_parser_term__list(
			_cast[*_i1.ParserTerm](p._sym.Peek(0)),
		)
	case 13:
		return p.on_parser_list(
			_cast[_i2.Token](p._sym.Peek(5)),
			_cast[_i2.Token](p._sym.Peek(4)),
			_cast[*_i1.ParserTerm](p._sym.Peek(3)),
			_cast[_i2.Token](p._sym.Peek(2)),
			_cast[*_i1.ParserTerm](p._sym.Peek(1)),
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 14:
		return p.on_parser_card(
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 15:
		return p.on_parser_card(
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 16:
		return p.on_parser_card(
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 17:
		return p.on_parser_qualif(
			_cast[_i2.Token](p._sym.Peek(3)),
			_cast[_i2.Token](p._sym.Peek(2)),
			_cast[_i2.Token](p._sym.Peek(1)),
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 18:
		return p.on_parser_qualif(
			_cast[_i2.Token](p._sym.Peek(3)),
			_cast[_i2.Token](p._sym.Peek(2)),
			_cast[_i2.Token](p._sym.Peek(1)),
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 19:
		return p.on_lexer_section(
			_cast[_i2.Token](p._sym.Peek(1)),
			_cast[[]_i1.Statement](p._sym.Peek(0)),
		)
	case 20:
		return p.on_lexer_statement(
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 21:
		return p.on_lexer_statement(
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 22:
		return p.on_lexer_rule(
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 23:
		return p.on_lexer_rule(
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 24:
		return p.on_lexer_rule(
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 25:
		return p.on_mode(
			_cast[_i2.Token](p._sym.Peek(4)),
			_cast[_i2.Token](p._sym.Peek(3)),
			_cast[_i2.Token](p._sym.Peek(2)),
			_cast[[]_i1.Statement](p._sym.Peek(1)),
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 26:
		return p.on_token_rule(
			_cast[_i2.Token](p._sym.Peek(4)),
			_cast[_i2.Token](p._sym.Peek(3)),
			_cast[*_i1.LexerExpr](p._sym.Peek(2)),
			_cast[[]_i1.Action](p._sym.Peek(1)),
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 27:
		return p.on_frag_rule(
			_cast[_i2.Token](p._sym.Peek(3)),
			_cast[*_i1.LexerExpr](p._sym.Peek(2)),
			_cast[[]_i1.Action](p._sym.Peek(1)),
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 28:
		return p.on_macro_rule(
			_cast[_i2.Token](p._sym.Peek(4)),
			_cast[_i2.Token](p._sym.Peek(3)),
			_cast[_i2.Token](p._sym.Peek(2)),
			_cast[*_i1.LexerExpr](p._sym.Peek(1)),
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 29:
		return p.on_lexer_expr(
			_cast[[]*_i1.LexerFactor](p._sym.Peek(0)),
		)
	case 30:
		return p.on_lexer_factor(
			_cast[[]*_i1.LexerTermCard](p._sym.Peek(0)),
		)
	case 31:
		return p.on_lexer_term_card(
			_cast[_i1.LexerTerm](p._sym.Peek(1)),
			_cast[_i1.Card](p._sym.Peek(0)),
		)
	case 32:
		return p.on_lexer_card(
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 33:
		return p.on_lexer_card(
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 34:
		return p.on_lexer_card(
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 35:
		return p.on_lexer_term__tok(
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 36:
		return p.on_lexer_term__tok(
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 37:
		return p.on_lexer_term__char_class_expr(
			_cast[_i1.CharClassExpr](p._sym.Peek(0)),
		)
	case 38:
		return p.on_lexer_term__expr(
			_cast[_i2.Token](p._sym.Peek(2)),
			_cast[*_i1.LexerExpr](p._sym.Peek(1)),
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 39:
		return p.on_char_class_expr__binary(
			_cast[_i1.CharClassExpr](p._sym.Peek(2)),
			_cast[_i2.Token](p._sym.Peek(1)),
			_cast[_i1.CharClassExpr](p._sym.Peek(0)),
		)
	case 40:
		return p.on_char_class_expr__char_class(
			_cast[*_i1.CharClass](p._sym.Peek(0)),
		)
	case 41:
		return p.on_char_class(
			_cast[_i2.Token](p._sym.Peek(3)),
			_cast[_i2.Token](p._sym.Peek(2)),
			_cast[[]_i2.Token](p._sym.Peek(1)),
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 42:
		return p.on_char_class_item(
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 43:
		return p.on_char_class_item(
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 44:
		return p.on_action(
			_cast[_i1.Action](p._sym.Peek(0)),
		)
	case 45:
		return p.on_action(
			_cast[_i1.Action](p._sym.Peek(0)),
		)
	case 46:
		return p.on_action(
			_cast[_i1.Action](p._sym.Peek(0)),
		)
	case 47:
		return p.on_action(
			_cast[_i1.Action](p._sym.Peek(0)),
		)
	case 48:
		return p.on_action_discard(
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 49:
		return p.on_action_push_mode(
			_cast[_i2.Token](p._sym.Peek(3)),
			_cast[_i2.Token](p._sym.Peek(2)),
			_cast[_i2.Token](p._sym.Peek(1)),
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 50:
		return p.on_action_pop_mode(
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 51:
		return p.on_action_emit(
			_cast[_i2.Token](p._sym.Peek(3)),
			_cast[_i2.Token](p._sym.Peek(2)),
			_cast[_i2.Token](p._sym.Peek(1)),
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 52: // ZeroOrMore
		return _cast[[][]_i1.Statement](p._sym.Peek(0))
	case 53: // ZeroOrMore
		{
			var zero [][]_i1.Statement
			return zero
		}
	case 54: // OneOrMore
		return append(
			_cast[[][]_i1.Statement](p._sym.Peek(1)),
			_cast[[]_i1.Statement](p._sym.Peek(0)),
		)
	case 55: // OneOrMore
		return [][]_i1.Statement{
			_cast[[]_i1.Statement](p._sym.Peek(0)),
		}
	case 56: // ZeroOrMore
		return _cast[[]_i1.Statement](p._sym.Peek(0))
	case 57: // ZeroOrMore
		{
			var zero []_i1.Statement
			return zero
		}
	case 58: // OneOrMore
		return append(
			_cast[[]_i1.Statement](p._sym.Peek(1)),
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 59: // OneOrMore
		return []_i1.Statement{
			_cast[_i1.Statement](p._sym.Peek(0)),
		}
	case 60: // ZeroOrOne
		return _cast[_i2.Token](p._sym.Peek(0))
	case 61: // ZeroOrOne
		{
			var zero _i2.Token
			return zero
		}
	case 62: // List
		return append(
			_cast[[]*_i1.ParserProd](p._sym.Peek(2)),
			_cast[*_i1.ParserProd](p._sym.Peek(0)),
		)
	case 63: // List
		return []*_i1.ParserProd{
			_cast[*_i1.ParserProd](p._sym.Peek(0)),
		}
	case 64: // OneOrMore
		return append(
			_cast[[]*_i1.ParserTerm](p._sym.Peek(1)),
			_cast[*_i1.ParserTerm](p._sym.Peek(0)),
		)
	case 65: // OneOrMore
		return []*_i1.ParserTerm{
			_cast[*_i1.ParserTerm](p._sym.Peek(0)),
		}
	case 66: // ZeroOrOne
		return _cast[*_i1.ProdQualifier](p._sym.Peek(0))
	case 67: // ZeroOrOne
		{
			var zero *_i1.ProdQualifier
			return zero
		}
	case 68: // ZeroOrOne
		return _cast[_i1.ParserTermType](p._sym.Peek(0))
	case 69: // ZeroOrOne
		{
			var zero _i1.ParserTermType
			return zero
		}
	case 70: // ZeroOrMore
		return _cast[[]_i1.Statement](p._sym.Peek(0))
	case 71: // ZeroOrMore
		{
			var zero []_i1.Statement
			return zero
		}
	case 72: // OneOrMore
		return append(
			_cast[[]_i1.Statement](p._sym.Peek(1)),
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 73: // OneOrMore
		return []_i1.Statement{
			_cast[_i1.Statement](p._sym.Peek(0)),
		}
	case 74: // ZeroOrMore
		return _cast[[]_i1.Statement](p._sym.Peek(0))
	case 75: // ZeroOrMore
		{
			var zero []_i1.Statement
			return zero
		}
	case 76: // OneOrMore
		return append(
			_cast[[]_i1.Statement](p._sym.Peek(1)),
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 77: // OneOrMore
		return []_i1.Statement{
			_cast[_i1.Statement](p._sym.Peek(0)),
		}
	case 78: // ZeroOrMore
		return _cast[[]_i1.Action](p._sym.Peek(0))
	case 79: // ZeroOrMore
		{
			var zero []_i1.Action
			return zero
		}
	case 80: // OneOrMore
		return append(
			_cast[[]_i1.Action](p._sym.Peek(1)),
			_cast[_i1.Action](p._sym.Peek(0)),
		)
	case 81: // OneOrMore
		return []_i1.Action{
			_cast[_i1.Action](p._sym.Peek(0)),
		}
	case 82: // List
		return append(
			_cast[[]*_i1.LexerFactor](p._sym.Peek(2)),
			_cast[*_i1.LexerFactor](p._sym.Peek(0)),
		)
	case 83: // List
		return []*_i1.LexerFactor{
			_cast[*_i1.LexerFactor](p._sym.Peek(0)),
		}
	case 84: // OneOrMore
		return append(
			_cast[[]*_i1.LexerTermCard](p._sym.Peek(1)),
			_cast[*_i1.LexerTermCard](p._sym.Peek(0)),
		)
	case 85: // OneOrMore
		return []*_i1.LexerTermCard{
			_cast[*_i1.LexerTermCard](p._sym.Peek(0)),
		}
	case 86: // ZeroOrOne
		return _cast[_i1.Card](p._sym.Peek(0))
	case 87: // ZeroOrOne
		{
			var zero _i1.Card
			return zero
		}
	case 88: // ZeroOrOne
		return _cast[_i2.Token](p._sym.Peek(0))
	case 89: // ZeroOrOne
		{
			var zero _i2.Token
			return zero
		}
	case 90: // OneOrMore
		return append(
			_cast[[]_i2.Token](p._sym.Peek(1)),
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 91: // OneOrMore
		return []_i2.Token{
			_cast[_i2.Token](p._sym.Peek(0)),
		}
	case 92: // ZeroOrOne
		return _cast[_i2.Token](p._sym.Peek(0))
	case 93: // ZeroOrOne
		{
			var zero _i2.Token
			return zero
		}
	default:
		panic("unreachable")
	}
}
