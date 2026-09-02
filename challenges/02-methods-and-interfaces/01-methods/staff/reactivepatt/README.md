# Reactive Stream

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A stream pipeline reads as a sentence: filter, then map, then collect. Each
operator returns the stream so the next call can chain. Here the operators work
in place — the pipeline mutates one slice rather than allocating a new one per
stage.

## Task

Implement `Filter` and `Map` on `*Stream` in [reactivepatt.go](reactivepatt.go):

1. `Filter(fn)` keeps only the elements for which `fn` returns true, in place,
   and returns `s`.
2. `Map(fn)` replaces each element with `fn(element)`, in place, and returns `s`.

**Constraint (staff):** the pipeline must run in place — filtering and mapping a 1M-element stream must allocate nothing.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  [1 2 3 4].Filter(even)
Output: [2 4]
```

**Example 2:**

```
Input:  [2 4].Map(x*10)
Output: [20 40]
```

**Example 3:**

```
Input:  [1 2 3 4].Filter(even).Map(x*10)
Output: [20 40]
```

_Explanation:_ each operator returns the receiver, so the calls chain.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Returning the receiver** | `return s` is what makes `.Filter(...).Map(...)` compile. |
| 2 | **Filter-in-place idiom** | `out := s.Data[:0]` reuses the backing array; `append` writes behind the read cursor. |
| 3 | **Index-based map** | `for i := range` to write back; the value form would discard the result. |

## Hint

The zero-allocation filter is `out := s.Data[:0]` followed by an `append` loop,
then `s.Data = out`. It is safe because the write index never overtakes the read
index.

## Validate

```bash
make verify
```
