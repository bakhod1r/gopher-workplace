# Range Copy Doesn't Mutate

**Level:** senior
**Topic:** 01-language-basics → 04-functions · _loops_

## Context

`for i, v := range xs` copies each element into `v`. Assigning to `v` changes
only the copy; to mutate the slice you must write through the index `xs[i]`.

## Task

Fix [rangemutate.go](rangemutate.go) so the slice is actually doubled.

Do **not** change the function signature or the tests.

## Examples

```go
DoubleAll([1 2 3]) // slice becomes [2 4 6]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Range value is a copy** | `v` does not alias `xs[i]`. |
| 2 | **Index to mutate** | Write `xs[i] = ...`. |
| 3 | **In-place mutation** | The caller's backing array must change. |

## Hint

Write through the index: `xs[i] = v * 2` (or `xs[i] *= 2`).

## Validate

```bash
make verify
```
