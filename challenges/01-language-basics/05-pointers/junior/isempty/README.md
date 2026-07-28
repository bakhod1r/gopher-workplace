# Nil or Empty List

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _nil-pointer-dereference_

## Context

A nil head is the canonical empty-list representation; the check is a simple
nil comparison.

## Task

Implement `IsEmpty` in [isempty.go](isempty.go).

Do **not** change the function signature or the tests.

## Examples

```go
IsEmpty(nil)      // => true
IsEmpty(&Node{})  // => false
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **nil as empty** | A nil head means no nodes. |
| 2 | **Nil comparison** | `head == nil`. |
| 3 | **No deref needed** | Just compare. |

## Hint

`return head == nil`.

## Validate

```bash
make verify
```
