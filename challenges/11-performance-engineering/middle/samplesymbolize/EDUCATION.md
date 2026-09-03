# From Addresses To Names

## Intuition

The symbol table partitions the address space into half-open intervals. Resolving an address means finding which interval it fell into, which is a predecessor query, not an equality search.

## Approach

1. Binary search for the first symbol starting *above* `addr`.
2. Step back one; if there is no previous symbol, the address is below the table.
3. `Symbolize` maps the stack, skipping failures.

## Solution

```go
func Resolve(table []Symbol, addr uint64) (string, bool) {
	i := sort.Search(len(table), func(i int) bool { return table[i].Start > addr })
	if i == 0 {
		return "", false
	}
	return table[i-1].Func, true
}

func Symbolize(table []Symbol, addrs []uint64) []string {
	out := make([]string, 0, len(addrs))
	for _, a := range addrs {
		if fn, ok := Resolve(table, a); ok {
			out = append(out, fn)
		}
	}
	return out
}
```

## Walkthrough

For `addr = 200` in a table starting at 100, 200, 300, the search returns index 2 (the first `Start > 200`), and `table[1]` is `b` — the symbol that starts exactly at the address.

## Pitfalls

- Searching for `>= addr`, which pushes an address landing exactly on a symbol boundary into the previous function.
- Forgetting the `i == 0` case and indexing `table[-1]`.
- Scanning the table linearly per address, which turns symbolisation into the slowest part of the tool.
