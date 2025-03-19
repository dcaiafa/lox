# Lox

Lox is a parser/lexer generator for Go.

```lox
@lexer

NUM = ('0' | [1-9][0-9]*) ('.' [0-9]+)? 
ADD = '+'
SUB = '-'
MUL = '*'
DIV = '/'
O_PAREN = '('
C_PAREN = ')'

@frag ' '+  @discard  // Discard whitespace. 

@parser

@start S = expr

expr = expr '+' expr  @left(1)
     | expr '-' expr  @left(1)
     | expr '*' expr  @left(2)
     | expr '/' expr  @left(2)
     | '(' expr ')'
     | num

num = NUM
    | '-' NUM
```

