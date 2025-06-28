package main

const (
	EOF      int = 0
	ERROR    int = 1
	OCURLY   int = 2
	CCURLY   int = 3
	OBRACKET int = 4
	CBRACKET int = 5
	COMMA    int = 6
	COLON    int = 7
	TRUE     int = 8
	FALSE    int = 9
	NULL     int = 10
	STRING   int = 11
	NUMBER   int = 12
)

func _TokenToString(t int) string {
	switch t {
	case EOF:
		return "EOF"
	case ERROR:
		return "ERROR"
	case OCURLY:
		return "OCURLY"
	case CCURLY:
		return "CCURLY"
	case OBRACKET:
		return "OBRACKET"
	case CBRACKET:
		return "CBRACKET"
	case COMMA:
		return "COMMA"
	case COLON:
		return "COLON"
	case TRUE:
		return "TRUE"
	case FALSE:
		return "FALSE"
	case NULL:
		return "NULL"
	case STRING:
		return "STRING"
	case NUMBER:
		return "NUMBER"
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
