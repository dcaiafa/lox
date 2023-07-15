package parser2

import "testing"

const parserLox = `
@lexer

@custom 
  ID
  LITERAL
  LABEL

  ZERO_OR_MANY
  ONE_OR_MANY
  ZERO_OR_ONE

  DEFINE
  OR
  SEMICOLON

  PARSER
  LEXER
  CUSTOM ;

@parser

Spec     = Section+ ;
Section  = Parser | Lexer;
Parser   = PARSER Pdecl* ;
Pdecl    = Prule ;
Prule    = ID DEFINE Pprods SEMICOLON ;
Pprods   = Pprods OR Pprod
         | Pprod ;
Pprod    = Pterm+ Label? ;
Pterm    = ID Pcard? ;
Pcard    = ZERO_OR_MANY | ONE_OR_MANY | ZERO_OR_ONE ; 
Label    = LABEL ;

Lexer    = LEXER Ldecl* ;
Ldecl    = Lcustom ;
Lcustom  = CUSTOM ID+ SEMICOLON ;
`

func TestParser(t *testing.T) {
	Parse("foo.lox", []byte(parserLox))
}
