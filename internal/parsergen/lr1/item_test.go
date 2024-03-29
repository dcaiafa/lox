package lr1

import "testing"

func TestItem(t *testing.T) {
	// S' = xs EOF
	// xs = xs x | x
	// x  = '+'
	g := NewGrammar()

	var (
		tAdd = g.AddTerminal("+")

		rXS = g.AddRule("XS")
		rX  = g.AddRule("X")
	)

	g.SetStart(rXS)
	g.AddProd(rXS, rXS, rX)
	g.AddProd(rXS, rX)
	g.AddProd(rX, tAdd)

	t.Run("1", func(t *testing.T) {
		i := Item{Prod: rXS.Prods[0].Index, Dot: 0, Lookahead: tAdd.Index}
		if i.IsKernel() {
			t.Fatalf("[%v] can't be a kernel", i.ToString(g))
		}
		const expected = "XS = .XS X, +"
		result := i.ToString(g)
		if result != expected {
			t.Fatalf("Expected: [%v] Actual: [%v]", expected, result)
		}
	})
	t.Run("2", func(t *testing.T) {
		i := Item{Prod: rXS.Prods[0].Index, Dot: 1, Lookahead: tAdd.Index}
		if !i.IsKernel() {
			t.Fatalf("[%v] should be a kernel", i.ToString(g))
		}
		const expected = "XS = XS .X, +"
		result := i.ToString(g)
		if result != expected {
			t.Fatalf("Expected: [%v] Actual: [%v]", expected, result)
		}
	})
	t.Run("3", func(t *testing.T) {
		i := Item{Prod: rXS.Prods[0].Index, Dot: 2, Lookahead: tAdd.Index}
		if !i.IsKernel() {
			t.Fatalf("[%v] should be a kernel", i.ToString(g))
		}
		const expected = "XS = XS X., +"
		result := i.ToString(g)
		if result != expected {
			t.Fatalf("Expected: [%v] Actual: [%v]", expected, result)
		}
	})
	t.Run("Sprime", func(t *testing.T) {
		i := Item{Prod: sPrimeProdIndex, Dot: 0, Lookahead: eofIndex}
		if !i.IsKernel() {
			t.Fatalf("[%v] should be a kernel", i.ToString(g))
		}
		const expected = "S' = .XS, EOF"
		result := i.ToString(g)
		if result != expected {
			t.Fatalf("Expected: [%v] Actual: [%v]", expected, result)
		}
	})
}
