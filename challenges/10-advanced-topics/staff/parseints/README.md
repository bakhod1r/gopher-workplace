# Parse Without Making A Single String

**Level:** staff
**Topic:** 10-advanced-topics / 02-escape-analysis

## Context

An ingest hot path splits each line with `strings.Split(string(line), ",")` and calls `strconv.Atoi` on every field. At a million lines a second the allocator is the bottleneck, not the parsing.

## Task

Implement [parseints.go](parseints.go):

1. Sum the decimal integers in `line`, separated by `sep`.
2. Return the total, how many fields were parsed, and `ErrSyntax` for anything that is not a decimal integer.
3. An empty input is 0, 0, nil; an empty field is a syntax error.
4. Zero allocations — no string conversion, no split slice.

Replace the stub body in [parseints.go](parseints.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  ParseInts([]byte("1,2,3"), ',')
Output: 6, 3, nil
```

**Example 2:**

```
Input:  ParseInts([]byte("-4,+6"), ',')
Output: 2, 2, nil
```

_Explanation:_ Signs are accepted.

**Example 3:**

```
Input:  ParseInts([]byte("1,,2"), ',')
Output: ErrSyntax
```

_Explanation:_ An empty field is invalid.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Parsing in place** | Digits can be folded into an accumulator straight from the bytes. |
| 2 | **Split allocates twice** | A conversion for the string and a slice for the pieces. |
| 3 | **Sentinel errors on the hot path** | A package-level error keeps the failure path allocation-free too. |
| 4 | **Boundary handling** | Running the loop to `len(line)` inclusive closes the final field. |

## Hint

`v = v*10 + int64(c-'0')` is the whole parser. The rest is where the fields start and end.

## Validate

```bash
make verify
```
