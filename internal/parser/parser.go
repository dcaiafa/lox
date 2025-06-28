package parser

import (
	gotoken "go/token"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/dcaiafa/loxlex/simplelexer"

	"github.com/dcaiafa/lox/internal/ast"
	"github.com/dcaiafa/lox/internal/base/errlogger"
)

type Token = simplelexer.Token

type parser struct {
	lox
	file *gotoken.File
	errs *errlogger.ErrLogger
	unit *ast.Unit
}

func Parse(
	file *gotoken.File,
	data []byte,
	errs *errlogger.ErrLogger,
) *ast.Unit {
	lex := newLexer(simplelexer.Config{
		StateMachine: new(_LexerStateMachine),
		File:         file,
		Input:        data,
	})

	p := &parser{
		file: file,
		errs: errs,
	}

	ok := p.parse(lex)
	if !ok {
		errs.Errorf(0, "Failed to parse")
	}

	return p.unit
}

func (p *parser) on_spec(_ []Token, sections [][]ast.Statement) *ast.Unit {
	n := 0
	for _, section := range sections {
		n += len(section)
	}
	stmts := make([]ast.Statement, 0, n)
	for _, section := range sections {
		stmts = append(stmts, section...)
	}
	p.unit = &ast.Unit{
		Statements: stmts,
	}
	return p.unit
}

func (p *parser) on_spec__error(e Error) *ast.Unit {
	if e.Token.Type == ERROR {
		p.errs.Errorf(e.Token.Pos, "lexer error: %v", e.Token.Err)
	} else {
		p.errs.Errorf(e.Token.Pos, "unexpected %v", _TokenToString(e.Token.Type))
	}
	return &ast.Unit{}
}

func (p *parser) on_section(sectionStmts []ast.Statement) []ast.Statement {
	return sectionStmts
}

// Parser
// ======

func (p *parser) on_parser_section(_, _ Token, stmts []ast.Statement) []ast.Statement {
	return stmts
}

func (p *parser) on_parser_statement(s ast.Statement) ast.Statement {
	return s
}

func (p *parser) on_parser_statement__nl(_ Token) ast.Statement {
	// This matches an empty line. Return a dummy statement that will be discarded
	// because of the *! cardinality.
	return ast.DiscardStatementSingleton
}

func (p *parser) on_parser_rule(start Token, name Token, _ Token, prods []*ast.ParserProd, _ Token) *ast.ParserRule {
	return &ast.ParserRule{
		IsStart: start.Type == START,
		Name:    string(name.Str),
		Prods:   prods,
	}
}

func (p *parser) on_parser_prod(terms []*ast.ParserTerm, qualif *ast.ProdQualifier) *ast.ParserProd {
	return &ast.ParserProd{
		Terms:     terms,
		Qualifier: qualif,
	}
}

func (p *parser) on_parser_prod__empty(_ Token) *ast.ParserProd {
	return &ast.ParserProd{}
}

func (p *parser) on_parser_term_card(term *ast.ParserTerm, typ ast.ParserTermType) *ast.ParserTerm {
	if typ == ast.ParserTermSimple || typ == ast.ParserTermError {
		return term
	}
	if term.Type == ast.ParserTermList {
		if typ != ast.ParserTermZeroOrOne {
			p.errs.Errorf(
				term.Bounds().Begin,
				"@list term can only use the zero-or-more '?' cardinality")
			return term
		}
		term.Type = ast.ParserTermListOpt
		return term
	}
	return &ast.ParserTerm{
		Type:  typ,
		Child: term,
	}
}

func (p *parser) on_parser_term__token(tok Token) *ast.ParserTerm {
	switch tok.Type {
	case ID:
		return &ast.ParserTerm{Name: string(tok.Str)}
	case LITERAL:
		return &ast.ParserTerm{Alias: fixLiteral(tok.Str)}
	case ERROR_KEYWORD:
		return &ast.ParserTerm{Type: ast.ParserTermError}
	default:
		panic("not-reached")
	}
}

func (p *parser) on_parser_term__list(listTerm *ast.ParserTerm) *ast.ParserTerm {
	return listTerm
}

func (p *parser) on_parser_list(_, _ Token, elem *ast.ParserTerm, _ Token, sep *ast.ParserTerm, _ Token) *ast.ParserTerm {
	return &ast.ParserTerm{
		Type:  ast.ParserTermList,
		Child: elem,
		Sep:   sep,
	}
}

func (p *parser) on_parser_card(card Token) ast.ParserTermType {
	switch card.Type {
	case ZERO_OR_MORE:
		return ast.ParserTermZeroOrMore
	case ZERO_OR_MORE_F:
		return ast.ParserTermZeroOrMoreF
	case ONE_OR_MORE:
		return ast.ParserTermOneOrMore
	case ZERO_OR_ONE:
		return ast.ParserTermZeroOrOne
	default:
		panic("unreachable")
	}
}

func (p *parser) on_parser_qualif(assoc Token, _ Token, prec Token, _ Token) *ast.ProdQualifier {
	q := &ast.ProdQualifier{}

	switch assoc.Type {
	case LEFT:
		q.Associativity = ast.Left
	case RIGHT:
		q.Associativity = ast.Right
	default:
		panic("not-reached")
	}

	var err error
	q.Precedence, err = strconv.Atoi(string(prec.Str))
	if err != nil {
		panic(err)
	}
	if q.Precedence <= 0 {
		panic("not-reached")
	}

	return q
}

// Lexer
// =====

func (p *parser) on_lexer_section(_, _ Token, stmts []ast.Statement) []ast.Statement {
	return stmts
}

func (p *parser) on_lexer_statement(s ast.Statement) ast.Statement {
	return s
}

func (p *parser) on_lexer_rule(r ast.Statement) ast.Statement {
	return r
}

func (p *parser) on_lexer_rule__nl(_ Token) ast.Statement {
	return &ast.DiscardStatement{}
}

func (p *parser) on_mode(_ Token, name Token, _ []Token, _ Token, rules []ast.Statement, _ Token) *ast.Mode {
	return &ast.Mode{
		Name:  string(name.Str),
		Rules: rules,
	}
}

func (p *parser) on_token_rule(id Token, _ Token, expr *ast.LexerExpr, actions []ast.Action, _ Token) *ast.TokenRule {
	return &ast.TokenRule{
		Name:    string(id.Str),
		Expr:    expr,
		Actions: actions,
	}
}

func (p *parser) on_frag_rule(_ Token, expr *ast.LexerExpr, actions []ast.Action, _ Token) *ast.FragRule {
	return &ast.FragRule{
		Expr:    expr,
		Actions: actions,
	}
}

func (p *parser) on_macro_rule(_ Token, name Token, _ Token, expr *ast.LexerExpr, _ Token) *ast.MacroRule {
	return &ast.MacroRule{
		Name: string(name.Str),
		Expr: expr,
	}
}

func (p *parser) on_external_rule(_ Token, names []*ast.ExternalName, _ Token) *ast.ExternalRule {
	return &ast.ExternalRule{
		Names: names,
	}
}

func (p *parser) on_external_name(name Token) *ast.ExternalName {
	return &ast.ExternalName{
		Name: string(name.Str),
	}
}

func (p *parser) on_lexer_expr(factors []*ast.LexerFactor) *ast.LexerExpr {
	return &ast.LexerExpr{
		Factors: factors,
	}
}

func (p *parser) on_lexer_factor(terms []*ast.LexerTermCard) *ast.LexerFactor {
	return &ast.LexerFactor{
		Terms: terms,
	}
}

func (p *parser) on_lexer_term_card(term ast.LexerTerm, card ast.Card) *ast.LexerTermCard {
	return &ast.LexerTermCard{
		Term: term,
		Card: card,
	}
}

func (p *parser) on_lexer_card(c Token) ast.Card {
	switch c.Type {
	case ZERO_OR_ONE:
		return ast.ZeroOrOne
	case ZERO_OR_MORE:
		return ast.ZeroOrMore
	case ZERO_OR_MORE_NG:
		return ast.ZeroOrMoreNG
	case ONE_OR_MORE:
		return ast.OneOrMore
	case ONE_OR_MORE_NG:
		return ast.OneOrMoreNG
	default:
		panic("unreachable")
	}
}

func (p *parser) on_lexer_term__tok(tok Token) ast.LexerTerm {
	switch tok.Type {
	case LITERAL:
		return &ast.LexerTermLiteral{
			Literal: fixLiteral(tok.Str),
		}
	case ID:
		return &ast.LexerTermRef{
			Ref: string(tok.Str),
		}
	case DOT:
		// Allow any unicode code point.
		return &ast.LexerTermCharClass{
			Expr: &ast.CharClass{
				CharClassItems: []*ast.CharClassItem{
					{
						From: 0x000000,
						To:   0x10FFFF,
					},
				},
			},
		}

	default:
		panic("unreachable")
	}
}

func (p *parser) on_lexer_term__char_class_expr(e ast.CharClassExpr) ast.LexerTerm {
	return &ast.LexerTermCharClass{
		Expr: e,
	}
}

func (p *parser) on_lexer_term__expr(_ Token, expr *ast.LexerExpr, _ Token) ast.LexerTerm {
	return expr
}

func (p *parser) on_char_class_expr__binary(l ast.CharClassExpr, _ Token, r ast.CharClassExpr) ast.CharClassExpr {
	return &ast.CharClassBinaryExpr{
		Op:    ast.CharClassBinaryExprSub,
		Left:  l,
		Right: r,
	}
}

func (p *parser) on_char_class_expr__char_class(c *ast.CharClass) ast.CharClassExpr {
	return c
}

func (p *parser) on_char_class(neg Token, _ Token, chars []Token, _ Token) *ast.CharClass {
	items := make([]*ast.CharClassItem, 0, len(chars))

	toRune := func(t Token) rune {
		r, _ := utf8.DecodeRuneInString(unescape(t.Str))
		return r
	}
	addItem := func(b, e Token) {
		items = append(items, &ast.CharClassItem{
			From: toRune(b),
			To:   toRune(e),
		})
	}

	for i := 0; i < len(chars); i++ {
		if i+2 > len(chars)-1 || chars[i+1].Type != CLASS_DASH {
			addItem(chars[i], chars[i])
		} else {
			addItem(chars[i], chars[i+2])
			i += 2
		}
	}

	return &ast.CharClass{
		Neg:            neg.Type != EOF,
		CharClassItems: items,
	}
}

func (p *parser) on_char_class_item(c Token) Token {
	return c
}

func (p *parser) on_action(action ast.Action) ast.Action {
	return action
}

func (p *parser) on_action_discard(_ Token) *ast.ActionDiscard {
	return &ast.ActionDiscard{}
}

func (p *parser) on_action_push_mode(_, _ Token, mode Token, _ Token) *ast.ActionPushMode {
	m := &ast.ActionPushMode{}
	// The `mode` token is optional. If not specified, mode.Type will be 0 (EOF).
	if mode.Type == ID {
		m.Mode = string(mode.Str)
	} else {
		m.Mode = ast.DefaultModeName
	}
	return m
}

func (p *parser) on_action_pop_mode(_ Token) *ast.ActionPopMode {
	return &ast.ActionPopMode{}
}

func (p *parser) on_action_emit(_, _ Token, tok Token, _ Token) *ast.ActionEmit {
	return &ast.ActionEmit{
		Name: string(tok.Str),
	}
}

func (p *parser) _onBounds(r any, begin, end Token) {
	rAST, ok := r.(ast.AST)
	if !ok {
		return
	}

	rAST.SetBounds(ast.Bounds{
		Begin: begin.Pos,
		End:   end.Pos + gotoken.Pos(len(end.Str)),
	})
}

func fixLiteral(lit []byte) string {
	return unescape(lit[1 : len(lit)-1])
}

// unescape returns a literal with all escape sequences evaluated and replaced.
// Because the literal has already been validated by the lexer, unescape assume
// that the escape sequences are well-formed.
func unescape(lit []byte) string {
	var str strings.Builder

	for i := 0; i < len(lit); i++ {
		if lit[i] == '\\' {
			switch lit[i+1] {
			case 'n':
				str.WriteRune('\n')
				i++
			case 'r':
				str.WriteRune('\r')
				i++
			case 't':
				str.WriteRune('\t')
				i++
			case '\'':
				str.WriteRune('\'')
				i++
			case '\\':
				str.WriteRune('\\')
				i++
			case '-':
				str.WriteRune('-')
				i++
			case 'x':
				str.WriteByte(byte(hexToRune(string(lit[i+2 : i+4]))))
				i += 3
			case 'u':
				str.WriteRune(hexToRune(string(lit[i+2 : i+6])))
				i += 5
			case 'U':
				str.WriteRune(hexToRune(string(lit[i+2 : i+10])))
				i += 9
			default:
				panic("unreachable")
			}
		} else {
			str.WriteByte(lit[i])
		}
	}

	return str.String()
}

func hexToRune(str string) rune {
	v, err := strconv.ParseUint(string(str), 16, 32)
	if err != nil {
		panic(err)
	}
	return rune(v)
}
