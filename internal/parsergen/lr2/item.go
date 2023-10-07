package lr2

import (
	"fmt"
	"sort"
	"strings"
)

type Item struct {
	Prod      int
	Dot       int
	Lookahead int
}

func (i Item) IsKernel() bool {
	return i.Prod == SPrimeProd || i.Dot != 0
}

func (i Item) ToString(g *Grammar) string {
	var str strings.Builder

	prod := g.prods[i.Prod]
	fmt.Fprintf(&str, "%v = ", g.GetRule(prod.Rule).Name)
	for j, term := range prod.Terms {
		if j != 0 {
			str.WriteString(" ")
		}
		if j == i.Dot {
			str.WriteString(".")
		}
		str.WriteString(g.GetSymbolName(term))
	}
	if i.Dot == len(prod.Terms) {
		str.WriteString(".")
	}
	str.WriteString(", ")
	str.WriteString(g.GetSymbolName(i.Lookahead))
	return str.String()
}

func SortItems(items []Item) {
	sort.Slice(items, func(i, j int) bool {
		switch {
		case items[i].Prod < items[j].Prod:
			return true
		case items[i].Prod > items[j].Prod:
			return false
		case items[i].Dot < items[j].Dot:
			return true
		case items[i].Dot > items[j].Dot:
			return false
		default:
			return items[i].Lookahead < items[j].Lookahead
		}
	})
}