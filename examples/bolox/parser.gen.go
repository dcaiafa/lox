package main

import (
	_i0 "github.com/dcaiafa/loxlex/simplelexer"
)

var _rules = []int32{
	0, 1, 1, 2, 3, 3, 3, 3, 3, 3, 4, 5, 6, 7,
	8, 9, 10, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 12, 12, 12, 13, 14, 14, 14, 14, 14, 15, 16, 16,
	17, 17, 18, 18, 19, 19, 20, 20, 21, 21, 22, 22, 23, 23,
	24, 24, 25, 25,
}

var _termCounts = []int32{
	1, 1, 1, 1, 2, 2, 2, 2, 2, 1, 5, 1, 7, 5,
	4, 3, 4, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 3,
	1, 0, 2, 1, 1, 0, 2, 1, 1, 0, 1, 0, 3, 1,
	1, 0, 2, 1,
}

var _actions = []int32{
	100, 115, 118, 121, 126, 141, 126, 156, 159, 162, 165, 168, 171, 186,
	191, 206, 209, 212, 227, 242, 257, 272, 287, 320, 355, 388, 126, 421,
	428, 461, 486, 519, 552, 585, 618, 651, 126, 676, 693, 126, 126, 126,
	126, 126, 126, 126, 126, 708, 126, 126, 126, 708, 721, 746, 751, 754,
	781, 806, 126, 813, 820, 823, 830, 833, 836, 869, 902, 935, 968, 1001,
	1034, 1067, 1100, 1133, 1166, 1199, 1232, 1265, 1298, 126, 1323, 1330, 1333, 1340,
	1347, 126, 1374, 1381, 1386, 1393, 1396, 1399, 1402, 1427, 708, 708, 1434, 1437,
	1440, 1447, 14, 11, 1, 0, -43, 1, 2, 12, 3, 7, 4, 32,
	5, 2, 6, 2, 32, 21, 2, 0, -2, 4, 16, 36, 13, 37,
	14, 6, 22, 12, 23, 28, 24, 10, 25, 13, 26, 29, 27, 5,
	28, 14, 18, -9, 11, -9, 0, -9, 12, -9, 7, -9, 32, -9,
	2, -9, 2, 0, -1, 2, 32, -11, 2, 32, 17, 2, 32, 20,
	2, 0, 2147483647, 14, 18, -45, 11, -45, 0, -45, 12, -45, 7, -45,
	32, -45, 2, -45, 4, 18, -3, 0, -3, 14, 18, -42, 11, 1,
	0, -42, 12, 3, 7, 4, 32, 5, 2, 6, 2, 32, 18, 2,
	32, 19, 14, 18, -4, 11, -4, 0, -4, 12, -4, 7, -4, 32,
	-4, 2, -4, 14, 18, -5, 11, -5, 0, -5, 12, -5, 7, -5,
	32, -5, 2, -5, 14, 18, -6, 11, -6, 0, -6, 12, -6, 7,
	-6, 32, -6, 2, -6, 14, 18, -7, 11, -7, 0, -7, 12, -7,
	7, -7, 32, -7, 2, -7, 14, 18, -8, 11, -8, 0, -8, 12,
	-8, 7, -8, 32, -8, 2, -8, 32, 4, -36, 18, -36, 15, -36,
	14, -36, 22, -36, 27, -36, 26, -36, 25, -36, 24, -36, 23, -36,
	20, -36, 32, -36, 17, -36, 3, -36, 19, -36, 21, -36, 34, 4,
	-33, 18, -33, 15, -33, 14, -33, 22, -33, 27, -33, 26, -33, 25,
	-33, 24, -33, 23, -33, 20, -33, 32, -33, 17, -33, 13, 37, 3,
	-33, 19, -33, 21, -33, 32, 4, -34, 18, -34, 15, -34, 14, -34,
	22, -34, 27, -34, 26, -34, 25, -34, 24, -34, 23, -34, 20, -34,
	32, -34, 17, -34, 3, -34, 19, -34, 21, -34, 32, 4, -37, 18,
	-37, 15, -37, 14, -37, 22, -37, 27, -37, 26, -37, 25, -37, 24,
	-37, 23, -37, 20, -37, 32, -37, 17, -37, 3, -37, 19, -37, 21,
	-37, 6, 31, 57, 17, 58, 30, -57, 32, 4, -35, 18, -35, 15,
	-35, 14, -35, 22, -35, 27, -35, 26, -35, 25, -35, 24, -35, 23,
	-35, 20, -35, 32, -35, 17, -35, 3, -35, 19, -35, 21, -35, 24,
	4, 39, 22, 40, 27, 41, 26, 42, 25, 43, 24, 44, 23, 45,
	20, 46, 17, 47, 3, 48, 19, 49, 21, 50, 32, 4, -30, 18,
	-30, 15, -30, 14, -30, 22, -30, 27, -30, 26, -30, 25, -30, 24,
	-30, 23, -30, 20, -30, 32, -30, 17, -30, 3, -30, 19, -30, 21,
	-30, 32, 4, -31, 18, -31, 15, -31, 14, -31, 22, -31, 27, -31,
	26, -31, 25, -31, 24, -31, 23, -31, 20, -31, 32, -31, 17, -31,
	3, -31, 19, -31, 21, -31, 32, 4, -29, 18, -29, 15, -29, 14,
	-29, 22, -29, 27, -29, 26, -29, 25, -29, 24, -29, 23, -29, 20,
	-29, 32, -29, 17, -29, 3, -29, 19, -29, 21, -29, 32, 4, -38,
	18, -38, 15, -38, 14, -38, 22, -38, 27, -38, 26, -38, 25, -38,
	24, -38, 23, -38, 20, -38, 32, -38, 17, -38, 3, -38, 19, -38,
	21, -38, 32, 4, -32, 18, -32, 15, -32, 14, -32, 22, -32, 27,
	-32, 26, -32, 25, -32, 24, -32, 23, -32, 20, -32, 32, -32, 17,
	-32, 3, -32, 19, -32, 21, -32, 24, 4, 39, 22, 40, 27, 41,
	26, 42, 25, 43, 24, 44, 23, 45, 20, 46, 17, 51, 3, 48,
	19, 49, 21, 50, 16, 14, -53, 6, 22, 12, 23, 28, 24, 10,
	25, 13, 26, 29, 27, 5, 28, 14, 18, -44, 11, -44, 0, -44,
	12, -44, 7, -44, 32, -44, 2, -44, 12, 18, -43, 11, 1, 12,
	3, 7, 4, 32, 5, 2, 6, 24, 4, 39, 22, 40, 27, 41,
	26, 42, 25, 43, 24, 44, 23, 45, 20, 46, 32, -15, 3, 48,
	19, 49, 21, 50, 4, 15, 79, 14, -52, 2, 14, 64, 26, 4,
	39, 15, -55, 14, -55, 22, 40, 27, 41, 26, 42, 25, 43, 24,
	44, 23, 45, 20, 46, 3, 48, 19, 49, 21, 50, 24, 4, 39,
	14, 65, 22, 40, 27, 41, 26, 42, 25, 43, 24, 44, 23, 45,
	20, 46, 3, 48, 19, 49, 21, 50, 6, 31, -40, 17, -40, 30,
	-40, 6, 31, -59, 17, -59, 30, -59, 2, 30, 77, 6, 31, 57,
	17, 58, 30, -56, 2, 18, 81, 2, 18, 82, 32, 4, -16, 18,
	-16, 15, -16, 14, -16, 22, -16, 27, -16, 26, -16, 25, -16, 24,
	-16, 23, -16, 20, -16, 32, -16, 17, -16, 3, -16, 19, -16, 21,
	-16, 32, 4, -28, 18, -28, 15, -28, 14, -28, 22, -28, 27, -28,
	26, -28, 25, -28, 24, -28, 23, -28, 20, -28, 32, -28, 17, -28,
	3, -28, 19, -28, 21, -28, 32, 4, -17, 18, -17, 15, -17, 14,
	-17, 22, -17, 27, -17, 26, -17, 25, -17, 24, -17, 23, -17, 20,
	-17, 32, -17, 17, -17, 3, -17, 19, -17, 21, -17, 32, 4, -18,
	18, -18, 15, -18, 14, -18, 22, -18, 27, -18, 26, -18, 25, -18,
	24, -18, 23, -18, 20, -18, 32, -18, 17, -18, 3, -18, 19, -18,
	21, -18, 32, 4, -19, 18, -19, 15, -19, 14, -19, 22, 40, 27,
	-19, 26, -19, 25, -19, 24, -19, 23, -19, 20, -19, 32, -19, 17,
	-19, 3, -19, 19, -19, 21, 50, 32, 4, -20, 18, -20, 15, -20,
	14, -20, 22, 40, 27, -20, 26, -20, 25, -20, 24, -20, 23, -20,
	20, -20, 32, -20, 17, -20, 3, -20, 19, -20, 21, 50, 32, 4,
	-21, 18, -21, 15, -21, 14, -21, 22, 40, 27, -21, 26, -21, 25,
	-21, 24, -21, 23, -21, 20, 46, 32, -21, 17, -21, 3, -21, 19,
	49, 21, 50, 32, 4, -22, 18, -22, 15, -22, 14, -22, 22, 40,
	27, -22, 26, -22, 25, -22, 24, -22, 23, -22, 20, 46, 32, -22,
	17, -22, 3, -22, 19, 49, 21, 50, 32, 4, -23, 18, -23, 15,
	-23, 14, -23, 22, 40, 27, -23, 26, -23, 25, -23, 24, -23, 23,
	-23, 20, 46, 32, -23, 17, -23, 3, -23, 19, 49, 21, 50, 32,
	4, -24, 18, -24, 15, -24, 14, -24, 22, 40, 27, -24, 26, -24,
	25, -24, 24, -24, 23, -24, 20, 46, 32, -24, 17, -24, 3, -24,
	19, 49, 21, 50, 32, 4, -25, 18, -25, 15, -25, 14, -25, 22,
	40, 27, -25, 26, -25, 25, -25, 24, -25, 23, -25, 20, 46, 32,
	-25, 17, -25, 3, -25, 19, 49, 21, 50, 32, 4, -26, 18, -26,
	15, -26, 14, -26, 22, 40, 27, 41, 26, 42, 25, 43, 24, 44,
	23, 45, 20, 46, 32, -26, 17, -26, 3, -26, 19, 49, 21, 50,
	32, 4, 39, 18, -27, 15, -27, 14, -27, 22, 40, 27, 41, 26,
	42, 25, 43, 24, 44, 23, 45, 20, 46, 32, -27, 17, -27, 3,
	-27, 19, 49, 21, 50, 32, 4, -39, 18, -39, 15, -39, 14, -39,
	22, -39, 27, -39, 26, -39, 25, -39, 24, -39, 23, -39, 20, -39,
	32, -39, 17, -39, 3, -39, 19, -39, 21, -39, 24, 4, 39, 18,
	83, 22, 40, 27, 41, 26, 42, 25, 43, 24, 44, 23, 45, 20,
	46, 3, 48, 19, 49, 21, 50, 6, 31, -58, 17, -58, 30, -58,
	2, 32, -10, 6, 8, 85, 9, -47, 32, -47, 6, 31, -41, 17,
	-41, 30, -41, 26, 4, 39, 15, -54, 14, -54, 22, 40, 27, 41,
	26, 42, 25, 43, 24, 44, 23, 45, 20, 46, 3, 48, 19, 49,
	21, 50, 6, 8, -49, 9, -49, 32, -49, 4, 9, 89, 32, -51,
	6, 8, 85, 9, -46, 32, -46, 2, 17, 95, 2, 32, -50, 2,
	32, -12, 24, 4, 39, 22, 40, 27, 41, 26, 42, 25, 43, 24,
	44, 23, 45, 20, 46, 17, 94, 3, 48, 19, 49, 21, 50, 6,
	8, -48, 9, -48, 32, -48, 2, 18, 98, 2, 18, 99, 6, 8,
	-13, 9, -13, 32, -13, 2, 32, -14,
}

var _goto = []int32{
	100, 121, 121, 121, 122, 121, 135, 121, 121, 121, 121, 121, 121, 121,
	148, 121, 121, 121, 121, 121, 121, 121, 121, 121, 121, 121, 161, 174,
	121, 121, 121, 121, 121, 121, 121, 121, 181, 194, 121, 211, 224, 237,
	250, 263, 276, 289, 302, 315, 334, 347, 360, 373, 121, 121, 121, 121,
	121, 121, 392, 121, 121, 405, 121, 121, 121, 121, 121, 121, 121, 121,
	121, 121, 121, 121, 121, 121, 121, 121, 121, 408, 121, 121, 421, 121,
	121, 428, 121, 441, 446, 121, 121, 121, 121, 121, 449, 468, 121, 121,
	121, 121, 20, 2, 7, 10, 8, 5, 9, 6, 10, 1, 11, 3,
	12, 17, 13, 18, 14, 9, 15, 4, 16, 0, 12, 11, 35, 10,
	30, 14, 31, 12, 32, 15, 33, 13, 34, 12, 11, 29, 10, 30,
	14, 31, 12, 32, 15, 33, 13, 34, 12, 10, 8, 5, 9, 6,
	10, 3, 38, 9, 15, 4, 16, 12, 11, 56, 10, 30, 14, 31,
	12, 32, 15, 33, 13, 34, 6, 16, 59, 24, 60, 25, 61, 12,
	11, 52, 10, 30, 14, 31, 12, 32, 15, 33, 13, 34, 16, 23,
	53, 22, 54, 11, 55, 10, 30, 14, 31, 12, 32, 15, 33, 13,
	34, 12, 11, 75, 10, 30, 14, 31, 12, 32, 15, 33, 13, 34,
	12, 11, 67, 10, 30, 14, 31, 12, 32, 15, 33, 13, 34, 12,
	11, 74, 10, 30, 14, 31, 12, 32, 15, 33, 13, 34, 12, 11,
	73, 10, 30, 14, 31, 12, 32, 15, 33, 13, 34, 12, 11, 72,
	10, 30, 14, 31, 12, 32, 15, 33, 13, 34, 12, 11, 71, 10,
	30, 14, 31, 12, 32, 15, 33, 13, 34, 12, 11, 70, 10, 30,
	14, 31, 12, 32, 15, 33, 13, 34, 12, 11, 69, 10, 30, 14,
	31, 12, 32, 15, 33, 13, 34, 18, 2, 62, 10, 8, 5, 9,
	6, 10, 3, 12, 17, 13, 18, 14, 9, 15, 4, 16, 12, 11,
	76, 10, 30, 14, 31, 12, 32, 15, 33, 13, 34, 12, 11, 68,
	10, 30, 14, 31, 12, 32, 15, 33, 13, 34, 12, 11, 66, 10,
	30, 14, 31, 12, 32, 15, 33, 13, 34, 18, 2, 63, 10, 8,
	5, 9, 6, 10, 3, 12, 17, 13, 18, 14, 9, 15, 4, 16,
	12, 11, 78, 10, 30, 14, 31, 12, 32, 15, 33, 13, 34, 2,
	16, 80, 12, 11, 84, 10, 30, 14, 31, 12, 32, 15, 33, 13,
	34, 6, 7, 86, 19, 87, 20, 88, 12, 11, 92, 10, 30, 14,
	31, 12, 32, 15, 33, 13, 34, 4, 8, 90, 21, 91, 2, 7,
	93, 18, 2, 96, 10, 8, 5, 9, 6, 10, 3, 12, 17, 13,
	18, 14, 9, 15, 4, 16, 18, 2, 97, 10, 8, 5, 9, 6,
	10, 3, 12, 17, 13, 18, 14, 9, 15, 4, 16,
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

type Error struct {
	Token    Token
	Expected []int
}

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

type _item struct {
	State  int32
	Sym    any
	Bounds _Bounds
}

type lox struct {
	_lex   _Lexer
	_stack _Stack[_item]

	_la    int
	_lasym any

	_qla    int
	_qlasym any
}

func (p *parser) parse(lex _Lexer) bool {
	const accept = 2147483647

	p._lex = lex
	p._qla = -1
	p._stack.Push(_item{})

	p._readToken()

	for {
		topState := p._stack.Peek(0).State
		action, ok := _Find(_actions, topState, int32(p._la))
		if !ok {
			if !p._recover() {
				return false
			}
			continue
		}
		if action == accept {
			break
		} else if action >= 0 { // shift
			latok, ok := p._lasym.(Token)
			if !ok {
				latok = p._lasym.(Error).Token
			}
			p._stack.Push(_item{
				State: action,
				Sym:   p._lasym,
				Bounds: _Bounds{
					Begin: latok,
					End:   latok,
				},
			})
			p._readToken()
		} else { // reduce
			prod := -action
			termCount := _termCounts[int(prod)]
			rule := _rules[int(prod)]
			res := p._act(prod)

			// Compute reduction token bounds.
			// Trim leading and trailing empty bounds.
			boundSlice := p._stack.PeekSlice(int(termCount))
			for len(boundSlice) > 0 && boundSlice[0].Bounds.Empty {
				boundSlice = boundSlice[1:]
			}
			for len(boundSlice) > 0 && boundSlice[len(boundSlice)-1].Bounds.Empty {
				boundSlice = boundSlice[:len(boundSlice)-1]
			}
			var bounds _Bounds
			if len(boundSlice) > 0 {
				bounds.Begin = boundSlice[0].Bounds.Begin
				bounds.End = boundSlice[len(boundSlice)-1].Bounds.End
			} else {
				bounds.Empty = true
			}
			if !bounds.Empty {
				p._onBounds(res, bounds.Begin, bounds.End)
			}
			p._stack.Pop(int(termCount))
			topState = p._stack.Peek(0).State
			nextState, _ := _Find(_goto, topState, rule)
			p._stack.Push(_item{
				State:  nextState,
				Sym:    res,
				Bounds: bounds,
			})
		}
	}

	return true
}

// recoverLookahead can be called during an error production action (an action
// for a production that has a @error term) to recover the lookahead that was
// possibly lost in the process of reducing the error production.
func (p *parser) recoverLookahead(typ int, tok Token) {
	if p._qla != -1 {
		panic("recovered lookahead already pending")
	}

	p._qla = p._la
	p._qlasym = p._lasym
	p._la = typ
	p._lasym = tok
}

func (p *parser) _readToken() {
	if p._qla != -1 {
		p._la = p._qla
		p._lasym = p._qlasym
		p._qla = -1
		p._qlasym = nil
		return
	}

	p._lasym, p._la = p._lex.ReadToken()
	if p._la == ERROR {
		p._lasym = p._makeError()
	}
}

func (p *parser) _recover() bool {
	errSym, ok := p._lasym.(Error)
	if !ok {
		errSym = p._makeError()
	}

	for p._la == ERROR {
		p._readToken()
	}

	for {
		save := p._stack

		for len(p._stack) >= 1 {
			state := p._stack.Peek(0).State

			for {
				action, ok := _Find(_actions, state, int32(ERROR))
				if !ok {
					break
				}

				if action < 0 {
					prod := -action
					rule := _rules[int(prod)]
					state, _ = _Find(_goto, state, rule)
					continue
				}

				state = action

				_, ok = _Find(_actions, state, int32(p._la))
				if !ok {
					break
				}

				p._qla = p._la
				p._qlasym = p._lasym
				p._la = ERROR
				p._lasym = errSym
				return true
			}

			p._stack.Pop(1)
		}

		if p._la == EOF {
			return false
		}

		p._stack = save
		p._readToken()
	}
}

func (p *parser) _makeError() Error {
	e := Error{
		Token: p._lasym.(Token),
	}

	// Compile list of allowed tokens at this state.
	// See _Find for the format of the _actions table.
	s := p._stack.Peek(0).State
	i := int(_actions[int(s)])
	count := int(_actions[i])
	i++
	end := i + count
	for ; i < end; i += 2 {
		e.Expected = append(e.Expected, int(_actions[i]))
	}

	return e
}

func (p *parser) _act(prod int32) any {
	switch prod {
	case 1:
		return p.on_program(
			_cast[*Block](p._stack.Peek(0).Sym),
		)
	case 2:
		return p.on_program__error(
			_cast[Error](p._stack.Peek(0).Sym),
		)
	case 3:
		return p.on_block(
			_cast[[]Statement](p._stack.Peek(0).Sym),
		)
	case 4:
		return p.on_stmt(
			_cast[Statement](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 5:
		return p.on_stmt(
			_cast[Statement](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 6:
		return p.on_stmt(
			_cast[Statement](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 7:
		return p.on_stmt(
			_cast[Statement](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 8:
		return p.on_stmt__kw(
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 9:
		return p.on_stmt__nl(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 10:
		return p.on_while_stmt(
			_cast[_i0.Token](p._stack.Peek(4).Sym),
			_cast[Expr](p._stack.Peek(3).Sym),
			_cast[_i0.Token](p._stack.Peek(2).Sym),
			_cast[*Block](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 11:
		return p.on_func_call_stmt(
			_cast[*FuncCall](p._stack.Peek(0).Sym),
		)
	case 12:
		return p.on_if_stmt(
			_cast[_i0.Token](p._stack.Peek(6).Sym),
			_cast[Expr](p._stack.Peek(5).Sym),
			_cast[_i0.Token](p._stack.Peek(4).Sym),
			_cast[*Block](p._stack.Peek(3).Sym),
			_cast[_i0.Token](p._stack.Peek(2).Sym),
			_cast[[]*Elif](p._stack.Peek(1).Sym),
			_cast[*Else](p._stack.Peek(0).Sym),
		)
	case 13:
		return p.on_elif(
			_cast[_i0.Token](p._stack.Peek(4).Sym),
			_cast[Expr](p._stack.Peek(3).Sym),
			_cast[_i0.Token](p._stack.Peek(2).Sym),
			_cast[*Block](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 14:
		return p.on_else(
			_cast[_i0.Token](p._stack.Peek(3).Sym),
			_cast[_i0.Token](p._stack.Peek(2).Sym),
			_cast[*Block](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 15:
		return p.on_var_assign(
			_cast[_i0.Token](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 16:
		return p.on_func_call(
			_cast[_i0.Token](p._stack.Peek(3).Sym),
			_cast[_i0.Token](p._stack.Peek(2).Sym),
			_cast[[]Expr](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 17:
		return p.on_expr__bin(
			_cast[Expr](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 18:
		return p.on_expr__bin(
			_cast[Expr](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 19:
		return p.on_expr__bin(
			_cast[Expr](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 20:
		return p.on_expr__bin(
			_cast[Expr](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 21:
		return p.on_expr__bin(
			_cast[Expr](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 22:
		return p.on_expr__bin(
			_cast[Expr](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 23:
		return p.on_expr__bin(
			_cast[Expr](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 24:
		return p.on_expr__bin(
			_cast[Expr](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 25:
		return p.on_expr__bin(
			_cast[Expr](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 26:
		return p.on_expr__bin(
			_cast[Expr](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 27:
		return p.on_expr__bin(
			_cast[Expr](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 28:
		return p.on_expr__paren(
			_cast[_i0.Token](p._stack.Peek(2).Sym),
			_cast[Expr](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 29:
		return p.on_expr__simple(
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 30:
		return p.on_simple_expr(
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 31:
		return p.on_simple_expr(
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 32:
		return p.on_simple_expr(
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 33:
		return p.on_var_ref(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 34:
		return p.on_literal__tok(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 35:
		return p.on_literal__tok(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 36:
		return p.on_literal__tok(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 37:
		return p.on_literal__tok(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 38:
		return p.on_literal__string(
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 39:
		return p.on_string(
			_cast[_i0.Token](p._stack.Peek(2).Sym),
			_cast[[]Expr](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 40:
		return p.on_string_part__char_seq(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 41:
		return p.on_string_part__expr(
			_cast[_i0.Token](p._stack.Peek(2).Sym),
			_cast[Expr](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 42: // ZeroOrMore
		return _cast[[]Statement](p._stack.Peek(0).Sym)
	case 43: // ZeroOrMore
		{
			var zero []Statement
			return zero
		}
	case 44:
		{ // OneOrMoreF
			l := _cast[[]Statement](p._stack.Peek(1).Sym)
			e := _cast[Statement](p._stack.Peek(0).Sym)
			if !e.Discard() {
				l = append(l, e)
			}
			return l
		}
	case 45:
		{ // OneOrMoreF
			var l []Statement
			e := _cast[Statement](p._stack.Peek(0).Sym)
			if !e.Discard() {
				l = append(l, e)
			}
			return l
		}
	case 46: // ZeroOrMore
		return _cast[[]*Elif](p._stack.Peek(0).Sym)
	case 47: // ZeroOrMore
		{
			var zero []*Elif
			return zero
		}
	case 48: // OneOrMore
		return append(
			_cast[[]*Elif](p._stack.Peek(1).Sym),
			_cast[*Elif](p._stack.Peek(0).Sym),
		)
	case 49: // OneOrMore
		return []*Elif{
			_cast[*Elif](p._stack.Peek(0).Sym),
		}
	case 50: // ZeroOrOne
		return _cast[*Else](p._stack.Peek(0).Sym)
	case 51: // ZeroOrOne
		{
			var zero *Else
			return zero
		}
	case 52: // ZeroOrOne
		return _cast[[]Expr](p._stack.Peek(0).Sym)
	case 53: // ZeroOrOne
		{
			var zero []Expr
			return zero
		}
	case 54: // List
		return append(
			_cast[[]Expr](p._stack.Peek(2).Sym),
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 55: // List
		return []Expr{
			_cast[Expr](p._stack.Peek(0).Sym),
		}
	case 56: // ZeroOrMore
		return _cast[[]Expr](p._stack.Peek(0).Sym)
	case 57: // ZeroOrMore
		{
			var zero []Expr
			return zero
		}
	case 58: // OneOrMore
		return append(
			_cast[[]Expr](p._stack.Peek(1).Sym),
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 59: // OneOrMore
		return []Expr{
			_cast[Expr](p._stack.Peek(0).Sym),
		}
	default:
		panic("unreachable")
	}
}
