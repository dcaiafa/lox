package main

import (
	_i0 "github.com/dcaiafa/loxlex/simplelexer"
)

var _rules = []int32{
	0, 1, 1, 2, 3, 3, 3, 3, 3, 3, 4, 5, 6, 7,
	8, 9, 10, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 12, 12, 12, 13, 14, 14, 14, 14, 14, 15, 15, 16,
	16, 17, 17, 18, 18, 19, 19, 20, 20, 21, 21,
}

var _termCounts = []int32{
	1, 1, 1, 1, 2, 2, 2, 2, 2, 1, 5, 1, 7, 5,
	4, 3, 4, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 2,
	1, 1, 0, 2, 1, 1, 0, 1, 0, 3, 1,
}

var _actions = []int32{
	90, 105, 108, 111, 116, 131, 116, 146, 149, 152, 155, 158, 161, 176,
	181, 196, 199, 202, 217, 232, 247, 262, 277, 308, 341, 372, 116, 403,
	434, 465, 490, 521, 552, 583, 614, 116, 639, 656, 116, 116, 116, 116,
	116, 116, 116, 116, 671, 116, 116, 116, 671, 684, 709, 714, 717, 744,
	769, 772, 775, 806, 837, 868, 899, 930, 961, 992, 1023, 1054, 1085, 1116,
	1147, 116, 1178, 1181, 1188, 116, 1215, 1222, 1227, 1234, 1237, 1240, 1243, 1268,
	671, 671, 1275, 1278, 1281, 1288, 14, 11, 1, 0, -40, 1, 2, 12,
	3, 7, 4, 30, 5, 2, 6, 2, 30, 21, 2, 0, -2, 4,
	16, 35, 13, 36, 14, 6, 22, 12, 23, 28, 24, 10, 25, 13,
	26, 29, 27, 5, 28, 14, 18, -9, 11, -9, 0, -9, 12, -9,
	7, -9, 30, -9, 2, -9, 2, 0, -1, 2, 30, -11, 2, 30,
	17, 2, 30, 20, 2, 0, 2147483647, 14, 18, -42, 11, -42, 0, -42,
	12, -42, 7, -42, 30, -42, 2, -42, 4, 18, -3, 0, -3, 14,
	18, -39, 11, 1, 0, -39, 12, 3, 7, 4, 30, 5, 2, 6,
	2, 30, 18, 2, 30, 19, 14, 18, -4, 11, -4, 0, -4, 12,
	-4, 7, -4, 30, -4, 2, -4, 14, 18, -5, 11, -5, 0, -5,
	12, -5, 7, -5, 30, -5, 2, -5, 14, 18, -6, 11, -6, 0,
	-6, 12, -6, 7, -6, 30, -6, 2, -6, 14, 18, -7, 11, -7,
	0, -7, 12, -7, 7, -7, 30, -7, 2, -7, 14, 18, -8, 11,
	-8, 0, -8, 12, -8, 7, -8, 30, -8, 2, -8, 30, 4, -37,
	15, -37, 14, -37, 22, -37, 27, -37, 26, -37, 25, -37, 24, -37,
	23, -37, 20, -37, 30, -37, 17, -37, 3, -37, 19, -37, 21, -37,
	32, 4, -33, 15, -33, 14, -33, 22, -33, 27, -33, 26, -33, 25,
	-33, 24, -33, 23, -33, 20, -33, 30, -33, 17, -33, 13, 36, 3,
	-33, 19, -33, 21, -33, 30, 4, -34, 15, -34, 14, -34, 22, -34,
	27, -34, 26, -34, 25, -34, 24, -34, 23, -34, 20, -34, 30, -34,
	17, -34, 3, -34, 19, -34, 21, -34, 30, 4, -38, 15, -38, 14,
	-38, 22, -38, 27, -38, 26, -38, 25, -38, 24, -38, 23, -38, 20,
	-38, 30, -38, 17, -38, 3, -38, 19, -38, 21, -38, 30, 4, -35,
	15, -35, 14, -35, 22, -35, 27, -35, 26, -35, 25, -35, 24, -35,
	23, -35, 20, -35, 30, -35, 17, -35, 3, -35, 19, -35, 21, -35,
	30, 4, -36, 15, -36, 14, -36, 22, -36, 27, -36, 26, -36, 25,
	-36, 24, -36, 23, -36, 20, -36, 30, -36, 17, -36, 3, -36, 19,
	-36, 21, -36, 24, 4, 38, 22, 39, 27, 40, 26, 41, 25, 42,
	24, 43, 23, 44, 20, 45, 17, 46, 3, 47, 19, 48, 21, 49,
	30, 4, -30, 15, -30, 14, -30, 22, -30, 27, -30, 26, -30, 25,
	-30, 24, -30, 23, -30, 20, -30, 30, -30, 17, -30, 3, -30, 19,
	-30, 21, -30, 30, 4, -31, 15, -31, 14, -31, 22, -31, 27, -31,
	26, -31, 25, -31, 24, -31, 23, -31, 20, -31, 30, -31, 17, -31,
	3, -31, 19, -31, 21, -31, 30, 4, -29, 15, -29, 14, -29, 22,
	-29, 27, -29, 26, -29, 25, -29, 24, -29, 23, -29, 20, -29, 30,
	-29, 17, -29, 3, -29, 19, -29, 21, -29, 30, 4, -32, 15, -32,
	14, -32, 22, -32, 27, -32, 26, -32, 25, -32, 24, -32, 23, -32,
	20, -32, 30, -32, 17, -32, 3, -32, 19, -32, 21, -32, 24, 4,
	38, 22, 39, 27, 40, 26, 41, 25, 42, 24, 43, 23, 44, 20,
	45, 17, 50, 3, 47, 19, 48, 21, 49, 16, 14, -50, 6, 22,
	12, 23, 28, 24, 10, 25, 13, 26, 29, 27, 5, 28, 14, 18,
	-41, 11, -41, 0, -41, 12, -41, 7, -41, 30, -41, 2, -41, 12,
	18, -40, 11, 1, 12, 3, 7, 4, 30, 5, 2, 6, 24, 4,
	38, 22, 39, 27, 40, 26, 41, 25, 42, 24, 43, 23, 44, 20,
	45, 30, -15, 3, 47, 19, 48, 21, 49, 4, 15, 71, 14, -49,
	2, 14, 58, 26, 4, 38, 15, -52, 14, -52, 22, 39, 27, 40,
	26, 41, 25, 42, 24, 43, 23, 44, 20, 45, 3, 47, 19, 48,
	21, 49, 24, 4, 38, 14, 59, 22, 39, 27, 40, 26, 41, 25,
	42, 24, 43, 23, 44, 20, 45, 3, 47, 19, 48, 21, 49, 2,
	18, 72, 2, 18, 73, 30, 4, -16, 15, -16, 14, -16, 22, -16,
	27, -16, 26, -16, 25, -16, 24, -16, 23, -16, 20, -16, 30, -16,
	17, -16, 3, -16, 19, -16, 21, -16, 30, 4, -28, 15, -28, 14,
	-28, 22, -28, 27, -28, 26, -28, 25, -28, 24, -28, 23, -28, 20,
	-28, 30, -28, 17, -28, 3, -28, 19, -28, 21, -28, 30, 4, -17,
	15, -17, 14, -17, 22, -17, 27, -17, 26, -17, 25, -17, 24, -17,
	23, -17, 20, -17, 30, -17, 17, -17, 3, -17, 19, -17, 21, -17,
	30, 4, -18, 15, -18, 14, -18, 22, -18, 27, -18, 26, -18, 25,
	-18, 24, -18, 23, -18, 20, -18, 30, -18, 17, -18, 3, -18, 19,
	-18, 21, -18, 30, 4, -19, 15, -19, 14, -19, 22, 39, 27, -19,
	26, -19, 25, -19, 24, -19, 23, -19, 20, -19, 30, -19, 17, -19,
	3, -19, 19, -19, 21, 49, 30, 4, -20, 15, -20, 14, -20, 22,
	39, 27, -20, 26, -20, 25, -20, 24, -20, 23, -20, 20, -20, 30,
	-20, 17, -20, 3, -20, 19, -20, 21, 49, 30, 4, -21, 15, -21,
	14, -21, 22, 39, 27, -21, 26, -21, 25, -21, 24, -21, 23, -21,
	20, 45, 30, -21, 17, -21, 3, -21, 19, 48, 21, 49, 30, 4,
	-22, 15, -22, 14, -22, 22, 39, 27, -22, 26, -22, 25, -22, 24,
	-22, 23, -22, 20, 45, 30, -22, 17, -22, 3, -22, 19, 48, 21,
	49, 30, 4, -23, 15, -23, 14, -23, 22, 39, 27, -23, 26, -23,
	25, -23, 24, -23, 23, -23, 20, 45, 30, -23, 17, -23, 3, -23,
	19, 48, 21, 49, 30, 4, -24, 15, -24, 14, -24, 22, 39, 27,
	-24, 26, -24, 25, -24, 24, -24, 23, -24, 20, 45, 30, -24, 17,
	-24, 3, -24, 19, 48, 21, 49, 30, 4, -25, 15, -25, 14, -25,
	22, 39, 27, -25, 26, -25, 25, -25, 24, -25, 23, -25, 20, 45,
	30, -25, 17, -25, 3, -25, 19, 48, 21, 49, 30, 4, -26, 15,
	-26, 14, -26, 22, 39, 27, 40, 26, 41, 25, 42, 24, 43, 23,
	44, 20, 45, 30, -26, 17, -26, 3, -26, 19, 48, 21, 49, 30,
	4, 38, 15, -27, 14, -27, 22, 39, 27, 40, 26, 41, 25, 42,
	24, 43, 23, 44, 20, 45, 30, -27, 17, -27, 3, -27, 19, 48,
	21, 49, 2, 30, -10, 6, 8, 75, 9, -44, 30, -44, 26, 4,
	38, 15, -51, 14, -51, 22, 39, 27, 40, 26, 41, 25, 42, 24,
	43, 23, 44, 20, 45, 3, 47, 19, 48, 21, 49, 6, 8, -46,
	9, -46, 30, -46, 4, 9, 79, 30, -48, 6, 8, 75, 9, -43,
	30, -43, 2, 17, 85, 2, 30, -47, 2, 30, -12, 24, 4, 38,
	22, 39, 27, 40, 26, 41, 25, 42, 24, 43, 23, 44, 20, 45,
	17, 84, 3, 47, 19, 48, 21, 49, 6, 8, -45, 9, -45, 30,
	-45, 2, 18, 88, 2, 18, 89, 6, 8, -13, 9, -13, 30, -13,
	2, 30, -14,
}

var _goto = []int32{
	90, 111, 111, 111, 112, 111, 123, 111, 111, 111, 111, 111, 111, 111,
	134, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 147, 111,
	111, 111, 111, 111, 111, 111, 111, 158, 169, 111, 184, 195, 206, 217,
	228, 239, 250, 261, 272, 291, 302, 313, 324, 111, 111, 111, 111, 111,
	111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111,
	111, 343, 111, 354, 111, 361, 111, 372, 377, 111, 111, 111, 111, 111,
	380, 399, 111, 111, 111, 111, 20, 2, 7, 10, 8, 5, 9, 6,
	10, 1, 11, 3, 12, 15, 13, 16, 14, 9, 15, 4, 16, 0,
	10, 11, 34, 10, 30, 14, 31, 12, 32, 13, 33, 10, 11, 29,
	10, 30, 14, 31, 12, 32, 13, 33, 12, 10, 8, 5, 9, 6,
	10, 3, 37, 9, 15, 4, 16, 10, 11, 55, 10, 30, 14, 31,
	12, 32, 13, 33, 10, 11, 51, 10, 30, 14, 31, 12, 32, 13,
	33, 14, 21, 52, 20, 53, 11, 54, 10, 30, 14, 31, 12, 32,
	13, 33, 10, 11, 69, 10, 30, 14, 31, 12, 32, 13, 33, 10,
	11, 61, 10, 30, 14, 31, 12, 32, 13, 33, 10, 11, 68, 10,
	30, 14, 31, 12, 32, 13, 33, 10, 11, 67, 10, 30, 14, 31,
	12, 32, 13, 33, 10, 11, 66, 10, 30, 14, 31, 12, 32, 13,
	33, 10, 11, 65, 10, 30, 14, 31, 12, 32, 13, 33, 10, 11,
	64, 10, 30, 14, 31, 12, 32, 13, 33, 10, 11, 63, 10, 30,
	14, 31, 12, 32, 13, 33, 18, 2, 56, 10, 8, 5, 9, 6,
	10, 3, 12, 15, 13, 16, 14, 9, 15, 4, 16, 10, 11, 70,
	10, 30, 14, 31, 12, 32, 13, 33, 10, 11, 62, 10, 30, 14,
	31, 12, 32, 13, 33, 10, 11, 60, 10, 30, 14, 31, 12, 32,
	13, 33, 18, 2, 57, 10, 8, 5, 9, 6, 10, 3, 12, 15,
	13, 16, 14, 9, 15, 4, 16, 10, 11, 74, 10, 30, 14, 31,
	12, 32, 13, 33, 6, 7, 76, 17, 77, 18, 78, 10, 11, 82,
	10, 30, 14, 31, 12, 32, 13, 33, 4, 8, 80, 19, 81, 2,
	7, 83, 18, 2, 86, 10, 8, 5, 9, 6, 10, 3, 12, 15,
	13, 16, 14, 9, 15, 4, 16, 18, 2, 87, 10, 8, 5, 9,
	6, 10, 3, 12, 15, 13, 16, 14, 9, 15, 4, 16,
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
	State int32
	Sym   any
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
			p._stack.Push(_item{
				State: action,
				Sym:   p._lasym,
			})
			p._readToken()
		} else { // reduce
			prod := -action
			termCount := _termCounts[int(prod)]
			rule := _rules[int(prod)]
			res := p._act(prod)
			p._stack.Pop(int(termCount))
			topState = p._stack.Peek(0).State
			nextState, _ := _Find(_goto, topState, rule)
			p._stack.Push(_item{
				State: nextState,
				Sym:   res,
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
		return p.on_literal(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 35:
		return p.on_literal(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 36:
		return p.on_literal(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 37:
		return p.on_literal(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 38:
		return p.on_literal(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 39: // ZeroOrMore
		return _cast[[]Statement](p._stack.Peek(0).Sym)
	case 40: // ZeroOrMore
		{
			var zero []Statement
			return zero
		}
	case 41:
		{ // OneOrMoreF
			l := _cast[[]Statement](p._stack.Peek(1).Sym)
			e := _cast[Statement](p._stack.Peek(0).Sym)
			if !e.Discard() {
				l = append(l, e)
			}
			return l
		}
	case 42:
		{ // OneOrMoreF
			var l []Statement
			e := _cast[Statement](p._stack.Peek(0).Sym)
			if !e.Discard() {
				l = append(l, e)
			}
			return l
		}
	case 43: // ZeroOrMore
		return _cast[[]*Elif](p._stack.Peek(0).Sym)
	case 44: // ZeroOrMore
		{
			var zero []*Elif
			return zero
		}
	case 45: // OneOrMore
		return append(
			_cast[[]*Elif](p._stack.Peek(1).Sym),
			_cast[*Elif](p._stack.Peek(0).Sym),
		)
	case 46: // OneOrMore
		return []*Elif{
			_cast[*Elif](p._stack.Peek(0).Sym),
		}
	case 47: // ZeroOrOne
		return _cast[*Else](p._stack.Peek(0).Sym)
	case 48: // ZeroOrOne
		{
			var zero *Else
			return zero
		}
	case 49: // ZeroOrOne
		return _cast[[]Expr](p._stack.Peek(0).Sym)
	case 50: // ZeroOrOne
		{
			var zero []Expr
			return zero
		}
	case 51: // List
		return append(
			_cast[[]Expr](p._stack.Peek(2).Sym),
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 52: // List
		return []Expr{
			_cast[Expr](p._stack.Peek(0).Sym),
		}
	default:
		panic("unreachable")
	}
}
