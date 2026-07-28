# SliceData of Empty Slice

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _nil-pointer-dereference_

## Context

For an empty slice `unsafe.SliceData(s)` may be nil; dereferencing it panics.
Guard on `len(s) == 0` before reading the element.

## Task

Fix [nilslicedata.go](nilslicedata.go) so an empty slice returns the default.

Do **not** change the function signature or the tests.

## Examples

```go
FirstOr(nil, 7) // => 7
FirstOr([]int{5}, 7) // => 5
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Empty guard** | Check `len(s) == 0` first. |
| 2 | **SliceData may be nil** | Empty slices have no data pointer to read. |
| 3 | **Default** | Return def when empty. |

## Hint

Guard length first: `if len(s) == 0 { return def }; return *unsafe.SliceData(s)`.

## Validate

```bash
make verify
```
