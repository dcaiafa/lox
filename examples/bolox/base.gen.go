package main

const (
	EOF       int = 0
	ERROR     int = 1
	WHILE     int = 2
	OR        int = 3
	AND       int = 4
	TRUE      int = 5
	FALSE     int = 6
	IF        int = 7
	ELIF      int = 8
	ELSE      int = 9
	NIL       int = 10
	CONTINUE  int = 11
	ID        int = 12
	OPAREN    int = 13
	CPAREN    int = 14
	COMMA     int = 15
	ASSIGN    int = 16
	OCURLY    int = 17
	CCURLY    int = 18
	PLUS      int = 19
	MINUS     int = 20
	TIMES     int = 21
	DIV       int = 22
	LT        int = 23
	LE        int = 24
	GT        int = 25
	GE        int = 26
	EQ        int = 27
	INT       int = 28
	STR_BEGIN int = 29
	STR_END   int = 30
	CHAR_SEQ  int = 31
	NL        int = 32
)

func _TokenToString(t int) string {
	switch t {
	case EOF:
		return "EOF"
	case ERROR:
		return "ERROR"
	case WHILE:
		return "WHILE"
	case OR:
		return "OR"
	case AND:
		return "AND"
	case TRUE:
		return "TRUE"
	case FALSE:
		return "FALSE"
	case IF:
		return "IF"
	case ELIF:
		return "ELIF"
	case ELSE:
		return "ELSE"
	case NIL:
		return "NIL"
	case CONTINUE:
		return "CONTINUE"
	case ID:
		return "ID"
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
	case STR_BEGIN:
		return "STR_BEGIN"
	case STR_END:
		return "STR_END"
	case CHAR_SEQ:
		return "CHAR_SEQ"
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
