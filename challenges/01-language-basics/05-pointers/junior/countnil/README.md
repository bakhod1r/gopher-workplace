# Count Nil Pointers

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _nil-pointer-dereference_

## Context

A pointer can be compared to `nil` to test whether it references anything.

## Task

Implement `CountNil` in [countnil.go](countnil.go).

Do **not** change the function signature or the tests.

## Examples

```go
CountNil([]*int{&a, nil, nil}) // => 2
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **nil comparison** | `p == nil` tests emptiness. |
| 2 | **Iterate pointers** | Range the slice. |
| 3 | **Counter** | Increment on nil. |

## Hint

Range the slice; `if p == nil { c++ }`.

## Validate

```bash
make verify
```
