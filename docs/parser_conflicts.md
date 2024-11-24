# Conflicts

If you process the following grammar snippet with `lox`, it will fail with the
error: `grammar has conflicts`:


```lox
expr = expr '+' expr
     | expr '*' expr
     | NUMBER
```

Running `lox` with the `-report` flag provides the missing details:

```
I5:
  expr = expr .PLUS expr, EOF
  expr = expr .PLUS expr, PLUS
  expr = expr .PLUS expr, MUL
  expr = expr PLUS expr., EOF
  expr = expr PLUS expr., PLUS
  expr = expr PLUS expr., MUL
  expr = expr .MUL expr, EOF
  expr = expr .MUL expr, PLUS
  expr = expr .MUL expr, MUL
    on EOF reduce expr
    on MUL reduce expr <== CONFLICT
    on MUL shift I3 <== CONFLICT
    on PLUS shift I4 <== CONFLICT
    on PLUS reduce expr <== CONFLICT
```
The excerpt above represents a single state in the parser state machine. In this
state, the parser has just parsed a `expr PLUS expr` (e.g. `2 + 2`). Say the
next token is a `MUL`, what should it do? It could reduce `expr PLUS expr` into
`expr` or it could shift `MUL`. The grammar does not make it clear which
action to take, so it is ambiguous.

One way to resolve this is by refactoring the grammar:

```lox
expr = expr '+' term
     | term

term = term '*' factor
     | factor

factor = number

number = NUMBER
       | '-' NUMBER
```

Some conflicts can only be resolved by refactoring but a large class of
conflicts can be resolved using 
[precedence qualifiers](./parser_reference#precedence_qualifiers):

```lox
expr = expr '+' expr  @left(1)
     | expr '*' expr  @left(2)
     | NUMBER
```

The `@left` qualifiers in the grammar tells `lox` that if it encounters a
conflict between `expr '+' expr` and `expr '*' expr`, then the latter should
take precendence over the former. 
