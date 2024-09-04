# Lexer Reference

The purpose of a lexer is to break the input into a series of tokens, providing
the parser with the fundamental building blocks it can expect. In a lexer
section, the most common declaration is the token rule. However, a lexer
can also define other constructs, such as fragments, macros, and modes, which
help model more complex grammars and facilitate more sophisticated tokenization
strategies.

## Declarations

Lexer declarations are contained in a lexer section which is signaled using the
keyword `@lexer`.

```lox
@lexer

// Lexer declarations

@parser

// Parser declarations
```

### Declaration Order

The order of lexer declarations is crucial because it determines which token is
emitted when a sequence of characters matches more than one lexical expression.
In cases where multiple expressions could match the same input, the lexer will
emit the token corresponding to the first matching expression encountered in the
declaration order.

For example, given the following grammar:
```lox
@lexer

TOO_GOOD = '2good'
NUMBER   = [0-9]+
ID       = [a-z0-9]+
```
* The input `2` emits `NUMBER`.
* The input `2x` emits `ID`.
* The input `2good` emits `TOO_GOOD`.

However, if `TOO_GOOD` were defined after `ID`, then `TOO_GOOD` would never be emitted
because `ID` would match `2good` first.

### Tokens

Tokens are the fundamental lexer building block. They define a lexical
expression that once recognized by the lexer state machine causes that token to
be emitted.

A token rule follows this format:

```text
NAME = <expression> <action>*
```
Where:
* `NAME` is a [lexical name](#lexical-names).
* `<expression>` is a [lexical expression](#lexical-expressions).
* `<action>` is a [lexical-action](#lexical-actions).

Tokens carry the action `@emit(NAME)` implicitly.

### Fragments

Fragments are similar to tokens in both syntax and semantics, but with some key
differences. Unlike tokens, fragments are unnamed and do not have a default
action. If a fragment does not specify an action, all characters recognized by
the fragment remain in the lexer's accumulator. This behavior is generally
useful within a sub-mode, where fragments can be used to break down a single
token into multiple expressions, allowing for more granular control over the
tokenization process.

In the default mode, fragments typically specify the `@discard` action, such as
when discarding whitespace or other insignificant characters that should not be
part of the token stream.

A fragment rule follows this format:

```text
@frag <expression> <action>*
```
Where:
* `<expression>` is a [lexical expression](#lexical-expressions).
* `<action>` is a [lexical-action](#lexical-actions).


### Macros

Macros, like tokens, define a lexical expression. However, unlike tokens, macros
do not directly influence the state machine and cannot be referenced by the
parser. Instead, macros serve as reusable components to simplify and streamline
lexer definitions.

For example, it is common to define a macro such as `@macro DIGIT=[0-9]` to
represent a digit. This macro can then be used within various token definitions,
reducing the need to repeat the expression `[0-9]` multiple times throughout the
lexer, thereby improving readability and maintainability.

A macro rule follows this format:

```text
@macro NAME = <expression>
```
Where:
* `NAME` is a [lexical name](#lexical-names).
* `<expression>` is a [lexical expression](#lexical-expressions).

### Modes

## Common Elements

### Lexical Expressions

Lexical expressions are used by tokens, fragments and macros to determine how
the lexer will match sequences of characters.
  
| Expression | Description |
| ---------- | ----------- |
| 'literal'  | Match a sequence of characters (e.g. 'func', '!=', ','). Special characters must be [escaped](#literal-escaping-rules).
| [*char_class*] | Match one of the characters in the set. `x-y` specifies a range of characters (e.g. `[A-Ca-c]` is equivalent to `[ABCabc]`). [Escaped](#literal-escaping-rules) characters are also allowed (e.g. `[a\-z]` will match `a`, `-` or `z`).
| `~`&nbsp;[*char_class*] | Is like the character class, but it matches the any characters **not** in the set.
| *expr*` `*expr* | Match one expression followed by another (e.g. `'//' ~[\n]*`).
| *expr*`\|`*expr* | Match either expression (e.g. `[1-9][0-9]* \| 'pi'`).
| `(`*expr*`)` | Group an expression (e.g. `('foo' \| 'bar')*`).
| *expr*`?` | Optionally match expression (e.g. `[1-9][0-9]*('.'[0-9]+)?` specifies a number with an optional fractional part).
| *expr*`*` | Match the expression zero or more times (e.g. `[1-9][0-9]*` matches sequences like `1` and `10`).
| *expr*`+` | Match the expression one or more times (e.g. `[1-9][0-9]+` matches `22`, `109`, but not `1`).

### Lexical Names

Names declared in a lexer section must conform to the following rules:

* **Must** be all uppercase.
* **Must** start with a letter.
* **May** contain letters, numbers, and underscores after the first character.
* **Must not** end with an underscore.
* **Must not** contain consecutive underscores.
* **Must** be unique.

### Lexical Actions

A lexical action determines the action executed by the Lexer when it matches a
token or a fragment.

| Keyword | Description |
| ------- | ----------- |
| `@emit(TOKEN)` | Emit the token with the provided name. Only valid in fragments.
| `@discard` | Discard all accumulated characters (e.g. the rule `@frag [ \n\r\t]+ @discard` will discard whitespaces)
| `@push_mode(MODE?)` | Push the current mode to the stack and enter the mode with name `MODE`. If `MODE` is not provided, it will enter the default mode.
| `@pop_mode` | Pop the name on the top of the mode stack and make it the current mode.

### Literal Escaping Rules:

| Escaped Sequence | Actual Character |
| ---------------- | ---------------- |
| `\n` | New line (a.k.a. carriage return) UTF-8: 0x0D. |
| `\r` | Line feed UTF-8: 0x0A. |
| `\t` | Horizontal tab UTF-8: 0x09. |
| `\'` | The single quote character `'` (only valid in token literal). |
| `\-` | The short dash character `-` (only valid in character class). |
| `\xXX` | Single byte unicode character in hexadecial (e.g. `\x2A` is `*`). |
| `\uXXXX` | Double byte unicode character in hexadecimal (e.g. `\u4E16` is `世`). |
| `\UXXXXXXXX` | Four byte unicode character (e.g. `\UF0938583` is `𓅃`). |
* **Must not** be one of the reserved names: `EOF`, `ERROR`.
