# First, Last, OK

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _multiple-return_

## Context

Returning an `ok bool` alongside the data lets the caller distinguish a real
result from the empty case without a sentinel value or panic.

## Task

Implement `FirstLast` in [firstlast.go](firstlast.go).

Do **not** change the function signature or the tests.

## Examples

```go
FirstLast(nil)              // => 0, 0, false
FirstLast([]int{7})         // => 7, 7, true
FirstLast([]int{2, 4, 6, 8}) // => 2, 8, true
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Comma-ok style return** | An extra `bool` signals validity. |
| 2 | **Empty guard** | `len(xs) == 0` returns the zero result early. |
| 3 | **Index math** | Last element is `xs[len(xs)-1]`. |

## Hint

Guard `len(xs)==0` first, else return `xs[0], xs[len(xs)-1], true`.

## Validate

```bash
make verify
```
