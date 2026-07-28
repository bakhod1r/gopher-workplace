# Pop Without Empty Guard

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _nil-pointer-dereference_

## Context

Popping from an empty queue dereferences a nil head. Guard `q.head == nil` and
return the not-ok result before touching the node.

## Task

Fix [popnoguard.go](popnoguard.go) so popping an empty queue returns false without panicking.

Do **not** change the function signature or the tests.

## Examples

```go
(&Queue{}).Pop() // => 0, false
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Empty guard** | Check `q.head == nil` first. |
| 2 | **Nil dereference** | `q.head.Val` panics on empty. |
| 3 | **Not-ok result** | Return 0,false when empty. |

## Hint

Guard first: `if q.head == nil { return 0, false }`, then pop.

## Validate

```bash
make verify
```
