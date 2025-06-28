package main

var _rules = []int32{
	0, 1, 2, 2, 2, 2, 2, 2, 2, 3, 4, 5, 6, 6,
	7, 7, 8, 8, 9, 9,
}

var _termCounts = []int32{
	1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 3, 3, 1, 0,
	3, 1, 1, 0, 3, 1,
}

var _actions = []int32{
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

var _goto = []int32{
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

func (p *jsoncParser) parse(lex _Lexer) bool {
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
func (p *jsoncParser) recoverLookahead(typ int, tok Token) {
	if p._qla != -1 {
		panic("recovered lookahead already pending")
	}

	p._qla = p._la
	p._qlasym = p._lasym
	p._la = typ
	p._lasym = tok
}

func (p *jsoncParser) _readToken() {
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

func (p *jsoncParser) _recover() bool {
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

func (p *jsoncParser) _makeError() Error {
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

func (p *jsoncParser) _act(prod int32) any {
	switch prod {
	case 1:
		return p.on_json(
			_cast[any](p._stack.Peek(0).Sym),
		)
	case 2:
		return p.on_value__object(
			_cast[map[string]any](p._stack.Peek(0).Sym),
		)
	case 3:
		return p.on_value__array(
			_cast[[]any](p._stack.Peek(0).Sym),
		)
	case 4:
		return p.on_value__tok(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 5:
		return p.on_value__tok(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 6:
		return p.on_value__tok(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 7:
		return p.on_value__tok(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 8:
		return p.on_value__tok(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 9:
		return p.on_object(
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[[]member](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 10:
		return p.on_member(
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[Token](p._stack.Peek(1).Sym),
			_cast[any](p._stack.Peek(0).Sym),
		)
	case 11:
		return p.on_array(
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[[]any](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 12: // ZeroOrOne
		return _cast[[]member](p._stack.Peek(0).Sym)
	case 13: // ZeroOrOne
		{
			var zero []member
			return zero
		}
	case 14: // List
		return append(
			_cast[[]member](p._stack.Peek(2).Sym),
			_cast[member](p._stack.Peek(0).Sym),
		)
	case 15: // List
		return []member{
			_cast[member](p._stack.Peek(0).Sym),
		}
	case 16: // ZeroOrOne
		return _cast[[]any](p._stack.Peek(0).Sym)
	case 17: // ZeroOrOne
		{
			var zero []any
			return zero
		}
	case 18: // List
		return append(
			_cast[[]any](p._stack.Peek(2).Sym),
			_cast[any](p._stack.Peek(0).Sym),
		)
	case 19: // List
		return []any{
			_cast[any](p._stack.Peek(0).Sym),
		}
	default:
		panic("unreachable")
	}
}
