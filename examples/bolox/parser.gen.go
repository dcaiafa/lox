package main

import (
	_i0 "github.com/dcaiafa/loxlex/simplelexer"
)

var _rules = []int32{
	0, 1, 2, 2, 3, 3, 3, 4, 5, 6, 7, 7, 7, 7,
	7, 7, 8, 8, 8, 9, 10, 10, 11, 11, 12, 12,
}

var _termCounts = []int32{
	1, 1, 3, 1, 1, 1, 0, 1, 3, 4, 3, 3, 3, 3,
	3, 1, 1, 1, 1, 1, 1, 1, 1, 0, 3, 1,
}

var _actions = []int32{
	37, 44, 49, 54, 59, 62, 67, 72, 37, 77, 86, 97, 102, 121,
	77, 138, 155, 168, 185, 202, 219, 236, 241, 244, 77, 77, 77, 77,
	257, 274, 77, 285, 302, 319, 336, 353, 370, 6, 0, -6, 12, 1,
	13, -6, 4, 5, 9, 2, 10, 4, 0, -7, 13, -7, 4, 0,
	-4, 13, -4, 2, 0, 2147483647, 4, 0, -3, 13, -3, 4, 0, -1,
	13, 8, 4, 0, -5, 13, -5, 8, 12, 12, 10, 13, 2, 14,
	11, 15, 10, 3, -23, 12, 12, 10, 13, 2, 14, 11, 15, 4,
	0, -2, 13, -2, 18, 4, -19, 3, -19, 9, -19, 0, -19, 7,
	-19, 13, -19, 2, 10, 6, -19, 8, -19, 16, 4, -20, 3, -20,
	9, -20, 0, -20, 7, -20, 13, -20, 6, -20, 8, -20, 16, 4,
	-21, 3, -21, 9, -21, 0, -21, 7, -21, 13, -21, 6, -21, 8,
	-21, 12, 9, 24, 0, -8, 7, 25, 13, -8, 6, 26, 8, 27,
	16, 4, -16, 3, -16, 9, -16, 0, -16, 7, -16, 13, -16, 6,
	-16, 8, -16, 16, 4, -17, 3, -17, 9, -17, 0, -17, 7, -17,
	13, -17, 6, -17, 8, -17, 16, 4, -15, 3, -15, 9, -15, 0,
	-15, 7, -15, 13, -15, 6, -15, 8, -15, 16, 4, -18, 3, -18,
	9, -18, 0, -18, 7, -18, 13, -18, 6, -18, 8, -18, 4, 4,
	30, 3, -22, 2, 3, 28, 12, 4, -25, 3, -25, 9, 24, 7,
	25, 6, 26, 8, 27, 16, 4, -9, 3, -9, 9, -9, 0, -9,
	7, -9, 13, -9, 6, -9, 8, -9, 10, 3, 31, 9, 24, 7,
	25, 6, 26, 8, 27, 16, 4, -14, 3, -14, 9, -14, 0, -14,
	7, -14, 13, -14, 6, -14, 8, -14, 16, 4, -10, 3, -10, 9,
	24, 0, -10, 7, -10, 13, -10, 6, -10, 8, 27, 16, 4, -11,
	3, -11, 9, 24, 0, -11, 7, -11, 13, -11, 6, -11, 8, 27,
	16, 4, -12, 3, -12, 9, -12, 0, -12, 7, -12, 13, -12, 6,
	-12, 8, -12, 16, 4, -13, 3, -13, 9, -13, 0, -13, 7, -13,
	13, -13, 6, -13, 8, -13, 12, 4, -24, 3, -24, 9, 24, 7,
	25, 6, 26, 8, 27,
}

var _goto = []int32{
	37, 50, 50, 50, 50, 50, 50, 50, 51, 60, 71, 50, 50, 50,
	86, 50, 50, 50, 50, 50, 50, 50, 50, 50, 97, 108, 119, 130,
	50, 50, 141, 50, 50, 50, 50, 50, 50, 12, 6, 2, 4, 3,
	1, 4, 3, 5, 2, 6, 5, 7, 0, 8, 6, 2, 4, 3,
	3, 11, 5, 7, 10, 7, 16, 6, 17, 10, 18, 8, 19, 9,
	20, 14, 12, 21, 11, 22, 7, 23, 6, 17, 10, 18, 8, 19,
	9, 20, 10, 7, 29, 6, 17, 10, 18, 8, 19, 9, 20, 10,
	7, 35, 6, 17, 10, 18, 8, 19, 9, 20, 10, 7, 33, 6,
	17, 10, 18, 8, 19, 9, 20, 10, 7, 32, 6, 17, 10, 18,
	8, 19, 9, 20, 10, 7, 34, 6, 17, 10, 18, 8, 19, 9,
	20, 10, 7, 36, 6, 17, 10, 18, 8, 19, 9, 20,
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
			_cast[[]Statement](p._stack.Peek(0).Sym),
		)
	case 2:
		return p.on_stmts__1(
			_cast[[]Statement](p._stack.Peek(2).Sym),
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
		return p.on_stmt__empty()
	case 7:
		return p.on_func_call_stmt(
			_cast[*FuncCall](p._stack.Peek(0).Sym),
		)
	case 8:
		return p.on_var_assign(
			_cast[_i0.Token](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 9:
		return p.on_func_call(
			_cast[_i0.Token](p._stack.Peek(3).Sym),
			_cast[_i0.Token](p._stack.Peek(2).Sym),
			_cast[[]Expr](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 10:
		return p.on_expr__bin(
			_cast[Expr](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 11:
		return p.on_expr__bin(
			_cast[Expr](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[Expr](p._stack.Peek(0).Sym),
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
		return p.on_expr__paren(
			_cast[_i0.Token](p._stack.Peek(2).Sym),
			_cast[Expr](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 15:
		return p.on_expr__simple(
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 16:
		return p.on_simple_expr(
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 17:
		return p.on_simple_expr(
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 18:
		return p.on_simple_expr(
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 19:
		return p.on_var_ref(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 20:
		return p.on_literal(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 21:
		return p.on_literal(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 22: // ZeroOrOne
		return _cast[[]Expr](p._stack.Peek(0).Sym)
	case 23: // ZeroOrOne
		{
			var zero []Expr
			return zero
		}
	case 24: // List
		return append(
			_cast[[]Expr](p._stack.Peek(2).Sym),
			_cast[Expr](p._stack.Peek(0).Sym),
		)
	case 25: // List
		return []Expr{
			_cast[Expr](p._stack.Peek(0).Sym),
		}
	default:
		panic("unreachable")
	}
}
