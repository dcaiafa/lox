package ast_test

import (
	gotoken "go/token"
	"strings"
	"testing"

	"github.com/dcaiafa/lox/internal/errlogger"
	"github.com/dcaiafa/lox/internal/lexergen/ast"
	"github.com/dcaiafa/lox/internal/lexergen/parser"
)

func requireEqualStr(t *testing.T, actual, expected string) {
	t.Helper()
	actual = strings.TrimSpace(actual)
	expected = strings.TrimSpace(expected)
	if actual != expected {
		t.Fatalf("not equal:\nexpected:\n%v\nactual:\n%v", expected, actual)
	}
}

func parse(t *testing.T, input string) (*ast.Spec, *ast.Context) {
	fset := gotoken.NewFileSet()
	errs := errlogger.New()
	file := fset.AddFile("input.lox", -1, len(input))
	unit := parser.Parse(file, []byte(input), errs)
	if errs.HasError() {
		t.Fatalf("Failed to parse")
	}
	spec := new(ast.Spec)
	spec.Units = []*ast.Unit{unit}
	ctx := ast.NewContext(fset, errs)
	return spec, ctx
}

func parseAndAnalyze(t *testing.T, input string) (*ast.Spec, *ast.Context) {
	spec, ctx := parse(t, input)
	if !ctx.Analyze(spec) {
		t.Fatalf("Failed to analyze")
	}
	return spec, ctx
}