# From Addresses To Names

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

The profiler records program counters, not names. Turning a raw profile into something readable means resolving every address against the symbol table — the function whose code starts at or below it and whose successor starts above it. With tens of thousands of symbols and millions of addresses, the lookup has to be logarithmic.

## Task

Implement both functions in [samplesymbolize.go](samplesymbolize.go):

1. `Resolve` returns the symbol with the greatest `Start` not above `addr`, using a binary search over the sorted table.
2. An address below the first symbol resolves to `"", false`.
3. `Symbolize` resolves a whole address stack, dropping unresolvable ones and returning a non-nil slice.

## Examples

**Example 1:**

```
Input:  Resolve([{100 a} {200 b}], 150)
Output: "a", true
```

**Example 2:**

```
Input:  Resolve([{100 a} {200 b}], 200)
Output: "b", true
```

**Example 3:**

```
Input:  Resolve([{100 a}], 99)
Output: "", false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Predecessor search** | Not "find this address" but "find the last symbol at or below it". |
| 2 | **`sort.Search` semantics** | It returns the first index where the predicate holds, which is one past the symbol you want. |
| 3 | **Symbolisation is the expensive step** | Doing it linearly per address is what makes naive profile tools crawl. |

## Topics used again

Binary search, slices, boundary conditions.

## Hint

Search for the first symbol whose `Start` is greater than `addr`; the one before it is the answer.

## Validate

```bash
make verify
```
