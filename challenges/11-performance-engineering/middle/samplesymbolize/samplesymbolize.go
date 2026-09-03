// Package samplesymbolize — Gopher Workplace challenge.
package samplesymbolize

// Symbol is one entry of a symbol table: the function's name and the address
// its code starts at.
type Symbol struct {
	Start uint64
	Func  string
}

// Resolve maps a program counter to the function containing it: the symbol
// with the greatest Start not above addr. The table must be sorted by Start
// ascending; use a binary search rather than a scan, because real tables hold
// tens of thousands of symbols and a profile holds millions of addresses.
// An address below the first symbol resolves to "", false.
//
// Examples:
//
//	Resolve([{100,"a"},{200,"b"}], 150) => "a", true
func Resolve(table []Symbol, addr uint64) (string, bool) {
	panic("not implemented")
}

// Symbolize resolves a whole stack of addresses, dropping the ones that fall
// outside the table. An empty result is non-nil.
//
// Examples:
//
//	Symbolize([{100,"a"}], []uint64{150, 1}) => ["a"]
func Symbolize(table []Symbol, addrs []uint64) []string {
	panic("not implemented")
}
