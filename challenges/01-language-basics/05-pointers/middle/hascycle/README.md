# Detect a Cycle

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

Floyd's algorithm: a slow pointer (one step) and a fast pointer (two steps).
If they ever meet, there is a cycle; if fast reaches nil, there isn't.

## Task

Implement `HasCycle` in [hascycle.go](hascycle.go).

Do **not** change the function signature or the tests.

## Examples

```go
HasCycle(acyclic) // => false
HasCycle(cyclic)  // => true
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Slow/fast pointers** | Different speeds close any gap in a cycle. |
| 2 | **Meeting means cycle** | Equal pointers detect the loop. |
| 3 | **nil means acyclic** | fast hits nil in a finite list. |

## Hint

Advance `slow` by 1 and `fast` by 2 while `fast != nil && fast.Next != nil`; if `slow == fast` return true; else return false.

## Validate

```bash
make verify
```
