# Parser Reference

A parser's purpose is to transform a sequence of tokens into high-level
programming constructs that represent the syntactical structure of the input
according to a defined grammar.

For example, the expression `1 + 2 * 3` can be transformed into an Abstract
Syntax Tree (AST) as shown below:

```
        Binary
      Expression
          +
        /   \
       /     \
  Literal    Binary
     1       Expression
                 *
               /   \
              /     \
          Literal  Literal
            2        3
```

## Parser Section

Parser declarations are contained in a parser section which is signaled using
the keyword `@parser`.

```lox
@parser

// Parser declarations

```

## Rules

Parser rules are the fundamental building blocks of the parser. They define
transformations of terminals (tokens) and non-terminals (other rules) and can
themselves be referenced as terms in other rule productions, including
recursively. The parser rule syntax resembles the
[extended-Backus-Naur form](https://en.wikipedia.org/wiki/Extended_Backus%E2%80%93Naur_form)
(EBNF).

The parser rule follows this format:
```
rule = prod1 | prod2 | ... | prodN
```
Where `rule` is the rule name. Rule names must be valid [Go
identifiers](https://go.dev/ref/spec#Identifiers) with the following additional
restrictions:
* **Must not** start with an underscore (`_`).
* **Must not** contain consecutive underscores.


