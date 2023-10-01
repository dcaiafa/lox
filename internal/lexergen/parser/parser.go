package parser

import (
	gotoken "go/token"
	"strconv"
	"unicode/utf8"

	"github.com/dcaiafa/lox/internal/errlogger"
	"github.com/dcaiafa/lox/internal/lexergen/ast"
)

type parser struct {
	lox
	file *gotoken.File
	errs *errlogger.ErrLogger
	unit *ast.Unit
}

func Parse(file *gotoken.File, data []byte, errs *errlogger.ErrLogger) *ast.Unit {
	lex := newLex(file, data, errs)
	p := &parser{
		file: file,
		errs: errs,
	}
	p.parse(lex)
	return p.unit
}

func (p *parser) on_spec(sections [][]ast.Statement) *ast.Unit {
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

func (p *parser) on_section(sectionStmts []ast.Statement) []ast.Statement {
	return sectionStmts
}

// Parser
// ======

func (p *parser) on_parser_section(_ Token, stmts []ast.Statement) []ast.Statement {
	return stmts
}

func (p *parser) on_parser_statement(s ast.Statement) ast.Statement {
	return s
}

func (p *parser) on_parser_rule(name Token, _ Token, prods []*ast.Prod, _ Token) *ast.Rule {
	return &ast.Rule{
		Name:  name.Str,
		Prods: prods,
	}
}

func (p *parser) on_parser_prod(terms []*ast.Term, qualif *ast.ProdQualifier) *ast.Prod {
	return &ast.Prod{
		Terms:     terms,
		Qualifier: qualif,
	}
}

func (p *parser) on_parser_term_card(term *ast.Term, typ ast.TermType) *ast.Term {
	if typ == ast.Simple || typ == ast.Error {
		return term
	}
	return &ast.Term{
		Type:  typ,
		Child: term,
	}
}

func (p *parser) on_parser_term__token(tok Token) *ast.Term {
	switch tok.Type {
	case ID:
		return &ast.Term{Name: tok.Str}
	case LITERAL:
		return &ast.Term{Alias: tok.Str}
	case ERROR_KEYWORD:
		return &ast.Term{Type: ast.Error}
	default:
		panic("not-reached")
	}
}

func (p *parser) on_parser_term__list(listTerm *ast.Term) *ast.Term {
	return listTerm
}

func (p *parser) on_parser_list(_, _ Token, elem *ast.Term, _ Token, sep *ast.Term, _ Token) *ast.Term {
	return &ast.Term{
		Type:  ast.List,
		Child: elem,
		Sep:   sep,
	}
}

func (p *parser) on_parser_card(card Token) ast.TermType {
	switch card.Type {
	case ZERO_OR_MORE:
		return ast.TermZeroOrMore
	case ONE_OR_MORE:
		return ast.TermOneOrMore
	case ZERO_OR_ONE:
		return ast.TermZeroOrOne
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
	q.Precedence, err = strconv.Atoi(prec.Str)
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

func (p *parser) on_lexer_section(_ Token, stmts []ast.Statement) []ast.Statement {
	return stmts
}

func (p *parser) on_lexer_statement(s ast.Statement) ast.Statement {
	return s
}

func (p *parser) on_lexer_rule(r ast.Statement) ast.Statement {
	return r
}

func (p *parser) on_mode(_ Token, name Token, _ Token, rules []ast.Statement, _ Token) *ast.Mode {
	return &ast.Mode{
		Name:  name.Str,
		Rules: rules,
	}
}

func (p *parser) on_token_rule(id Token, _ Token, expr *ast.LexerExpr, actions []ast.Action, _ Token) *ast.TokenRule {
	return &ast.TokenRule{
		Name:    id.Str,
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
		Name: name.Str,
		Expr: expr,
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
	case ONE_OR_MORE:
		return ast.OneOrMore
	default:
		panic("unreachable")
	}
}

func (p *parser) on_lexer_term__tok(tok Token) ast.LexerTerm {
	switch tok.Type {
	case LITERAL:
		return &ast.LexerTermLiteral{
			Literal: tok.Str,
		}
	case ID:
		return &ast.LexerTermRef{
			Ref: tok.Str,
		}
	default:
		panic("unreachable")
	}
}

func (p *parser) on_lexer_term__char_class(charClass *ast.LexerTermCharClass) ast.LexerTerm {
	return charClass
}

func (p *parser) on_lexer_term__expr(_ Token, expr *ast.LexerExpr, _ Token) ast.LexerTerm {
	return expr
}

func (p *parser) on_char_class(neg Token, _ Token, chars []Token, _ Token) *ast.LexerTermCharClass {
	items := make([]*ast.CharClassItem, 0, len(chars))

	toRune := func(t Token) rune {
		r, _ := utf8.DecodeRuneInString(t.Str)
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

	return &ast.LexerTermCharClass{
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

func (p *parser) on_action_skip(_ Token) *ast.ActionSkip {
	return &ast.ActionSkip{}
}

func (p *parser) on_action_push_mode(_, _ Token, mode Token, _ Token) *ast.ActionPushMode {
	return &ast.ActionPushMode{
		Mode: mode.Str,
	}
}

func (p *parser) on_action_pop_mode(_ Token) *ast.ActionPopMode {
	return &ast.ActionPopMode{}
}

func (p *parser) onReduce(r any, begin, end Token) {
	rAST, ok := r.(ast.AST)
	if !ok {
		return
	}
	rAST.SetBounds(ast.Bounds{
		Begin: begin.Pos,
		End:   end.Pos,
	})
}

func (p *parser) onError() {
	p.errs.Errorf(
		p.file.Position(p.errorToken().Pos),
		"unexpected %v", p.errorToken())
}
