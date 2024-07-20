package main

const (
	EOF    int = 0
	ERROR  int = 1
	OPAREN int = 2
	CPAREN int = 3
	COMMA  int = 4
	ASSIGN int = 5
	PLUS   int = 6
	MINUS  int = 7
	TIMES  int = 8
	DIV    int = 9
	INT    int = 10
	STR    int = 11
	ID     int = 12
	NL     int = 13
)

func _TokenToString(t int) string {
	switch t {
	case EOF:
		return "EOF"
	case ERROR:
		return "ERROR"
	case OPAREN:
		return "OPAREN"
	case CPAREN:
		return "CPAREN"
	case COMMA:
		return "COMMA"
	case ASSIGN:
		return "ASSIGN"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case TIMES:
		return "TIMES"
	case DIV:
		return "DIV"
	case INT:
		return "INT"
	case STR:
		return "STR"
	case ID:
		return "ID"
	case NL:
		return "NL"
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
