# Deterministic Map Output

**Level:** middle  
**Topic:** 03-generics

## Context

A snapshot test kept failing at random because the fixture was rendered straight from a map.

## Task

Implement the stub(s) in [detmapgen.go](detmapgen.go):

1. Implement `Entries`, returning the map's entries sorted by key.
2. Return an empty (non-nil) slice for empty or nil input.
3. The result must be identical across runs.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Entries({b:1, a:2})
Output: [{a 2} {b 1}]
```

**Example 2:**

```
Input:  repeated calls
Output: identical output
```

**Example 3:**

```
Input:  Entries(nil)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Randomised map order** | Go randomises range order deliberately, to stop code depending on it. |
| 2 | **Sorting for determinism** | Ordered keys make output reproducible in tests and diffs. |
| 3 | **Pairs as a struct** | `Entry[K, V]` keeps key and value together through the sort. |

## Hint

Collect first, then sort by key — never rely on the range order.

## Validate

```bash
make verify
```
