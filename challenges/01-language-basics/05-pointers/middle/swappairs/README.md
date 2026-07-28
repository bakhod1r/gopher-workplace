# Swap First Two Nodes

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

Swapping the first two nodes rewires three pointers and returns the second node
as the new head — a pointer-surgery exercise.

## Task

Implement `SwapHead` in [swappairs.go](swappairs.go).

Do **not** change the function signature or the tests.

## Examples

```go
SwapHead(1->2->3) // => 2->1->3
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Guard short lists** | Need at least two nodes. |
| 2 | **Rewire pointers** | second.Next = head; head.Next = rest. |
| 3 | **New head** | Return the old second node. |

## Hint

If `head == nil || head.Next == nil { return head }`; `second := head.Next; head.Next = second.Next; second.Next = head; return second`.

## Validate

```bash
make verify
```
