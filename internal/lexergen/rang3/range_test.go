package rang3

import (
	"math/rand"
	"reflect"
	"slices"
	"testing"

	"github.com/dcaiafa/lox/internal/base/set"
)

func TestRangeContains(t *testing.T) {
	test := func(name string, ab, ae, bb, be rune, expect bool) {
		t.Run(name, func(t *testing.T) {
			a := Range{B: ab, E: ae}
			b := Range{B: bb, E: be}
			result := a.Contains(b)
			if result != expect {
				t.Fatalf(
					"Ranges: %v %v expected: %v actual: %v",
					a, b, expect, result)
			}
		})
	}
	//    1     2     3     4     5     6    7    8
	// a ---    ---   --   ----  --    --  --      --
	// b  ---  ---   ----   --    --  --     --  --
	//    F      F     F    T     F    F     F    F
	test("1", 'a', 'd', 'b', 'e', false)
	test("2", 'b', 'e', 'a', 'f', false)
	test("3", 'c', 'd', 'a', 'f', false)
	test("4", 'a', 'f', 'c', 'd', true)
	test("5", 'a', 'b', 'b', 'c', false)
	test("6", 'b', 'c', 'a', 'b', false)
	test("7", 'a', 'b', 'c', 'd', false)
	test("8", 'c', 'd', 'a', 'b', false)
	test("9", 'c', 'd', 'c', 'd', true)
}

func TestRangeIntersects(t *testing.T) {
	test := func(name string, ab, ae, bb, be rune, expect bool) {
		t.Run(name, func(t *testing.T) {
			a := Range{B: ab, E: ae}
			b := Range{B: bb, E: be}
			result := a.Intersects(b)
			if result != expect {
				t.Fatalf(
					"Ranges: %v %v expected: %v actual: %v",
					a, b, expect, result)
			}
		})
	}

	//    1     2     3     4     5     6    7    8
	// a ---    ---   --   ----  --    --  --      --
	// b  ---  ---   ----   --    --  --     --  --
	//    T      T     T    T     T    T     F    F
	test("1", 'a', 'd', 'b', 'e', true)
	test("2", 'b', 'e', 'a', 'd', true)
	test("3", 'c', 'd', 'a', 'f', true)
	test("4", 'a', 'f', 'c', 'd', true)
	test("5", 'a', 'b', 'b', 'c', true)
	test("6", 'b', 'c', 'a', 'b', true)
	test("7", 'a', 'b', 'c', 'd', false)
	test("8", 'c', 'd', 'a', 'b', false)
}

func TestRangeTouches(t *testing.T) {
	test := func(name string, ab, ae, bb, be rune, expect bool) {
		t.Run(name, func(t *testing.T) {
			a := Range{B: ab, E: ae}
			b := Range{B: bb, E: be}
			result := a.Touches(b)
			if result != expect {
				t.Fatalf(
					"Ranges: %v %v expected: %v actual: %v",
					a, b, expect, result)
			}
		})
	}

	//    1     2     3     4     5     6    7    8      9    10
	// a ---    ---   --   ----  --    --  --      --  --       --
	// b  ---  ---   ----   --    --  --     --  --       -- --
	//    T      T     T    T     T    T     T    T      F    F
	test("1", 'a', 'd', 'b', 'e', true)
	test("2", 'b', 'e', 'a', 'd', true)
	test("3", 'c', 'd', 'a', 'f', true)
	test("4", 'a', 'f', 'c', 'd', true)
	test("5", 'a', 'b', 'b', 'c', true)
	test("6", 'b', 'c', 'a', 'b', true)
	test("7", 'a', 'b', 'c', 'd', true)
	test("8", 'c', 'd', 'a', 'b', true)
	test("9", 'a', 'b', 'd', 'e', false)
	test("10", 'd', 'e', 'a', 'b', false)
}

func Itor(input []int) []Range {
	inputr := make([]Range, 0, len(input)/2)
	for i := 0; i < len(input); i += 2 {
		inputr = append(inputr, Range{
			B: rune(input[i]),
			E: rune(input[i+1]),
		})
	}
	return inputr
}

func ItorS(input []int) []Range {
	r := Itor(input)
	rand.Shuffle(len(r), func(i, j int) {
		r[i], r[j] = r[j], r[i]
	})
	return r
}

func DumpRanges(t *testing.T, rs []Range) {
	for _, r := range rs {
		t.Log(int(r.B), "-", int(r.E))
	}
}

func TestFlattenRanges(t *testing.T) {
	test := func(name string, input, output []int) {
		t.Run(name, func(t *testing.T) {
			res := Flatten(Itor(input), nil)
			expected := Itor(output)
			if !reflect.DeepEqual(res, expected) {
				t.Log("Expected:")
				DumpRanges(t, expected)
				t.Log("Actual:")
				DumpRanges(t, res)
				t.Fatalf("Unexpected result")
			}
		})
	}

	//    0123456789
	//    ---
	//            --
	//     ---
	// R: ----    --
	test("1", []int{0, 2, 8, 9, 1, 3}, []int{0, 3, 8, 9})

	//    0123456789
	//        ---
	//    ---    ---
	// R: --- ------
	test("2", []int{4, 6, 0, 2, 7, 9}, []int{0, 2, 4, 9})

	//    0123456789
	//        ---
	//    ---     --
	// R: --- --- --
	test("3", []int{4, 6, 0, 2, 8, 9}, []int{0, 2, 4, 6, 8, 9})

	//    0123456789
	//             -
	//            -
	//           -
	//          -
	//        -
	//       -
	//      -
	// R:   --- ----
	test("4", []int{9, 9, 8, 8, 7, 7, 6, 6, 4, 4, 3, 3, 2, 2}, []int{2, 4, 6, 9})
}

func TestNormalizeRanges(t *testing.T) {
	test := func(name string, input, output []int) {
		t.Run(name, func(t *testing.T) {
			inputRanges := ItorS(input)
			res := set.Set[Range]{}
			res.AddSlice(Itor(input))
			Normalize(inputRanges, func(o, a, b, c Range) {
				if !res.Has(o) {
					t.Fatalf("Range doesn't exist: %v", o)
				}
				res.Remove(o)
				res.Add(a)
				res.Add(b)
				if c != b {
					res.Add(c)
				}
			})
			resSlice := res.Elements()
			slices.SortFunc(resSlice, Compare)
			expected := Itor(output)
			if !reflect.DeepEqual(resSlice, expected) {
				t.Log("Expected:")
				DumpRanges(t, expected)
				t.Log("Actual:")
				DumpRanges(t, resSlice)
				t.Fatalf("Unexpected result")
			}
		})
	}
	//   0123456789
	// x ---
	// y ------
	// ===================
	// x ---
	// a    ---
	test("1", []int{0, 2, 0, 5}, []int{0, 2, 3, 5})

	//   0123456789
	// x ------
	// y    ---
	// ===================
	// a ---
	// y    ---
	test("2", []int{0, 5, 3, 5}, []int{0, 2, 3, 5})

	//   0123456789
	// x ------
	// y    ------
	// ==================
	// a ---
	// b    ---
	// c       ---
	test("3", []int{0, 5, 3, 8}, []int{0, 2, 3, 5, 6, 8})

	//   0123456789
	//   ---------
	//      ---
	// ==================
	//   ---
	//      ---
	//         ---
	test("4", []int{0, 8, 3, 5}, []int{0, 2, 3, 5, 6, 8})

	//   0123456789
	//   ---------
	//      ---
	//      -
	//           -
	//            -
	// ==================
	//   ---
	//      -
	//       --
	//         --
	//           -
	//            -
	test("complex1",
		[]int{
			0, 8,
			3, 5,
			3, 3,
			8, 8,
			9, 9,
		},
		[]int{
			0, 2,
			3, 3,
			4, 5,
			6, 7,
			8, 8,
			9, 9,
		})

	//   0123456789
	//     --------
	//        -
	//         ---
	//         ----
	//  ===========
	//     ---
	//        -
	//         ---
	//            -
	test("complex2",
		[]int{
			2, 9,
			5, 5,
			6, 8,
			6, 9,
		},
		[]int{
			2, 4,
			5, 5,
			6, 8,
			9, 9,
		},
	)
}

func TestSubtract(t *testing.T) {
	test := func(name string, inputA, inputB, output []int) {
		t.Run(name, func(t *testing.T) {
			res := Subtract(ItorS(inputA), ItorS(inputB))
			expected := Itor(output)
			if !reflect.DeepEqual(res, expected) {
				t.Log("Expected:")
				DumpRanges(t, expected)
				t.Log("Actual:")
				DumpRanges(t, res)
				t.Fatalf("Unexpected result")
			}
		})
	}

	//   0123456789
	// a ----- ----
	// b --  -- --
	// r   --  -  -
	test("1", []int{0, 4, 6, 9}, []int{0, 1, 4, 5, 7, 8}, []int{2, 3, 6, 6, 9, 9})

	//   0123456789
	// a  ---- ----
	// b ---
	// r    -- ----
	test("2", []int{1, 4, 6, 9}, []int{0, 2}, []int{3, 4, 6, 9})

	//   0123456789
	// a ----- ----
	// b     ------
	// r ----
	test("3", []int{0, 4, 6, 9}, []int{4, 9}, []int{0, 3})

	//   0123456789
	// a ----- ----
	// b        ---
	// r ----- -
	test("4", []int{0, 4, 6, 9}, []int{7, 9}, []int{0, 4, 6, 6})
}
