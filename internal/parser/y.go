// Code generated by goyacc parser.y. DO NOT EDIT.

//line parser.y:2
package parser

import __yyfmt__ "fmt"

//line parser.y:2

import (
	"github.com/dcaiafa/lox/internal/grammar"
	"github.com/dcaiafa/lox/internal/token"
)

func cast[T any](v any) T {
	cv, _ := v.(T)
	return cv
}

/*

syntax = production* .

production = ID '=' expression '.' .

expression = term                  # e1
           | expression '|' term   # e2
           .

term = qfactor+ label? .

qfactor = factor qualifier? .

factor = ID | LITERAL .

qualifier = '+' | '*' | '?' .

label = '#' ID .

---

func syntax(t0 []any) any {
  return nil
}

func production(id0 Token, t1 Token, expression2 any, t3 token) any {
  return nil
}

func expression_e1(t0 any) any {
}

func expression_e2(expression0 any, t1 Token, term2 any) any {
}

func term(qfactor []any) any {
}

func qfactor(factor any, qualifier any) any {
}

func factor(t0 Token) any {
}

func qualifier(t0 Token) any {
}

func label(t0 Token, id Token) any {
}

*/

//line parser.y:69
type yySymType struct {
	yys  int
	tok  token.Token
	prod interface{}
}

const EOF = 57346
const LEXERR = 57347
const ID = 57348
const LITERAL = 57349

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"EOF",
	"LEXERR",
	"ID",
	"LITERAL",
	"'='",
	"'.'",
	"'|'",
	"'*'",
	"'+'",
	"'?'",
	"'#'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 31

var yyAct = [...]int8{
	11, 15, 16, 25, 26, 27, 9, 13, 29, 22,
	17, 18, 15, 16, 6, 7, 5, 1, 19, 28,
	20, 8, 21, 23, 24, 14, 12, 10, 3, 4,
	2,
}

var yyPact = [...]int16{
	8, -1000, 11, -1000, 8, -1000, -2, -1000, -1000, 6,
	1, -1000, -5, -1000, -8, -1000, -1000, -1000, 6, -1000,
	-1000, -1000, 2, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
}

var yyPgo = [...]int8{
	0, 30, 29, 28, 16, 27, 0, 26, 7, 25,
	24, 23, 22, 18, 17,
}

var yyR1 = [...]int8{
	0, 14, 1, 2, 2, 3, 3, 4, 5, 5,
	6, 7, 7, 8, 9, 9, 10, 10, 10, 11,
	11, 12, 13, 13,
}

var yyR2 = [...]int8{
	0, 2, 1, 2, 1, 1, 0, 4, 3, 1,
	2, 2, 1, 2, 1, 1, 1, 1, 1, 1,
	0, 2, 1, 0,
}

var yyChk = [...]int16{
	-1000, -14, -1, -3, -2, -4, 6, 4, -4, 8,
	-5, -6, -7, -8, -9, 6, 7, 9, 10, -13,
	-8, -12, 14, -11, -10, 11, 12, 13, -6, 6,
}

var yyDef = [...]int8{
	6, -2, 0, 2, 5, 4, 0, 1, 3, 0,
	0, 9, 23, 12, 20, 14, 15, 7, 0, 10,
	11, 22, 0, 13, 19, 16, 17, 18, 8, 21,
}

var yyTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 14, 3, 3, 3, 3,
	3, 3, 11, 12, 3, 3, 9, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 8, 3, 13, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 10,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = int(yyPact[yystate])
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:100
		{
			yyVAL.prod = &grammar.Syntax{
				Productions: cast[[]*grammar.Production](yyDollar[1].prod),
			}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:108
		{
			yyVAL.prod = append(
				cast[[]*grammar.Production](yyDollar[1].prod),
				cast[*grammar.Production](yyDollar[2].prod),
			)
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:115
		{
			yyVAL.prod = []*grammar.Production{
				cast[*grammar.Production](yyDollar[1].prod),
			}
		}
	case 6:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:123
		{
			yyVAL.prod = nil
		}
	case 7:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:128
		{
			yyVAL.prod = &grammar.Production{
				Name:  yyDollar[1].tok.Str,
				Terms: cast[[]*grammar.Term](yyDollar[2].tok),
			}
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:137
		{
			yyVAL.prod = append(
				cast[[]*grammar.Term](yyDollar[1].prod),
				cast[*grammar.Term](yyDollar[2].tok),
			)
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:144
		{
			yyVAL.prod = []*grammar.Term{
				cast[*grammar.Term](yyDollar[1].prod),
			}
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:152
		{
			yyVAL.prod = &grammar.Term{
				Factors: cast[[]*grammar.Factor](yyDollar[1].prod),
				Label:   cast[*grammar.Label](yyDollar[2].prod),
			}
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:161
		{
			yyVAL.prod = append(
				cast[[]*grammar.Factor](yyDollar[1].prod),
				cast[*grammar.Factor](yyDollar[2].prod),
			)
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:168
		{
			yyVAL.prod = []*grammar.Factor{
				cast[*grammar.Factor](yyDollar[1].prod),
			}
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:176
		{
			factor := cast[*grammar.Factor](yyDollar[1].prod)
			qualifier := cast[grammar.Qualifier](yyDollar[2].prod)
			factor.Qualifier = qualifier
			yyVAL.prod = factor
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:184
		{
			yyVAL.prod = &grammar.Factor{Name: yyDollar[1].tok.Str}
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:185
		{
			yyVAL.prod = &grammar.Factor{Literal: yyDollar[1].tok.Str}
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:188
		{
			yyVAL.prod = grammar.ZeroOrMore
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:189
		{
			yyVAL.prod = grammar.OneOrMore
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:190
		{
			yyVAL.prod = grammar.ZeroOrOne
		}
	case 20:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:194
		{
			yyVAL.prod = grammar.NoQualifier
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:198
		{
			yyVAL.prod = &grammar.Label{
				Label: yyDollar[2].tok.Str,
			}
		}
	case 23:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:206
		{
			yyVAL.prod = nil
		}
	}
	goto yystack /* stack new state and value */
}
