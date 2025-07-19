package rang3

import (
	"fmt"
	"slices"
	"sort"
	"strings"
	"unicode"

	"github.com/dcaiafa/lox/internal/base/stack"
)

const MaxRune rune = 0x10FFFF

type Range struct {
	B rune
	E rune
}

func (r Range) Contains(o Range) bool {
	return r.B <= o.B && r.E >= o.E
}

func (r Range) Intersects(o Range) bool {
	a, b := r, o
	if a.B > b.B {
		a, b = b, a
	}
	return b.B <= a.E
}

func (r Range) Touches(o Range) bool {
	a, b := r, o
	if a.B > b.B {
		a, b = b, a
	}
	return b.B <= a.E || b.B-1 == a.E
}

func (r Range) String() string {
	p := func(c rune) string {
		var buf strings.Builder
		switch c {
		case '\n':
			buf.WriteString(`\n`)
		case '\r':
			buf.WriteString(`\r`)
		case '\t':
			buf.WriteString(`\t`)
		case '-':
			buf.WriteString(`\-`)
		default:
			if unicode.IsGraphic(c) {
				buf.WriteRune(c)
			} else {
				fmt.Fprintf(&buf, "\\u%04x", c)
			}
		}
		return buf.String()
	}
	if r.B == r.E {
		return p(r.B)
	} else {
		return p(r.B) + "-" + p(r.E)
	}
}

func Compare(a, b Range) int {
	switch {
	case a.B < b.B:
		return -1
	case a.B > b.B:
		return 1
	case a.E < b.E:
		return -1
	case a.E > b.E:
		return 1
	default:
		return 0
	}
}

func Normalize(ranges []Range, onChange func(o, a, b, c Range)) {
	rh := newRangeHeap(ranges)

	for rh.Len() > 1 {
		x := rh.Pop()
		y := rh.Peek()

		if x == y {
			continue
		}

		if !x.Intersects(y) {
			continue
		}

		switch {
		case x.B == y.B && x.E < y.E:
			// x ---
			// y ------
			// ===================
			// x ---
			// a    ---
			a := Range{x.E + 1, y.E}
			onChange(y, x, a, a)
			rh.Pop() // pop y
			rh.Push(x)
			rh.Push(a)

		case x.B < y.B && x.E == y.E:
			// x ------
			// y    ---
			// ===================
			// a ---
			// y    ---
			a := Range{x.B, y.B - 1}
			onChange(x, a, y, y)
			rh.Push(a)

		case x.B < y.B && x.E < y.E:
			// x ------
			// y    ------
			// ==================
			// a ---
			// b    ---
			// c       ---
			a := Range{x.B, y.B - 1}
			b := Range{y.B, x.E}
			c := Range{x.E + 1, y.E}
			onChange(x, a, b, b)
			onChange(y, b, c, c)
			rh.Pop() // pop y
			rh.Push(a)
			rh.Push(b)
			rh.Push(c)

		case x.B < y.B && x.E > y.E:
			// x ---------
			// y    ---
			// ==================
			// a ---
			// y    ---
			// b       ---
			a := Range{x.B, y.B - 1}
			b := Range{y.E + 1, x.E}
			onChange(x, a, y, b)
			rh.Push(a)
			rh.Push(b)

		default:
			panic("not reached")
		}
	}
}

func Flatten(ranges []Range, onChange func(oa, ob, n Range)) []Range {
	slices.SortFunc(ranges, func(a, b Range) int {
		return Compare(a, b)
	})

	sort.Slice(ranges, func(i, j int) bool {
		switch {
		case ranges[i].B < ranges[j].B:
			return true
		case ranges[i].B > ranges[j].B:
			return false
		case ranges[i].E < ranges[j].B:
			return true
		default:
			return false
		}
	})
	min := func(a, b rune) rune {
		if a < b {
			return a
		} else {
			return b
		}
	}
	max := func(a, b rune) rune {
		if a > b {
			return a
		} else {
			return b
		}
	}

	ranges2 := stack.Stack[Range]{}
	ranges2.SetCapacity(len(ranges))
	for _, r := range ranges {
		if !ranges2.Empty() {
			tip := ranges2.Peek()
			if tip.Touches(r) {
				n := Range{B: min(tip.B, r.B), E: max(tip.E, r.E)}
				ranges2.Pop()
				ranges2.Push(n)
				if onChange != nil {
					onChange(tip, r, n)
				}
				continue
			}
		}
		ranges2.Push(r)
	}

	return ranges2.Elements()
}

func Subtract(a []Range, b []Range) []Range {
	if len(a) == 0 || len(b) == 0 {
		return a
	}

	a = Flatten(a, nil)
	b = Flatten(b, nil)

	r := stack.Stack[Range]{}
	r.SetCapacity(len(a))

	for len(b) > 0 {
		// a --
		// b   --
		if r.Empty() || r.Peek().E < b[0].B {
			if len(a) == 0 {
				break
			}
			r.Push(a[0])
			a = a[1:]
			continue
		}

		ea := r.Peek()
		eb := b[0]

		switch {
		// a   --
		// b --
		case ea.B > eb.E:
			b = b[1:]

		// a   --
		// b ------
		// r
		case ea.B >= eb.B && ea.E <= eb.E:
			r.Pop()

			// a ------
			// b   --
			// r --  --
		case ea.B < eb.B && ea.E > eb.E:
			r.Pop()
			r.Push(Range{B: ea.B, E: eb.B - 1})
			r.Push(Range{B: eb.B + 1, E: ea.E})

			// a ----
			// b   ----
			// r --
		case ea.B < eb.B && ea.E <= eb.E:
			r.Pop()
			r.Push(Range{B: ea.B, E: eb.B - 1})

			// a   ----
			// b ----
			// r     --
		case ea.B >= eb.B && ea.E > eb.E:
			r.Pop()
			r.Push(Range{B: eb.E + 1, E: ea.E})

		default:
			panic("unreachable")
		}
	}

	return append(r.Elements(), a...)
}
