# Lexer Reference

The purpose of a lexer is to break the input into a series of tokens, providing
the parser with the fundamental building blocks it can expect. In a lexer
section, the most fundamental declaration is the token rule. However, a lexer
can also define other constructs, such as fragments, macros, and modes, which
help model more complex grammars and facilitate more sophisticated tokenization
strategies.

## Token Rule

A token declaration follows this format:

```
NAME = <expression> <action>*
```

### Token Name

The token name uniquely identifies the token, and allows it to be referenced in
the parser grammar.

* **Must** be all uppercase.
* **Must** start with a letter.
* **May** contain letters, numbers, and underscores after the first character.
* **Must not** end with an underscore.
* **Must not** contain consecutive underscores.
* **Must** be unique.
* **Must not** be one of the reserved names: `EOF`, `ERROR`.

### Token Expression

The token expression determines how a token
  
| Expression | Description |
| ---------- | ----------- |
| 'literal'  | Match a sequence of characters (e.g. 'func', '!=', ','). Special characters must be [escaped](#literal-escaping-rules).
| [*char_class*] | Match one of the characters in the set. `x-y` specifies a range of characters (e.g. `[A-Ca-c]` is equivalent to `[ABCabc]`). [Escaped](#literal-escaping-rules) characters are also allowed (e.g. `[a\-z]` will match `a`, `-` or `z`).
| `~`&nbsp;[*char_class*] | Is like the character class, but it matches the any characters **not** in the set.
| *expr*` `*expr* | Match one expression followed by another (e.g. `'//' ~[\n]*`).
| *expr*`\|`*expr* | Match either expression (e.g. `[1-9][0-9]* \| 'nan'`).
| `(`*expr*`)` | Group an expression (e.g. `('foo' \| 'bar')*`).
| *expr*`?` | Optionally match expression (e.g. `[1-9][0-9]*('.'[0-9]+)?` specifies a number with an optional fractional part).
| *expr*`*` | Match the expression zero or more times (e.g. `[1-9][0-9]*` matches sequences like `1` and `10`).
| *expr*`+` | Match the expression one or more times (e.g. `[1-9][0-9]+` matches `22`, `109`, but not `1`).

#### Literal Escaping Rules:
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

