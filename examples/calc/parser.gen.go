package main

import (
	_i0 "github.com/dcaiafa/loxlex/simplelexer"
)

var _rules = []int32{
	0, 1, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3,
}

var _termCounts = []int32{
	1, 1, 3, 3, 3, 3, 3, 3, 3, 1, 1, 2,
}

var _actions = []int32{
	22, 29, 22, 46, 49, 52, 67, 22, 22, 22, 22, 22, 22, 84,
	99, 116, 133, 150, 167, 184, 201, 218, 6, 10, 1, 8, 2, 3,
	4, 16, 2, -10, 9, -10, 5, -10, 0, -10, 4, -10, 7, -10,
	6, -10, 3, -10, 2, 0, 2147483647, 2, 10, 14, 14, 2, 7, 5,
	8, 0, -1, 4, 9, 7, 10, 6, 11, 3, 12, 16, 2, -9,
	9, -9, 5, -9, 0, -9, 4, -9, 7, -9, 6, -9, 3, -9,
	14, 2, 7, 9, 15, 5, 8, 4, 9, 7, 10, 6, 11, 3,
	12, 16, 2, -11, 9, -11, 5, -11, 0, -11, 4, -11, 7, -11,
	6, -11, 3, -11, 16, 2, -8, 9, -8, 5, -8, 0, -8, 4,
	-8, 7, -8, 6, -8, 3, -8, 16, 2, -2, 9, -2, 5, 8,
	0, -2, 4, 9, 7, 10, 6, 11, 3, -2, 16, 2, -3, 9,
	-3, 5, 8, 0, -3, 4, 9, 7, 10, 6, 11, 3, -3, 16,
	2, -4, 9, -4, 5, -4, 0, -4, 4, -4, 7, 10, 6, -4,
	3, -4, 16, 2, -5, 9, -5, 5, -5, 0, -5, 4, -5, 7,
	10, 6, -5, 3, -5, 16, 2, -6, 9, -6, 5, -6, 0, -6,
	4, -6, 7, 10, 6, -6, 3, -6, 16, 2, -7, 9, -7, 5,
	-7, 0, -7, 4, -7, 7, -7, 6, -7, 3, -7,
}

var _goto = []int32{
	22, 29, 30, 29, 29, 29, 29, 35, 40, 45, 50, 55, 60, 29,
	29, 29, 29, 29, 29, 29, 29, 29, 6, 1, 3, 2, 5, 3,
	6, 0, 4, 2, 13, 3, 6, 4, 2, 16, 3, 6, 4, 2,
	19, 3, 6, 4, 2, 18, 3, 6, 4, 2, 21, 3, 6, 4,
	2, 20, 3, 6, 4, 2, 17, 3, 6,
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

func (p *calcParser) parse(lex _Lexer) bool {
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
func (p *calcParser) recoverLookahead(typ int, tok Token) {
	if p._qla != -1 {
		panic("recovered lookahead already pending")
	}

	p._qla = p._la
	p._qlasym = p._lasym
	p._la = typ
	p._lasym = tok
}

func (p *calcParser) _readToken() {
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

func (p *calcParser) _recover() bool {
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

func (p *calcParser) _makeError() Error {
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

func (p *calcParser) _act(prod int32) any {
	switch prod {
	case 1:
		return p.on_S__foo(
			_cast[float64](p._stack.Peek(0).Sym),
		)
	case 2:
		return p.on_expr__binary(
			_cast[float64](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[float64](p._stack.Peek(0).Sym),
		)
	case 3:
		return p.on_expr__binary(
			_cast[float64](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[float64](p._stack.Peek(0).Sym),
		)
	case 4:
		return p.on_expr__binary(
			_cast[float64](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[float64](p._stack.Peek(0).Sym),
		)
	case 5:
		return p.on_expr__binary(
			_cast[float64](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[float64](p._stack.Peek(0).Sym),
		)
	case 6:
		return p.on_expr__binary(
			_cast[float64](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[float64](p._stack.Peek(0).Sym),
		)
	case 7:
		return p.on_expr__binary(
			_cast[float64](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[float64](p._stack.Peek(0).Sym),
		)
	case 8:
		return p.on_expr__paren(
			_cast[_i0.Token](p._stack.Peek(2).Sym),
			_cast[float64](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 9:
		return p.on_expr__num(
			_cast[float64](p._stack.Peek(0).Sym),
		)
	case 10:
		return p.on_num(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 11:
		return p.on_num__minus(
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	default:
		panic("unreachable")
	}
}
