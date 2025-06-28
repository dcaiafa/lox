# Lox

Lox is a lexer and parser generator for Go.

## Documentation

For comprehensive documentation, examples, and tutorials, visit **[dcaiafa.github.io/lox](https://dcaiafa.github.io/lox/)**.

## Quick Start

```bash
# Install
# Download the latest release from https://github.com/dcaiafa/lox/releases/latest

# Generate parser from grammar
lox .
```

## Example Grammar

```
@lexer
NUM = [0-9]+
ADD = '+'

@parser
@start expr
expr = expr ADD expr  @left(1)
     | NUM
```

## Features

- Self-hosted parser generator
- LR(1) parser generation
- DFA-based lexer generation
- Type-safe Go code generation
- Comprehensive error reporting

## License

MIT
