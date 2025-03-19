# Getting started

Lox is a parser/lexer generator for the Go language. It consists of the `lox`
tool which analyzes a project's grammar and Go sources, and generates a
dependency-free parser in Go.

## Installing Lox

Download the [latest](https://github.com/dcaiafa/lox/releases/latest) version of
`lox` matching your platform. Unpack it and copy the `lox` binary to a directory
in your `PATH`.

{.notice}
Because `lox` analyzes Go sources, it is very important to update it often,
especially when upgrading your Go tooling.

## Experiment with an example

The best way to start learning Lox is to play with an
[example](https://github.com/dcaiafa/lox/tree/main/examples). In this section,
we are going to use the
[calculator](https://github.com/dcaiafa/lox/tree/main/examples/calc).

First, try running the example by going to the `calc` directory in your terminal
and running it using `go run .`. Enter an expression like `2 * (1+3)` and you
should see the result.

Next, let's add built-in function. First, we need to 
