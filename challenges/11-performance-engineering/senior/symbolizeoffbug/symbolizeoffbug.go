// Package symbolizeoffbug — Gopher Workplace challenge.
package symbolizeoffbug

import "sort"

// Symbol is one entry of a symbol table: a function and the address its code
// starts at. Tables are sorted by Start ascending.
type Symbol struct {
	Start uint64
	Func  string
}

// Resolve maps a program counter to the function containing it: the symbol
// with the greatest Start not above addr. An address below the first symbol
// resolves to "", false.
//
// Examples:
//
//	Resolve([{100,"a"},{200,"b"}], 150) => "a", true
func Resolve(table []Symbol, addr uint64) (string, bool) {
	// CHANGE CODE BELOW THIS LINE
	i := sort.Search(len(table), func(i int) bool { return table[i].Start >= addr })
	// CHANGE CODE ABOVE THIS LINE
	if i == 0 {
		return "", false
	}
	return table[i-1].Func, true
}

// Symbolize resolves a whole stack of addresses, dropping the unresolvable
// ones.
//
// Examples:
//
//	Symbolize([{100,"a"}], []uint64{150, 1}) => ["a"]
func Symbolize(table []Symbol, addrs []uint64) []string {
	out := make([]string, 0, len(addrs))
	for _, a := range addrs {
		if fn, ok := Resolve(table, a); ok {
			out = append(out, fn)
		}
	}
	return out
}
