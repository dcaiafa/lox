package main

import (
	_i0 "errors"
	_i1 "github.com/dcaiafa/lox/internal/base/baselexer"
)

var _LHS = []int32{
	0, 1, 2, 2, 2, 2, 2, 2, 2, 3, 4, 5, 6, 6,
	7, 7, 8, 8, 9, 9,
}

var _TermCounts = []int32{
	1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 3, 3, 1, 0,
	3, 1, 1, 0, 3, 1,
}

var _Actions = []int32{
	27, 42, 51, 60, 69, 86, 91, 100, 109, 118, 121, 130, 133, 138,
	141, 144, 149, 154, 157, 162, 27, 171, 180, 27, 183, 188, 193, 14,
	9, 1, 10, 2, 12, 3, 4, 4, 2, 5, 11, 6, 8, 7,
	8, 5, -7, 3, -7, 6, -7, 0, -7, 8, 5, -8, 3, -8,
	6, -8, 0, -8, 8, 5, -5, 3, -5, 6, -5, 0, -5, 16,
	5, -17, 9, 1, 10, 2, 12, 3, 4, 4, 2, 5, 11, 6,
	8, 7, 4, 3, -13, 11, 14, 8, 5, -4, 3, -4, 6, -4,
	0, -4, 8, 5, -6, 3, -6, 6, -6, 0, -6, 8, 5, -3,
	3, -3, 6, -3, 0, -3, 2, 0, 2147483647, 8, 5, -2, 3, -2,
	6, -2, 0, -2, 2, 0, -1, 4, 3, -12, 6, 22, 2, 3,
	19, 2, 7, 20, 4, 3, -15, 6, -15, 4, 5, -16, 6, 23,
	2, 5, 21, 4, 5, -19, 6, -19, 8, 5, -9, 3, -9, 6,
	-9, 0, -9, 8, 5, -11, 3, -11, 6, -11, 0, -11, 2, 11,
	14, 4, 3, -10, 6, -10, 4, 3, -14, 6, -14, 4, 5, -18,
	6, -18,
}

var _Goto = []int32{
	27, 36, 36, 36, 37, 48, 36, 36, 36, 36, 36, 36, 36, 36,
	36, 36, 36, 36, 36, 36, 55, 36, 62, 65, 36, 36, 36, 8,
	5, 8, 1, 9, 3, 10, 2, 11, 0, 10, 9, 16, 8, 17,
	5, 8, 3, 10, 2, 18, 6, 7, 12, 6, 13, 4, 15, 6,
	5, 8, 3, 10, 2, 24, 2, 4, 25, 6, 5, 8, 3, 10,
	2, 26,
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
	_lex   _Lexer
	_state _Stack[int32]
	_sym   _Stack[any]

	_lookahead     Token
	_lookaheadType int
	_errorToken    Token
}

func (p *jsonParser) parse(lex _Lexer) bool {
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
			p._ReadToken()
		} else { // reduce
			prod := -action
			termCount := _TermCounts[int(prod)]
			rule := _LHS[int(prod)]
			res := p._Act(prod)
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

func (p *jsonParser) errorToken() Token {
	return p._errorToken
}

func (p *jsonParser) _ReadToken() {
	p._lookahead, p._lookaheadType = p._lex.ReadToken()
}

func (p *jsonParser) _Recover() (int32, bool) {
	p._errorToken = p._lookahead
	p._onError()

	for {
		for p._lookaheadType == ERROR {
			p._ReadToken()
		}

		saveState := p._state
		saveSym := p._sym

		for len(p._state) > 1 {
			topState := p._state.Peek(0)
			action, ok := _Find(_Actions, topState, int32(ERROR))
			if ok {
				action2, ok := _Find(
					_Actions, action, int32(p._lookaheadType))
				if ok {
					p._state.Push(action)
					p._sym.Push(_errorPlaceholder)
					return action2, true
				}
			}
			p._state.Pop(1)
			p._sym.Pop(1)
		}

		if p._lookaheadType == EOF {
			return 0, false
		}

		p._ReadToken()
		p._state = saveState
		p._sym = saveSym
	}
}

func (p *jsonParser) _Act(prod int32) any {
	switch prod {
	case 1:
		return p.on_json(
			_cast[any](p._sym.Peek(0)),
		)
	case 2:
		return p.on_value__object(
			_cast[map[string]any](p._sym.Peek(0)),
		)
	case 3:
		return p.on_value__array(
			_cast[[]any](p._sym.Peek(0)),
		)
	case 4:
		return p.on_value__tok(
			_cast[_i1.Token](p._sym.Peek(0)),
		)
	case 5:
		return p.on_value__tok(
			_cast[_i1.Token](p._sym.Peek(0)),
		)
	case 6:
		return p.on_value__tok(
			_cast[_i1.Token](p._sym.Peek(0)),
		)
	case 7:
		return p.on_value__tok(
			_cast[_i1.Token](p._sym.Peek(0)),
		)
	case 8:
		return p.on_value__tok(
			_cast[_i1.Token](p._sym.Peek(0)),
		)
	case 9:
		return p.on_object(
			_cast[_i1.Token](p._sym.Peek(2)),
			_cast[[]member](p._sym.Peek(1)),
			_cast[_i1.Token](p._sym.Peek(0)),
		)
	case 10:
		return p.on_member(
			_cast[_i1.Token](p._sym.Peek(2)),
			_cast[_i1.Token](p._sym.Peek(1)),
			_cast[any](p._sym.Peek(0)),
		)
	case 11:
		return p.on_array(
			_cast[_i1.Token](p._sym.Peek(2)),
			_cast[[]any](p._sym.Peek(1)),
			_cast[_i1.Token](p._sym.Peek(0)),
		)
	case 12: // ZeroOrOne
		return _cast[[]member](p._sym.Peek(0))
	case 13: // ZeroOrOne
		{
			var zero []member
			return zero
		}
	case 14: // List
		return append(
			_cast[[]member](p._sym.Peek(2)),
			_cast[member](p._sym.Peek(0)),
		)
	case 15: // List
		return []member{
			_cast[member](p._sym.Peek(0)),
		}
	case 16: // ZeroOrOne
		return _cast[[]any](p._sym.Peek(0))
	case 17: // ZeroOrOne
		{
			var zero []any
			return zero
		}
	case 18: // List
		return append(
			_cast[[]any](p._sym.Peek(2)),
			_cast[any](p._sym.Peek(0)),
		)
	case 19: // List
		return []any{
			_cast[any](p._sym.Peek(0)),
		}
	default:
		panic("unreachable")
	}
}
