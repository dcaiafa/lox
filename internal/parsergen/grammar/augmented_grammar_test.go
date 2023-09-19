package grammar

import (
	"testing"

	"github.com/dcaiafa/lox/internal/errlogger"
	"github.com/dcaiafa/lox/internal/util/set"
)

func TestFirst1(t *testing.T) {
	// Dragon book section 4.11.
	//	E  -> TE'
	//	E' -> +TE' | ε
	//	T  -> FT'
	//	T' -> *FT' | ε
	//	F  -> ( E ) | id
	sg := &Grammar{
		Terminals: []*Terminal{
			{Name: "id"},
			{Name: "+"},
			{Name: "*"},
			{Name: "("},
			{Name: ")"},
		},
		Rules: []*Rule{
			{
				Name: "E",
				Prods: []*Prod{
					{Terms: []*Term{{Name: "T"}, {Name: "E'"}}},
				},
			},
			{
				Name: "E'",
				Prods: []*Prod{
					{Terms: []*Term{{Name: "+"}, {Name: "T"}, {Name: "E'"}}},
					{Terms: []*Term{}},
				},
			},
			{
				Name: "T",
				Prods: []*Prod{
					{Terms: []*Term{{Name: "F"}, {Name: "T'"}}},
				},
			},
			{
				Name: "T'",
				Prods: []*Prod{
					{Terms: []*Term{{Name: "*"}, {Name: "F"}, {Name: "T'"}}},
					{Terms: []*Term{}},
				},
			},
			{
				Name: "F",
				Prods: []*Prod{
					{Terms: []*Term{{Name: "("}, {Name: "E"}, {Name: ")"}}},
					{Terms: []*Term{{Name: "id"}}},
				},
			},
		},
	}

	errs := errlogger.New()

	g := sg.ToAugmentedGrammar(errs)
	if errs.HasError() {
		t.Fatalf("ToAugmentedGrammar failed")
	}

	assertTerminalSetEq := func(t *testing.T, symSet *set.Set[*Terminal], symNames ...string) {
		t.Helper()

		expected := new(set.Set[string])
		expected.AddSlice(symNames)
		actual := new(set.Set[string])
		symSet.ForEach(func(terminal *Terminal) {
			actual.Add(terminal.Name)
		})

		if !actual.Equal(expected) {
			t.Log("Expected: ", expected.Elements())
			t.Log("Actual: ", actual.Elements())
			t.Fatalf("Terminal set did not match expectation")
		}
	}

	sym := func(name string) Symbol {
		return g.GetSymbol(name)
	}

	assertTerminalSetEq(t, g.first(sym("E")), "(", "id")
	assertTerminalSetEq(t, g.first(sym("E'")), "+", "ε")
	assertTerminalSetEq(t, g.first(sym("T'")), "*", "ε")
	assertTerminalSetEq(t, g.First([]Symbol{sym("E'"), sym("E")}), "+", "(", "id")
	assertTerminalSetEq(t, g.First([]Symbol{sym("E'"), sym("T'")}), "+", "*", "ε")
	assertTerminalSetEq(t, g.First([]Symbol{sym("E'"), sym("id")}), "+", "id")
}

func TestFirst2(t *testing.T) {
	// X -> Y Z '*'
	// Y -> '+' | ε
	// Z -> '-' | ε
	sg := &Grammar{
		Terminals: []*Terminal{
			{Name: "+"},
			{Name: "-"},
			{Name: "*"},
		},
		Rules: []*Rule{
			{
				Name: "X",
				Prods: []*Prod{
					{Terms: []*Term{{Name: "Y"}, {Name: "Z"}, {Name: "*"}}},
				},
			},
			{
				Name: "Y",
				Prods: []*Prod{
					{Terms: []*Term{{Name: "+"}}},
					{Terms: []*Term{}},
				},
			},
			{
				Name: "Z",
				Prods: []*Prod{
					{Terms: []*Term{{Name: "-"}}},
					{Terms: []*Term{}},
				},
			},
		},
	}

	errs := errlogger.New()

	g := sg.ToAugmentedGrammar(errs)
	if errs.HasError() {
		t.Fatalf("ToAugmentedGrammar failed")
	}

	assertTerminalSetEq := func(t *testing.T, symSet *set.Set[*Terminal], symNames ...string) {
		t.Helper()

		expected := new(set.Set[string])
		expected.AddSlice(symNames)
		actual := new(set.Set[string])
		symSet.ForEach(func(terminal *Terminal) {
			actual.Add(terminal.Name)
		})

		if !actual.Equal(expected) {
			t.Log("Expected: ", expected.Elements())
			t.Log("Actual: ", actual.Elements())
			t.Fatalf("Terminal set did not match expectation")
		}
	}

	sym := func(name string) Symbol {
		return g.GetSymbol(name)
	}

	assertTerminalSetEq(t, g.first(sym("X")), "+", "-", "*", "ε")
}
