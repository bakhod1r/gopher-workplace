# Word Frequency

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types
**Estimated time:** 10 min

## Context

A log analyzer needs a tally of how often each word shows up. A `map` from
word to count is the natural tool — but a `nil` map must not be written to, and
the empty case should still return a usable map.

## Task

Implement `Count` in [wordfreq.go](wordfreq.go) so it returns a
`map[string]int` of each word's occurrence count. Return an empty (non-nil) map
for `nil` or empty input.

Do **not** change the function signature or the tests.

## Examples

```go
Count([]string{"a", "b", "a"}) // => map[string]int{"a": 2, "b": 1}
Count([]string{"x"})           // => map[string]int{"x": 1}
Count(nil)                     // => map[string]int{}
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Map creation** | A map literal `map[string]int{}` or `make(map[string]int)` gives a writable, non-nil map. |
| 2 | **Nil map writes panic** | The zero value of a map is `nil`; reading it is fine, but writing to it panics — always initialize first. |
| 3 | **Zero value on read** | Reading a missing key yields the value type's zero (`0` for int), so `m[w]++` works even the first time. |

## Hint

Initialize the result with `map[string]int{}` up front, then range over the
words doing `counts[w]++`. The missing-key read returning `0` makes the
increment work without a special first-time case.

## Validate

```bash
make verify   # fmt-check + vet + test
```

Green tests + clean `vet`/`gofmt` = challenge passed.
