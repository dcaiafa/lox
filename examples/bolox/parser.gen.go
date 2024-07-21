package main

import (
	_i0 "github.com/dcaiafa/loxlex/simplelexer"
)

var _rules = []int32{
	0, 1, 1, 2, 3, 3, 3, 3, 4, 5, 6, 7, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 9, 9, 9,
	10, 11, 11, 12, 12, 13, 13, 14, 14, 15, 15,
}

var _termCounts = []int32{
	1, 1, 1, 1, 2, 2, 2, 1, 5, 1, 3, 4, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 1, 1, 1,
	1, 1, 1, 1, 0, 2, 1, 1, 0, 3, 1,
}

var _actions = []int32{
	63, 74, 77, 82, 93, 102, 105, 108, 111, 114, 125, 130, 141, 144,
	147, 158, 169, 180, 213, 93, 244, 275, 300, 331, 362, 393, 93, 424,
	435, 93, 93, 93, 93, 93, 93, 93, 93, 446, 93, 93, 93, 455,
	480, 485, 488, 515, 540, 543, 574, 605, 636, 667, 698, 729, 760, 791,
	822, 853, 884, 915, 93, 946, 949, 10, 0, -32, 1, 1, 22, 2,
	23, 3, 19, 4, 2, 0, -2, 4, 5, 26, 2, 27, 10, 7,
	-7, 0, -7, 22, -7, 23, -7, 19, -7, 8, 22, 17, 17, 18,
	2, 19, 18, 20, 2, 0, -1, 2, 23, -9, 2, 23, 14, 2,
	0, 2147483647, 10, 7, -34, 0, -34, 22, -34, 23, -34, 19, -34, 4,
	7, -3, 0, -3, 10, 7, -31, 0, -31, 22, 2, 23, 3, 19,
	4, 2, 23, 15, 2, 23, 16, 10, 7, -4, 0, -4, 22, -4,
	23, -4, 19, -4, 10, 7, -5, 0, -5, 22, -5, 23, -5, 19,
	-5, 10, 7, -6, 0, -6, 22, -6, 23, -6, 19, -6, 32, 21,
	-28, 4, -28, 3, -28, 11, -28, 16, -28, 15, -28, 14, -28, 13,
	-28, 12, -28, 9, -28, 23, -28, 6, -28, 2, 27, 20, -28, 8,
	-28, 10, -28, 30, 21, -29, 4, -29, 3, -29, 11, -29, 16, -29,
	15, -29, 14, -29, 13, -29, 12, -29, 9, -29, 23, -29, 6, -29,
	20, -29, 8, -29, 10, -29, 30, 21, -30, 4, -30, 3, -30, 11,
	-30, 16, -30, 15, -30, 14, -30, 13, -30, 12, -30, 9, -30, 23,
	-30, 6, -30, 20, -30, 8, -30, 10, -30, 24, 21, 29, 11, 30,
	16, 31, 15, 32, 14, 33, 13, 34, 12, 35, 9, 36, 6, 37,
	20, 38, 8, 39, 10, 40, 30, 21, -25, 4, -25, 3, -25, 11,
	-25, 16, -25, 15, -25, 14, -25, 13, -25, 12, -25, 9, -25, 23,
	-25, 6, -25, 20, -25, 8, -25, 10, -25, 30, 21, -26, 4, -26,
	3, -26, 11, -26, 16, -26, 15, -26, 14, -26, 13, -26, 12, -26,
	9, -26, 23, -26, 6, -26, 20, -26, 8, -26, 10, -26, 30, 21,
	-24, 4, -24, 3, -24, 11, -24, 16, -24, 15, -24, 14, -24, 13,
	-24, 12, -24, 9, -24, 23, -24, 6, -24, 20, -24, 8, -24, 10,
	-24, 30, 21, -27, 4, -27, 3, -27, 11, -27, 16, -27, 15, -27,
	14, -27, 13, -27, 12, -27, 9, -27, 23, -27, 6, -27, 20, -27,
	8, -27, 10, -27, 10, 3, -36, 22, 17, 17, 18, 2, 19, 18,
	20, 10, 7, -33, 0, -33, 22, -33, 23, -33, 19, -33, 8, 7,
	-32, 22, 2, 23, 3, 19, 4, 24, 21, 29, 11, 30, 16, 31,
	15, 32, 14, 33, 13, 34, 12, 35, 9, 36, 23, -10, 20, 38,
	8, 39, 10, 40, 4, 4, 60, 3, -35, 2, 3, 47, 26, 21,
	29, 4, -38, 3, -38, 11, 30, 16, 31, 15, 32, 14, 33, 13,
	34, 12, 35, 9, 36, 20, 38, 8, 39, 10, 40, 24, 21, 29,
	3, 48, 11, 30, 16, 31, 15, 32, 14, 33, 13, 34, 12, 35,
	9, 36, 20, 38, 8, 39, 10, 40, 2, 7, 61, 30, 21, -11,
	4, -11, 3, -11, 11, -11, 16, -11, 15, -11, 14, -11, 13, -11,
	12, -11, 9, -11, 23, -11, 6, -11, 20, -11, 8, -11, 10, -11,
	30, 21, -23, 4, -23, 3, -23, 11, -23, 16, -23, 15, -23, 14,
	-23, 13, -23, 12, -23, 9, -23, 23, -23, 6, -23, 20, -23, 8,
	-23, 10, -23, 30, 21, 29, 4, -12, 3, -12, 11, 30, 16, 31,
	15, 32, 14, 33, 13, 34, 12, 35, 9, -12, 23, -12, 6, -12,
	20, 38, 8, -12, 10, 40, 30, 21, 29, 4, -13, 3, -13, 11,
	30, 16, 31, 15, 32, 14, 33, 13, 34, 12, 35, 9, -13, 23,
	-13, 6, -13, 20, 38, 8, -13, 10, 40, 30, 21, 29, 4, -14,
	3, -14, 11, -14, 16, 31, 15, 32, 14, 33, 13, 34, 12, 35,
	9, -14, 23, -14, 6, -14, 20, 38, 8, -14, 10, -14, 30, 21,
	29, 4, -15, 3, -15, 11, -15, 16, 31, 15, 32, 14, 33, 13,
	34, 12, 35, 9, -15, 23, -15, 6, -15, 20, 38, 8, -15, 10,
	-15, 30, 21, 29, 4, -16, 3, -16, 11, -16, 16, -16, 15, -16,
	14, -16, 13, -16, 12, -16, 9, -16, 23, -16, 6, -16, 20, 38,
	8, -16, 10, -16, 30, 21, 29, 4, -17, 3, -17, 11, -17, 16,
	-17, 15, -17, 14, -17, 13, -17, 12, -17, 9, -17, 23, -17, 6,
	-17, 20, 38, 8, -17, 10, -17, 30, 21, 29, 4, -18, 3, -18,
	11, -18, 16, -18, 15, -18, 14, -18, 13, -18, 12, -18, 9, -18,
	23, -18, 6, -18, 20, 38, 8, -18, 10, -18, 30, 21, 29, 4,
	-19, 3, -19, 11, -19, 16, -19, 15, -19, 14, -19, 13, -19, 12,
	-19, 9, -19, 23, -19, 6, -19, 20, 38, 8, -19, 10, -19, 30,
	21, 29, 4, -20, 3, -20, 11, -20, 16, -20, 15, -20, 14, -20,
	13, -20, 12, -20, 9, -20, 23, -20, 6, -20, 20, 38, 8, -20,
	10, -20, 30, 21, -21, 4, -21, 3, -21, 11, -21, 16, -21, 15,
	-21, 14, -21, 13, -21, 12, -21, 9, -21, 23, -21, 6, -21, 20,
	38, 8, -21, 10, -21, 30, 21, -22, 4, -22, 3, -22, 11, -22,
	16, -22, 15, -22, 14, -22, 13, -22, 12, -22, 9, -22, 23, -22,
	6, -22, 20, -22, 8, -22, 10, -22, 2, 23, -8, 26, 21, 29,
	4, -37, 3, -37, 11, 30, 16, 31, 15, 32, 14, 33, 13, 34,
	12, 35, 9, 36, 20, 38, 8, 39, 10, 40,
}

var _goto = []int32{
	63, 82, 82, 82, 83, 82, 82, 82, 82, 82, 82, 94, 82, 82,
	82, 82, 82, 82, 82, 105, 82, 82, 82, 82, 82, 82, 116, 127,
	82, 142, 153, 164, 175, 186, 197, 208, 219, 230, 247, 258, 269, 82,
	82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82,
	82, 82, 82, 82, 280, 82, 82, 18, 2, 5, 7, 6, 5, 7,
	1, 8, 3, 9, 12, 10, 13, 11, 6, 12, 4, 13, 0, 10,
	8, 21, 7, 22, 11, 23, 9, 24, 10, 25, 10, 7, 6, 5,
	7, 3, 28, 6, 12, 4, 13, 10, 8, 45, 7, 22, 11, 23,
	9, 24, 10, 25, 10, 8, 41, 7, 22, 11, 23, 9, 24, 10,
	25, 14, 15, 42, 14, 43, 8, 44, 7, 22, 11, 23, 9, 24,
	10, 25, 10, 8, 58, 7, 22, 11, 23, 9, 24, 10, 25, 10,
	8, 52, 7, 22, 11, 23, 9, 24, 10, 25, 10, 8, 57, 7,
	22, 11, 23, 9, 24, 10, 25, 10, 8, 56, 7, 22, 11, 23,
	9, 24, 10, 25, 10, 8, 55, 7, 22, 11, 23, 9, 24, 10,
	25, 10, 8, 54, 7, 22, 11, 23, 9, 24, 10, 25, 10, 8,
	53, 7, 22, 11, 23, 9, 24, 10, 25, 10, 8, 50, 7, 22,
	11, 23, 9, 24, 10, 25, 16, 2, 46, 7, 6, 5, 7, 3,
	9, 12, 10, 13, 11, 6, 12, 4, 13, 10, 8, 59, 7, 22,
	11, 23, 9, 24, 10, 25, 10, 8, 49, 7, 22, 11, 23, 9,
	24, 10, 25, 10, 8, 51, 7, 22, 11, 23, 9, 24, 10, 25,
	10, 8, 62, 7, 22, 11, 23, 9, 24, 10, 25,
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
		return p.on_stmt__nl(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
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
	case 31: // ZeroOrMore
		return _cast[[]Statement](p._stack.Peek(0).Sym)
	case 32: // ZeroOrMore
		{
			var zero []Statement
			return zero
		}
	case 33:
		{ // OneOrMoreF
			l := _cast[[]Statement](p._stack.Peek(1).Sym)
			e := _cast[Statement](p._stack.Peek(0).Sym)
			if !e.Discard() {
				l = append(l, e)
			}
			return l
		}
	case 34:
		{ // OneOrMoreF
			var l []Statement
			e := _cast[Statement](p._stack.Peek(0).Sym)
			if !e.Discard() {
				l = append(l, e)
			}
			return l
		}
	case 35: // ZeroOrOne
		return _cast[[]Expr](p._stack.Peek(0).Sym)
	case 36: // ZeroOrOne
		{
			var zero []Expr
			return zero
		}
	case 37: // List
		return append(
			_cast[[]Expr](p._stack.Peek(2).Sym),
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 38: // List
		return []Expr{
			_cast[Expr](p._stack.Peek(0).Sym),
		}
	default:
		panic("unreachable")
	}
}
