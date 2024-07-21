package main

const (
	EOF    int = 0
	ERROR  int = 1
	OPAREN int = 2
	CPAREN int = 3
	COMMA  int = 4
	ASSIGN int = 5
	OCURLY int = 6
	CCURLY int = 7
	PLUS   int = 8
	MINUS  int = 9
	TIMES  int = 10
	DIV    int = 11
	LT     int = 12
	LE     int = 13
	GT     int = 14
	GE     int = 15
	EQ     int = 16
	INT    int = 17
	STR    int = 18
	WHILE  int = 19
	OR     int = 20
	AND    int = 21
	ID     int = 22
	NL     int = 23
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
	case OCURLY:
		return "OCURLY"
	case CCURLY:
		return "CCURLY"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case TIMES:
		return "TIMES"
	case DIV:
		return "DIV"
	case LT:
		return "LT"
	case LE:
		return "LE"
	case GT:
		return "GT"
	case GE:
		return "GE"
	case EQ:
		return "EQ"
	case INT:
		return "INT"
	case STR:
		return "STR"
	case WHILE:
		return "WHILE"
	case OR:
		return "OR"
	case AND:
		return "AND"
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
