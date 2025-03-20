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

## Try an example

The best way to start learning Lox is to play with an
[example](https://github.com/dcaiafa/lox/tree/main/examples). In this section,
we are going to use the
[calculator](https://github.com/dcaiafa/lox/tree/main/examples/calc).

First, try running the example by going to the `calc` directory in your terminal
and running it using `go run .`. Enter an expression like `2 * (1+3)` and you
should see the result.

Next, let's extend the project to support built-in functions so you can evaluate
something like `pow(2, ) / 2`. We need an identifier token that will match
function names (`sqrt`, `pow`, etc.). We will also need a comma token to use as
the parameter separator.

Open `calc.lox` in an editor, then add the following to the end of the `@lexer`
section, right after the definition for `NUM`:

```lox
COMMA = ','
ID = [A-Za-z][A-Za-z0-9]*
```

Now add a parser rule for parsing a function call. Add the following to the end
of the `@parser` section (at the end of the file):

```lox
func_call = ID '(' @list(expr, ',')? ')'
```
The `ID` `'('` `')'` part should look pretty straightforward. The `@list(expr,
',')` term will match one or more `expr` separated by `,`. The `?` qualifier
makes it optional so you can have functions that take no parameters (e.g.
`pi()`). Notice that you can reference tokens in your grammar using their name
(e.g. `COMMA`) or their literal value (e.g. `','`). The latter is usually easier
to read, but is only allowed when the token consists of a simple literal.

Now, change the definition of `expr` to include `func_call`:

```lox
expr = expr '+' expr  @left(1)
     | expr '-' expr  @left(1)
     | expr '*' expr  @left(2)
     | expr '/' expr  @left(2)
     | expr '%' expr  @left(2)
     | expr '^' expr  @right(3)
     | '(' expr ')'
     | num
     | func_call
```

We are done with the grammar changes. Before moving on to the Go changes, run
`lox .` from the `calc` directory. You should see an error like the following:

```
 $ lox .
calc.lox:35:1: rule missing action method: func_call
Error: errors ocurred
```
You got this error because Lox expects every production from every rule to have
a corresponding action.

Let's add one for `func_call`. Open `parser.go` in your IDE, then add the
following to the end of the file:

```go
func (p *calcParser) on_func_call(
	id Token, _ Token, args []float64, _ Token,
) float64 {
	name := string(id.Str)
	switch name {
	case "pow":
		if len(args) != 2 {
			p.errLogger.Errorf(id.Pos, "pow takes 2 arguments")
			return 0
		}
		return math.Pow(args[0], args[1])

	default:
		p.errLogger.Errorf(id.Pos, "invalid function %v", name)
		return 0
	}
}
```
Try running `lox .` again. You should get no errors. And you should be able to
run the project again and get a result for an expression like `pow(2, 8) - 1`.

Hopefully the implementation of `on_func_call` should look pretty
straightforward. But there is a lot of magic happening in the method signature.
Here are a few useful things you should know:
* The `on_` prefix is how Lox knows that this is an action method. It should be
  followed by the name of the rule. It might include a `__*` suffix which will
  be ignored by Lox, but that is sometimes necessary when you need multiple
  action methods for a single rule. For example: `on_num` and `on_num__minus`
  both reference the `num` rule; the `__minus` suffix is arbitrary.
* The number and type of parameters must match the production being targeted.
  The single `func_call` production has 4 terms, so `on_func_call` must have 4
  parameters.
* The name of the parameters do not matter. Since we don't care about the
  value of the parenthesis token terms, we named them `_`.
* The type of the parameters is very important, and must match the term action
  type. For example: `on_func_call` returns `float64`, so an action that has a
  parameter for a `func_call` term must use a type that `float64` is assignable
  to.
* Tokens don't have actions, and their action type will always be `Token`.
* The action type of `@list(X, Y)` is a slice of the action type of `X`. In this
  case, it is a slice of the action type of `expr` which is `float64`. Hence the
  action type of `@list(expr, ',')` is `[]float64`.
* Notice that I said that the term action type must be **assignable** to the
  corresponding parameter type. So you could use `any` as the parameter type.
  This is often a bad idea as you would not be letting Lox do the type-checking
  for you (and you would need to cast the values yourself).
* Also, notice that you did not need to provide an action method for the
  `func_call` production that you added to `expr`. This is because the
  production will match to `on_expr__num` which will now handle `num` and
  `func_call` (and it should probably be renamed to reflect that). The match
  occurs because both `num` and `func_call` have a single term whose action type
  is `float64`. If `func_call`'s action type was not assignable to `float64`,
  you would need to provide a dedicated action for it.
