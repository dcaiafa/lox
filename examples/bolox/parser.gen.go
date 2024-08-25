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
	671, 671, 1275, 1278, 1281, 1288, 14, 28, 1, 0, -40, 1, 2, 29,
	3, 24, 4, 30, 5, 19, 6, 2, 30, 21, 2, 0, -2, 4,
	5, 35, 2, 36, 14, 23, 22, 29, 23, 17, 24, 27, 25, 2,
	26, 18, 27, 22, 28, 14, 7, -9, 28, -9, 0, -9, 29, -9,
	24, -9, 30, -9, 19, -9, 2, 0, -1, 2, 30, -11, 2, 30,
	17, 2, 30, 20, 2, 0, 2147483647, 14, 7, -42, 28, -42, 0, -42,
	29, -42, 24, -42, 30, -42, 19, -42, 4, 7, -3, 0, -3, 14,
	7, -39, 28, 1, 0, -39, 29, 3, 24, 4, 30, 5, 19, 6,
	2, 30, 18, 2, 30, 19, 14, 7, -4, 28, -4, 0, -4, 29,
	-4, 24, -4, 30, -4, 19, -4, 14, 7, -5, 28, -5, 0, -5,
	29, -5, 24, -5, 30, -5, 19, -5, 14, 7, -6, 28, -6, 0,
	-6, 29, -6, 24, -6, 30, -6, 19, -6, 14, 7, -7, 28, -7,
	0, -7, 29, -7, 24, -7, 30, -7, 19, -7, 14, 7, -8, 28,
	-8, 0, -8, 29, -8, 24, -8, 30, -8, 19, -8, 30, 21, -37,
	4, -37, 3, -37, 11, -37, 16, -37, 15, -37, 14, -37, 13, -37,
	12, -37, 9, -37, 30, -37, 6, -37, 20, -37, 8, -37, 10, -37,
	32, 21, -33, 4, -33, 3, -33, 11, -33, 16, -33, 15, -33, 14,
	-33, 13, -33, 12, -33, 9, -33, 30, -33, 6, -33, 2, 36, 20,
	-33, 8, -33, 10, -33, 30, 21, -34, 4, -34, 3, -34, 11, -34,
	16, -34, 15, -34, 14, -34, 13, -34, 12, -34, 9, -34, 30, -34,
	6, -34, 20, -34, 8, -34, 10, -34, 30, 21, -38, 4, -38, 3,
	-38, 11, -38, 16, -38, 15, -38, 14, -38, 13, -38, 12, -38, 9,
	-38, 30, -38, 6, -38, 20, -38, 8, -38, 10, -38, 30, 21, -35,
	4, -35, 3, -35, 11, -35, 16, -35, 15, -35, 14, -35, 13, -35,
	12, -35, 9, -35, 30, -35, 6, -35, 20, -35, 8, -35, 10, -35,
	30, 21, -36, 4, -36, 3, -36, 11, -36, 16, -36, 15, -36, 14,
	-36, 13, -36, 12, -36, 9, -36, 30, -36, 6, -36, 20, -36, 8,
	-36, 10, -36, 24, 21, 38, 11, 39, 16, 40, 15, 41, 14, 42,
	13, 43, 12, 44, 9, 45, 6, 46, 20, 47, 8, 48, 10, 49,
	30, 21, -30, 4, -30, 3, -30, 11, -30, 16, -30, 15, -30, 14,
	-30, 13, -30, 12, -30, 9, -30, 30, -30, 6, -30, 20, -30, 8,
	-30, 10, -30, 30, 21, -31, 4, -31, 3, -31, 11, -31, 16, -31,
	15, -31, 14, -31, 13, -31, 12, -31, 9, -31, 30, -31, 6, -31,
	20, -31, 8, -31, 10, -31, 30, 21, -29, 4, -29, 3, -29, 11,
	-29, 16, -29, 15, -29, 14, -29, 13, -29, 12, -29, 9, -29, 30,
	-29, 6, -29, 20, -29, 8, -29, 10, -29, 30, 21, -32, 4, -32,
	3, -32, 11, -32, 16, -32, 15, -32, 14, -32, 13, -32, 12, -32,
	9, -32, 30, -32, 6, -32, 20, -32, 8, -32, 10, -32, 24, 21,
	38, 11, 39, 16, 40, 15, 41, 14, 42, 13, 43, 12, 44, 9,
	45, 6, 50, 20, 47, 8, 48, 10, 49, 16, 3, -50, 23, 22,
	29, 23, 17, 24, 27, 25, 2, 26, 18, 27, 22, 28, 14, 7,
	-41, 28, -41, 0, -41, 29, -41, 24, -41, 30, -41, 19, -41, 12,
	7, -40, 28, 1, 29, 3, 24, 4, 30, 5, 19, 6, 24, 21,
	38, 11, 39, 16, 40, 15, 41, 14, 42, 13, 43, 12, 44, 9,
	45, 30, -15, 20, 47, 8, 48, 10, 49, 4, 4, 71, 3, -49,
	2, 3, 58, 26, 21, 38, 4, -52, 3, -52, 11, 39, 16, 40,
	15, 41, 14, 42, 13, 43, 12, 44, 9, 45, 20, 47, 8, 48,
	10, 49, 24, 21, 38, 3, 59, 11, 39, 16, 40, 15, 41, 14,
	42, 13, 43, 12, 44, 9, 45, 20, 47, 8, 48, 10, 49, 2,
	7, 72, 2, 7, 73, 30, 21, -16, 4, -16, 3, -16, 11, -16,
	16, -16, 15, -16, 14, -16, 13, -16, 12, -16, 9, -16, 30, -16,
	6, -16, 20, -16, 8, -16, 10, -16, 30, 21, -28, 4, -28, 3,
	-28, 11, -28, 16, -28, 15, -28, 14, -28, 13, -28, 12, -28, 9,
	-28, 30, -28, 6, -28, 20, -28, 8, -28, 10, -28, 30, 21, -17,
	4, -17, 3, -17, 11, -17, 16, -17, 15, -17, 14, -17, 13, -17,
	12, -17, 9, -17, 30, -17, 6, -17, 20, -17, 8, -17, 10, -17,
	30, 21, -18, 4, -18, 3, -18, 11, -18, 16, -18, 15, -18, 14,
	-18, 13, -18, 12, -18, 9, -18, 30, -18, 6, -18, 20, -18, 8,
	-18, 10, -18, 30, 21, -19, 4, -19, 3, -19, 11, 39, 16, -19,
	15, -19, 14, -19, 13, -19, 12, -19, 9, -19, 30, -19, 6, -19,
	20, -19, 8, -19, 10, 49, 30, 21, -20, 4, -20, 3, -20, 11,
	39, 16, -20, 15, -20, 14, -20, 13, -20, 12, -20, 9, -20, 30,
	-20, 6, -20, 20, -20, 8, -20, 10, 49, 30, 21, -21, 4, -21,
	3, -21, 11, 39, 16, -21, 15, -21, 14, -21, 13, -21, 12, -21,
	9, 45, 30, -21, 6, -21, 20, -21, 8, 48, 10, 49, 30, 21,
	-22, 4, -22, 3, -22, 11, 39, 16, -22, 15, -22, 14, -22, 13,
	-22, 12, -22, 9, 45, 30, -22, 6, -22, 20, -22, 8, 48, 10,
	49, 30, 21, -23, 4, -23, 3, -23, 11, 39, 16, -23, 15, -23,
	14, -23, 13, -23, 12, -23, 9, 45, 30, -23, 6, -23, 20, -23,
	8, 48, 10, 49, 30, 21, -24, 4, -24, 3, -24, 11, 39, 16,
	-24, 15, -24, 14, -24, 13, -24, 12, -24, 9, 45, 30, -24, 6,
	-24, 20, -24, 8, 48, 10, 49, 30, 21, -25, 4, -25, 3, -25,
	11, 39, 16, -25, 15, -25, 14, -25, 13, -25, 12, -25, 9, 45,
	30, -25, 6, -25, 20, -25, 8, 48, 10, 49, 30, 21, -26, 4,
	-26, 3, -26, 11, 39, 16, 40, 15, 41, 14, 42, 13, 43, 12,
	44, 9, 45, 30, -26, 6, -26, 20, -26, 8, 48, 10, 49, 30,
	21, 38, 4, -27, 3, -27, 11, 39, 16, 40, 15, 41, 14, 42,
	13, 43, 12, 44, 9, 45, 30, -27, 6, -27, 20, -27, 8, 48,
	10, 49, 2, 30, -10, 6, 25, 75, 26, -44, 30, -44, 26, 21,
	38, 4, -51, 3, -51, 11, 39, 16, 40, 15, 41, 14, 42, 13,
	43, 12, 44, 9, 45, 20, 47, 8, 48, 10, 49, 6, 25, -46,
	26, -46, 30, -46, 4, 26, 79, 30, -48, 6, 25, 75, 26, -43,
	30, -43, 2, 6, 85, 2, 30, -47, 2, 30, -12, 24, 21, 38,
	11, 39, 16, 40, 15, 41, 14, 42, 13, 43, 12, 44, 9, 45,
	6, 84, 20, 47, 8, 48, 10, 49, 6, 25, -45, 26, -45, 30,
	-45, 2, 7, 88, 2, 7, 89, 6, 25, -13, 26, -13, 30, -13,
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
