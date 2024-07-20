package parser

const (
	EOF             int = 0
	ERROR           int = 1
	COMMA           int = 2
	EQ              int = 3
	OR              int = 4
	OCURLY          int = 5
	CCURLY          int = 6
	TILDE           int = 7
	OPAREN          int = 8
	CPAREN          int = 9
	SUB             int = 10
	ZERO_OR_ONE     int = 11
	ZERO_OR_MORE    int = 12
	ONE_OR_MORE     int = 13
	ZERO_OR_MORE_NG int = 14
	ONE_OR_MORE_NG  int = 15
	PARSER          int = 16
	LEXER           int = 17
	START           int = 18
	DISCARD         int = 19
	MACRO           int = 20
	FRAG            int = 21
	MODE            int = 22
	PUSH_MODE       int = 23
	POP_MODE        int = 24
	ERROR_KEYWORD   int = 25
	LEFT            int = 26
	LIST            int = 27
	RIGHT           int = 28
	EMIT            int = 29
	EMPTY           int = 30
	EXTERNAL        int = 31
	KEYWORD         int = 32
	ID              int = 33
	NUM             int = 34
	LITERAL         int = 35
	OBRACKET        int = 36
	CBRACKET        int = 37
	CLASS_DASH      int = 38
	CLASS_CHAR      int = 39
	NL              int = 40
	EXTEND          int = 41
)

func _TokenToString(t int) string {
	switch t {
	case EOF:
		return "EOF"
	case ERROR:
		return "ERROR"
	case COMMA:
		return "COMMA"
	case EQ:
		return "EQ"
	case OR:
		return "OR"
	case OCURLY:
		return "OCURLY"
	case CCURLY:
		return "CCURLY"
	case TILDE:
		return "TILDE"
	case OPAREN:
		return "OPAREN"
	case CPAREN:
		return "CPAREN"
	case SUB:
		return "SUB"
	case ZERO_OR_ONE:
		return "ZERO_OR_ONE"
	case ZERO_OR_MORE:
		return "ZERO_OR_MORE"
	case ONE_OR_MORE:
		return "ONE_OR_MORE"
	case ZERO_OR_MORE_NG:
		return "ZERO_OR_MORE_NG"
	case ONE_OR_MORE_NG:
		return "ONE_OR_MORE_NG"
	case PARSER:
		return "PARSER"
	case LEXER:
		return "LEXER"
	case START:
		return "START"
	case DISCARD:
		return "DISCARD"
	case MACRO:
		return "MACRO"
	case FRAG:
		return "FRAG"
	case MODE:
		return "MODE"
	case PUSH_MODE:
		return "PUSH_MODE"
	case POP_MODE:
		return "POP_MODE"
	case ERROR_KEYWORD:
		return "ERROR_KEYWORD"
	case LEFT:
		return "LEFT"
	case LIST:
		return "LIST"
	case RIGHT:
		return "RIGHT"
	case EMIT:
		return "EMIT"
	case EMPTY:
		return "EMPTY"
	case EXTERNAL:
		return "EXTERNAL"
	case KEYWORD:
		return "KEYWORD"
	case ID:
		return "ID"
	case NUM:
		return "NUM"
	case LITERAL:
		return "LITERAL"
	case OBRACKET:
		return "OBRACKET"
	case CBRACKET:
		return "CBRACKET"
	case CLASS_DASH:
		return "CLASS_DASH"
	case CLASS_CHAR:
		return "CLASS_CHAR"
	case NL:
		return "NL"
	case EXTEND:
		return "EXTEND"
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
