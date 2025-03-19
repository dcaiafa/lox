package main

const (
	EOF     int = 0
	ERROR   int = 1
	ADD     int = 2
	SUB     int = 3
	MUL     int = 4
	DIV     int = 5
	REM     int = 6
	POW     int = 7
	O_PAREN int = 8
	C_PAREN int = 9
	NUM     int = 10
)

func _TokenToString(t int) string {
	switch t {
	case EOF:
		return "EOF"
	case ERROR:
		return "ERROR"
	case ADD:
		return "ADD"
	case SUB:
		return "SUB"
	case MUL:
		return "MUL"
	case DIV:
		return "DIV"
	case REM:
		return "REM"
	case POW:
		return "POW"
	case O_PAREN:
		return "O_PAREN"
	case C_PAREN:
		return "C_PAREN"
	case NUM:
		return "NUM"
	default:
		return "???"
	}
}

type _Stack[T any] []T

func (s *_Stack[T]) Push(x T) {
	*s = append(*s, x)
}

func (s *_Stack[T]) Pop(n int) {
	*s = (*s)[:len(*s)-n]
}

func (s _Stack[T]) Peek(n int) T {
	return s[len(s)-n-1]
}

func (s _Stack[T]) PeekSlice(n int) []T {
	return s[len(s)-n:]
}
