package main

import (
	_i0 "github.com/dcaiafa/loxlex/simplelexer"
)

var _rules = []int32{
	0, 1, 2, 2, 3, 3, 3, 3, 4, 5, 6, 7, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 9, 9, 9,
	10, 11, 11, 12, 12, 13, 13,
}

var _termCounts = []int32{
	1, 1, 3, 1, 1, 1, 1, 0, 5, 1, 3, 4, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 1, 1, 1,
	1, 1, 1, 1, 0, 3, 1,
}

var _actions = []int32{
	57, 66, 71, 80, 87, 94, 97, 104, 109, 116, 123, 134, 171, 71,
	206, 241, 266, 301, 336, 371, 71, 406, 417, 71, 71, 71, 71, 71,
	71, 71, 71, 424, 71, 71, 71, 433, 462, 467, 470, 497, 522, 527,
	562, 597, 632, 667, 702, 737, 772, 807, 842, 877, 912, 947, 71, 982,
	989, 8, 0, -7, 22, 1, 23, -7, 19, 2, 4, 5, 20, 2,
	21, 8, 22, 11, 17, 12, 2, 13, 18, 14, 6, 7, -9, 0,
	-9, 23, -9, 6, 7, -4, 0, -4, 23, -4, 2, 0, 2147483647, 6,
	7, -3, 0, -3, 23, -3, 4, 0, -1, 23, 10, 6, 7, -5,
	0, -5, 23, -5, 6, 7, -6, 0, -6, 23, -6, 10, 7, -7,
	0, -7, 22, 1, 23, -7, 19, 2, 36, 21, -28, 7, -28, 4,
	-28, 3, -28, 11, -28, 0, -28, 16, -28, 15, -28, 14, -28, 13,
	-28, 12, -28, 9, -28, 23, -28, 6, -28, 2, 21, 20, -28, 8,
	-28, 10, -28, 34, 21, -29, 7, -29, 4, -29, 3, -29, 11, -29,
	0, -29, 16, -29, 15, -29, 14, -29, 13, -29, 12, -29, 9, -29,
	23, -29, 6, -29, 20, -29, 8, -29, 10, -29, 34, 21, -30, 7,
	-30, 4, -30, 3, -30, 11, -30, 0, -30, 16, -30, 15, -30, 14,
	-30, 13, -30, 12, -30, 9, -30, 23, -30, 6, -30, 20, -30, 8,
	-30, 10, -30, 24, 21, 23, 11, 24, 16, 25, 15, 26, 14, 27,
	13, 28, 12, 29, 9, 30, 6, 31, 20, 32, 8, 33, 10, 34,
	34, 21, -25, 7, -25, 4, -25, 3, -25, 11, -25, 0, -25, 16,
	-25, 15, -25, 14, -25, 13, -25, 12, -25, 9, -25, 23, -25, 6,
	-25, 20, -25, 8, -25, 10, -25, 34, 21, -26, 7, -26, 4, -26,
	3, -26, 11, -26, 0, -26, 16, -26, 15, -26, 14, -26, 13, -26,
	12, -26, 9, -26, 23, -26, 6, -26, 20, -26, 8, -26, 10, -26,
	34, 21, -24, 7, -24, 4, -24, 3, -24, 11, -24, 0, -24, 16,
	-24, 15, -24, 14, -24, 13, -24, 12, -24, 9, -24, 23, -24, 6,
	-24, 20, -24, 8, -24, 10, -24, 34, 21, -27, 7, -27, 4, -27,
	3, -27, 11, -27, 0, -27, 16, -27, 15, -27, 14, -27, 13, -27,
	12, -27, 9, -27, 23, -27, 6, -27, 20, -27, 8, -27, 10, -27,
	10, 3, -32, 22, 11, 17, 12, 2, 13, 18, 14, 6, 7, -2,
	0, -2, 23, -2, 8, 7, -7, 22, 1, 23, -7, 19, 2, 28,
	21, 23, 7, -10, 11, 24, 0, -10, 16, 25, 15, 26, 14, 27,
	13, 28, 12, 29, 9, 30, 23, -10, 20, 32, 8, 33, 10, 34,
	4, 4, 54, 3, -31, 2, 3, 41, 26, 21, 23, 4, -34, 3,
	-34, 11, 24, 16, 25, 15, 26, 14, 27, 13, 28, 12, 29, 9,
	30, 20, 32, 8, 33, 10, 34, 24, 21, 23, 3, 42, 11, 24,
	16, 25, 15, 26, 14, 27, 13, 28, 12, 29, 9, 30, 20, 32,
	8, 33, 10, 34, 4, 7, 55, 23, 10, 34, 21, -11, 7, -11,
	4, -11, 3, -11, 11, -11, 0, -11, 16, -11, 15, -11, 14, -11,
	13, -11, 12, -11, 9, -11, 23, -11, 6, -11, 20, -11, 8, -11,
	10, -11, 34, 21, -23, 7, -23, 4, -23, 3, -23, 11, -23, 0,
	-23, 16, -23, 15, -23, 14, -23, 13, -23, 12, -23, 9, -23, 23,
	-23, 6, -23, 20, -23, 8, -23, 10, -23, 34, 21, 23, 7, -12,
	4, -12, 3, -12, 11, 24, 0, -12, 16, 25, 15, 26, 14, 27,
	13, 28, 12, 29, 9, -12, 23, -12, 6, -12, 20, 32, 8, -12,
	10, 34, 34, 21, 23, 7, -13, 4, -13, 3, -13, 11, 24, 0,
	-13, 16, 25, 15, 26, 14, 27, 13, 28, 12, 29, 9, -13, 23,
	-13, 6, -13, 20, 32, 8, -13, 10, 34, 34, 21, 23, 7, -14,
	4, -14, 3, -14, 11, -14, 0, -14, 16, 25, 15, 26, 14, 27,
	13, 28, 12, 29, 9, -14, 23, -14, 6, -14, 20, 32, 8, -14,
	10, -14, 34, 21, 23, 7, -15, 4, -15, 3, -15, 11, -15, 0,
	-15, 16, 25, 15, 26, 14, 27, 13, 28, 12, 29, 9, -15, 23,
	-15, 6, -15, 20, 32, 8, -15, 10, -15, 34, 21, 23, 7, -16,
	4, -16, 3, -16, 11, -16, 0, -16, 16, -16, 15, -16, 14, -16,
	13, -16, 12, -16, 9, -16, 23, -16, 6, -16, 20, 32, 8, -16,
	10, -16, 34, 21, 23, 7, -17, 4, -17, 3, -17, 11, -17, 0,
	-17, 16, -17, 15, -17, 14, -17, 13, -17, 12, -17, 9, -17, 23,
	-17, 6, -17, 20, 32, 8, -17, 10, -17, 34, 21, 23, 7, -18,
	4, -18, 3, -18, 11, -18, 0, -18, 16, -18, 15, -18, 14, -18,
	13, -18, 12, -18, 9, -18, 23, -18, 6, -18, 20, 32, 8, -18,
	10, -18, 34, 21, 23, 7, -19, 4, -19, 3, -19, 11, -19, 0,
	-19, 16, -19, 15, -19, 14, -19, 13, -19, 12, -19, 9, -19, 23,
	-19, 6, -19, 20, 32, 8, -19, 10, -19, 34, 21, 23, 7, -20,
	4, -20, 3, -20, 11, -20, 0, -20, 16, -20, 15, -20, 14, -20,
	13, -20, 12, -20, 9, -20, 23, -20, 6, -20, 20, 32, 8, -20,
	10, -20, 34, 21, -21, 7, -21, 4, -21, 3, -21, 11, -21, 0,
	-21, 16, -21, 15, -21, 14, -21, 13, -21, 12, -21, 9, -21, 23,
	-21, 6, -21, 20, 32, 8, -21, 10, -21, 34, 21, -22, 7, -22,
	4, -22, 3, -22, 11, -22, 0, -22, 16, -22, 15, -22, 14, -22,
	13, -22, 12, -22, 9, -22, 23, -22, 6, -22, 20, -22, 8, -22,
	10, -22, 6, 7, -8, 0, -8, 23, -8, 26, 21, 23, 4, -33,
	3, -33, 11, 24, 16, 25, 15, 26, 14, 27, 13, 28, 12, 29,
	9, 30, 20, 32, 8, 33, 10, 34,
}

var _goto = []int32{
	57, 72, 73, 72, 72, 72, 72, 72, 72, 72, 84, 72, 72, 95,
	72, 72, 72, 72, 72, 72, 106, 117, 72, 132, 143, 154, 165, 176,
	187, 198, 209, 220, 233, 244, 255, 72, 72, 72, 72, 72, 72, 72,
	72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 266, 72,
	72, 14, 7, 3, 5, 4, 1, 5, 3, 6, 2, 7, 6, 8,
	4, 9, 0, 10, 8, 15, 7, 16, 11, 17, 9, 18, 10, 19,
	10, 7, 3, 5, 4, 3, 22, 6, 8, 4, 9, 10, 8, 39,
	7, 16, 11, 17, 9, 18, 10, 19, 10, 8, 35, 7, 16, 11,
	17, 9, 18, 10, 19, 14, 13, 36, 12, 37, 8, 38, 7, 16,
	11, 17, 9, 18, 10, 19, 10, 8, 52, 7, 16, 11, 17, 9,
	18, 10, 19, 10, 8, 46, 7, 16, 11, 17, 9, 18, 10, 19,
	10, 8, 51, 7, 16, 11, 17, 9, 18, 10, 19, 10, 8, 50,
	7, 16, 11, 17, 9, 18, 10, 19, 10, 8, 49, 7, 16, 11,
	17, 9, 18, 10, 19, 10, 8, 48, 7, 16, 11, 17, 9, 18,
	10, 19, 10, 8, 47, 7, 16, 11, 17, 9, 18, 10, 19, 10,
	8, 44, 7, 16, 11, 17, 9, 18, 10, 19, 12, 7, 3, 5,
	4, 3, 6, 2, 40, 6, 8, 4, 9, 10, 8, 53, 7, 16,
	11, 17, 9, 18, 10, 19, 10, 8, 43, 7, 16, 11, 17, 9,
	18, 10, 19, 10, 8, 45, 7, 16, 11, 17, 9, 18, 10, 19,
	10, 8, 56, 7, 16, 11, 17, 9, 18, 10, 19,
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
		return p.on_stmts__1(
			_cast[*Block](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[Statement](p._stack.Peek(0).Sym),
		)
	case 3:
		return p.on_stmts__2(
			_cast[Statement](p._stack.Peek(0).Sym),
		)
	case 4:
		return p.on_stmt(
			_cast[Statement](p._stack.Peek(0).Sym),
		)
	case 5:
		return p.on_stmt(
			_cast[Statement](p._stack.Peek(0).Sym),
		)
	case 6:
		return p.on_stmt(
			_cast[Statement](p._stack.Peek(0).Sym),
		)
	case 7:
		return p.on_stmt__empty()
	case 8:
		return p.on_while_stmt(
			_cast[_i0.Token](p._stack.Peek(4).Sym),
			_cast[Expr](p._stack.Peek(3).Sym),
			_cast[_i0.Token](p._stack.Peek(2).Sym),
			_cast[*Block](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 9:
		return p.on_func_call_stmt(
			_cast[*FuncCall](p._stack.Peek(0).Sym),
		)
	case 10:
		return p.on_var_assign(
			_cast[_i0.Token](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 11:
		return p.on_func_call(
			_cast[_i0.Token](p._stack.Peek(3).Sym),
			_cast[_i0.Token](p._stack.Peek(2).Sym),
			_cast[[]Expr](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 12:
		return p.on_expr__bin(
			_cast[Expr](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 13:
		return p.on_expr__bin(
			_cast[Expr](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 14:
		return p.on_expr__bin(
			_cast[Expr](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 15:
		return p.on_expr__bin(
			_cast[Expr](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 16:
		return p.on_expr__bin(
			_cast[Expr](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[Expr](p._stack.Peek(0).Sym),
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
		return p.on_expr__paren(
			_cast[_i0.Token](p._stack.Peek(2).Sym),
			_cast[Expr](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 24:
		return p.on_expr__simple(
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 25:
		return p.on_simple_expr(
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 26:
		return p.on_simple_expr(
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 27:
		return p.on_simple_expr(
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 28:
		return p.on_var_ref(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 29:
		return p.on_literal(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 30:
		return p.on_literal(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 31: // ZeroOrOne
		return _cast[[]Expr](p._stack.Peek(0).Sym)
	case 32: // ZeroOrOne
		{
			var zero []Expr
			return zero
		}
	case 33: // List
		return append(
			_cast[[]Expr](p._stack.Peek(2).Sym),
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 34: // List
		return []Expr{
			_cast[Expr](p._stack.Peek(0).Sym),
		}
	default:
		panic("unreachable")
	}
}
