# What Boxing An `int` Costs

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

An `any` is two words: a type pointer and a data pointer. Putting an `int` in one therefore needs somewhere for the int to live — the heap, one allocation per value. Except for 0..255, which the runtime keeps in a static table, so those box for free. That is why a profile of `[]any` code sometimes shows a million allocations and sometimes none.

## Task

Implement both functions in [boxedintcount.go](boxedintcount.go):

1. `Box` converts each `int` to `any`, building the result in a single allocation.
2. `Unbox` extracts the `int` values and returns how many elements were not ints.
3. Empty inputs return empty, non-nil slices.

## Examples

**Example 1:**

```
Input:  Box([1 2 3])
Output: [1 2 3] as []any
```

**Example 2:**

```
Input:  Unbox([1 "x" 3 nil])
Output: [1 3], 2
```

**Example 3:**

```
Input:  Box(100 values in 0..255)
Output: 1 allocation total
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Interfaces hold pointers** | Any value wider than a word must be heap-allocated to be stored in one. |
| 2 | **The small-int table** | Values 0..255 box without allocating, which makes microbenchmarks lie. |
| 3 | **Type assertion with comma-ok** | The safe way to get the value back out without panicking on the wrong type. |

## Topics used again

`any`, type assertions, `make` with a capacity hint.

## Hint

Both functions are a preallocated slice and one loop; `Unbox` needs the comma-ok form.

## Validate

```bash
make verify
```
