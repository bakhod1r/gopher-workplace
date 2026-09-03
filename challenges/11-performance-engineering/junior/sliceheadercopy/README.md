# The Header Is A Copy, The Array Is Not

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

Passing a slice copies three words: pointer, length, capacity. Writes through it reach the caller because the pointer is shared; `append` does not, because the length that changed was the local copy. Every "why didn't my append show up" bug, and every accidental aliasing bug, comes from this one fact.

## Task

Implement both functions in [sliceheadercopy.go](sliceheadercopy.go):

1. `Fill` overwrites every element of `s` with `v`, allocating nothing.
2. `Fill` handles nil and empty slices without panicking.
3. `AppendLocal` appends `v` to its local copy of the header — so the caller's length is unchanged, while the spare capacity really was written.

## Examples

**Example 1:**

```
Input:  s = [1 2 3]; Fill(s, 7)
Output: s is [7 7 7]
```

**Example 2:**

```
Input:  s = make([]int, 3, 8); AppendLocal(s, 7)
Output: len(s) is still 3
```

**Example 3:**

```
Input:  the same s, read as s[:4]
Output: s[:4][3] is 7 — the array was written
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Three words, passed by value** | Pointer, length and capacity are copied; the array behind them is not. |
| 2 | **Why `append` must return** | It may change the length, the capacity, or the pointer, and only the return carries that. |
| 3 | **Writing past `len` is real** | The spare capacity belongs to the shared array, so a stray append is visible to anyone resslicing it. |

## Topics used again

Slices, `range`, `append`, pointers.

## Hint

`AppendLocal` genuinely is a one-line `append` whose result goes nowhere — that is the lesson, not a mistake.

## Validate

```bash
make verify
```
