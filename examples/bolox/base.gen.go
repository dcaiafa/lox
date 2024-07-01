package main

const (
	EOF    int = 0
	ERROR  int = 1
	OPAREN int = 2
	CPAREN int = 3
	COMMA  int = 4
	ASSIGN int = 5
	INT    int = 6
	STR    int = 7
	ID     int = 8
	NL     int = 9
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
