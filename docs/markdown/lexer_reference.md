# Lexer Reference

The purpose of a lexer is to break the input into tokens which can then be used
by a parser to build complex syntatic structures.

In a lexer section, the most common declaration is the token rule. However, a
lexer can also define other lexical elements, such as fragments, macros, and
modes, which help model more complex grammars and facilitate more sophisticated
tokenization strategies.

## Declarations

Lexer declarations are contained in a lexer section signaled by the `@lexer`
keyword.

```lox
@lexer

// Lexer declarations
```
### Line Continuation

Declarations are terminated by an end-of-line. The line-continuation backslash
`\` must be used to split a declaration into multiple lines. The exception is
when the first token in the next line is the vertical bar `|`. In this case, the
line-continuation is implicit. 

For example:

```lox
@macro INTEGER = DIGIT | \
                 ONE_NINE DIGIT+
```
Is equivalent to:
```lox
@macro INTEGER = DIGIT
               | ONE_NINE DIGIT+
```
It is idiomatic to use the latter form, and only use the `\` when a declaration
cannot be split at a `|`.

### Declaration Order

The order of lexer declarations determines which token is emitted when a
sequence of characters matches more than one lexical expression. In cases where
multiple expressions could match the same input, the lexer will emit the token
corresponding to the first matching expression encountered in the declaration
order.

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

Token rules are the fundamental lexer building block. They define a lexical
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

Tokens carry the [action](#lexical-actions) `@emit(NAME)` implicitly.

### Fragments

Fragments are similar to tokens in both syntax and semantics, but with some key
differences. Unlike tokens, fragments are unnamed and do not have a default
action. If a fragment does not specify an action, all characters recognized by
the fragment remain in the lexer's accumulator. This behavior is generally
useful within a mode, where fragments can be used to break down a single token
into multiple expressions, allowing for more granular control over the
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

A mode is a group of lexical expressions that allows you to switch between
different sets of tokens or fragments depending on the context. Tokens or
fragments declared outside of any specific mode belong to the default mode. The
`@push_mode` and `@pop_mode` [actions](#lexical-actions) are used to switch
between modes during lexing.

Example:
```lox
@lexer
PLUS   = '+'
MINUS  = '-'
OPAREN = '(' @push_mode(Alt)

@mode Alt {
    DASH   = '-'
    CPAREN = ')' @pop_mode
}
```
Given the input sequence `+-(--)-+`, the lexer would emit the following tokens:
`PLUS`, `MINUS`, `OPAREN`, `DASH`, `DASH`, `CPAREN`, `MINUS`, `PLUS`. 

**Explanation:**
* After emitting `OPAREN`, the lexer switches to `Alt` mode, where `-` is now
  recognized as `DASH` instead of `MINUS`.
* In `Alt` mode, encountering a `+` would result in an error since `+` is not
  defined in Alt.
* When `CPAREN` is encountered, the lexer pops `Alt` mode, returning to the
  default mode, where `-` is once again recognized as `MINUS` and `+` as `PLUS`.

The lexer maintains a stack of modes, starting in the default mode. It is
possible to push the default mode onto the stack by using `@push_mode()` without
specifying a mode name.

## Common Elements

### Lexical Expressions

Lexical expressions are used by tokens, fragments and macros to determine how
the lexer will match sequences of characters.
  
| Expression | Description |
| ---------- | ----------- |
| 'literal'  | Match a sequence of characters (e.g. 'func', '!=', ','). Special characters must be [escaped](#literal-escaping-rules).
| `.` | Match any character.
| [*char_class*] | Match one of the characters in the set. x-y specifies a range of characters (e.g. [A-Ca-c] is equivalent to [ABCabc]). [Escaped](#literal-escaping-rules) characters are also allowed (e.g. [a\-z] will match a, `-` or z).
| ~[*char_class*] | Like [*char_class*], but it matches characters **not** in the set.
| *cc* - *cc* | Matches the difference between two character classes (e.g. [A-Z] - [IJK] matches any character between A and Z that is not I, J or K).
| *expr* *expr* | Match one expression followed by another (e.g. '//' ~[\n]*).
| *expr* \| *expr* | Match either expression (e.g. [1-9][0-9]* \| 'pi').
| (*expr*) | Group an expression (e.g. ('foo' \| 'bar')*).
| *expr* ? | Optionally match expression (e.g. [1-9][0-9]*('.'[0-9]+)? specifies a number with an optional fractional part).
| *expr* * | Match the expression zero or more times (e.g. [1-9][0-9]* matches sequences like 1 and 123).
| *expr* + | Match the expression one or more times (e.g. [1-9][0-9]+ matches 22, 109, but not 1).
| *expr* *? | Like `*`, but non-greedy. 
| *expr* +? | Like `+`, but non-greedy. 

{.notice}
The non-greedy cardinalities `*?` and `+?` will consume the least amount of
input required. As such they make no sense in the end of term expressions. The
expression `[0-9]+?` by itself will never match more than one digit. On the
other hand, the C comment lexer expression `'/*' .*? '*/'` will not work
correctly without the `?`. This is because `.*` will also match `*/`. 

### Lexical Names

Names declared in a lexer section must conform to the following rules:

* **Must** be all uppercase.
* **Must** start with a letter.
* **May** contain letters, numbers, and underscores after the first character.
* **Must** be unique.
* **Must not** end with an underscore.
* **Must not** contain consecutive underscores.
* **Must not** be one of the reserved names: EOF, ERROR.

### Lexical Actions

A lexical action determines the action executed by the Lexer when it matches a
token or a fragment.

| Keyword | Description |
| ------- | ----------- |
| @emit(TOKEN) | Emit the token referenced by the given name. Only valid in fragments.
| @discard | Discard all accumulated characters (e.g. the rule @frag [ \n\r\t]+ @discard will discard whitespaces)
| @push_mode(MODE?) | Push the current mode onto the stack and enter the mode with name MODE. If MODE is not provided, it will enter the default mode.
| @pop_mode | Pop the name on the top of the mode stack and make it the current mode.

### Literal Escaping Rules

| Escaped Sequence | Actual Character |
| ---------------- | ---------------- |
| \\n | New line (a.k.a. carriage return) UTF-8: 0x0D. |
| \\r | Line feed UTF-8: 0x0A. |
| \\t | Horizontal tab UTF-8: 0x09. |
| \\' | The single quote character ' (only valid in token literal). |
| \\- | The short dash character - (only valid in character class). |
| \\xXX | Single byte unicode character in hexadecial (e.g. \\x2A is *). |
| \\uXXXX | Double byte unicode character in hexadecimal (e.g. \\u4E16 is 世). |
| \\UXXXXXXXX | Four byte unicode character (e.g. \\UF0938583 is 𓅃). |

## Examples

### Keywords and Identifiers

```lox
// Keywords
WHILE    = 'while'
CONTINUE = 'continue'
IF       = 'if'
ELSE     = 'else'

// Identifier
ID = [A-Za-z_] [A-Za-z0-9_]*
```
Keywords are often specific cases of the identifier lexical expression. If this
is the case in your grammar, make sure that you declare the keywords before the
identifier. Otherwise, the identifier will supersede all keywords, causing the
lexer to recognize them as identifiers instead of their respective keyword
tokens.

### Number Literals
```lox
@macro ONE_NINE = [1-9]
@macro DIGIT    = '0' | ONE_NINE
@macro INTEGER  = DIGIT
                | ONE_NINE DIGIT+
                | '-' DIGIT
                | '-' ONE_NINE DIGIT+
@macro FRACTION = '.' DIGIT+
@macro EXPONENT = [eE] [+-]? ONE_NINE DIGIT*
NUMBER = INTEGER FRACTION? EXPONENT?
```
This example defines a `NUMBER` literal that includes an integer part and
optional fraction and exponent parts. Macros are used to break the token
declaration into smaller, more readable components.

### Line Continuation

```lox
NL = '\n'
@frag '\\' [ \r\n\t]* '\n' @discard
```
In languages where newlines are used for statement termination, you might need a
mechanism to allow statements to span multiple lines. The example above
demonstrates how to handle this using a backslash (`\`) as a line continuation
character.

**Explanation:**
* **NL Token**: The `NL` token represents a newline character (`\n`).
* **Line Continuation Fragment**: The fragment `@frag '\\' [ \r\n\t]* '\n' @discard`
  handles the line continuation. When a backslash (`\`) appears at the end of a
  line, followed by optional whitespace, the newline character is discarded,
  preventing it from being emitted as an `NL` token.

This setup allows the lexer to treat a backslash followed by a newline as a
continuation of the same statement, effectively ignoring the newline.

### String interpolation

```lox
NUM = [0-9]+
PLUS = '+'

STR_BEGIN = '"' @push_mode(String)
@mode String {
  STR_END = '"' @pop_mode
  CHAR_SEQ = (~["\n{}\\] | '\\' ["nrt{}\\])*
  OCURLY = '{' @push_mode() @emit(OCURLY)
}
CCURLY = '}' @pop_mode
```
This example demonstrates how to implement string interpolation using modes in a
lexer. The grammar allows for embedded expressions within strings, as seen in
the input  `"1 + 2 = {1+2}"`, which would be parsed as `STR_BEGIN`,
`CHAR_SEQ(1 + 2 =)`, `OCURLY`, `NUM(1)`, `PLUS`, `NUM(2)`, `CCURLY`.

**Explanation:**

* **NUM and PLUS**: These tokens represent numbers and the plus sign within the
  interpolated expression.
* **STR_BEGIN and STR_END**: These tokens mark the beginning and end of a
  string. `STR_BEGIN` pushes the lexer into `String` mode, while `STR_END` pops
  the mode, returning to the previous state.
* **CHAR_SEQ**: This fragment matches sequences of characters within the string,
  while also handling escaped characters such as `\n`, `\t`, `{`, and `}`.
* **OCURLY**: The opening curly brace (`{`) within the String mode triggers a
  mode push to the default mode (indicated by the lack of mode parameter),
  allowing the lexer to parse the embedded expression. It also emits the
  `OCURLY` token.
* **CCURLY**: The closing curly brace (`}`) pops the current mode, signaling the
  end of the interpolated expression.
