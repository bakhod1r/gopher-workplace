# List Copy Shares the Tail

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

Reusing `head.Next` shares the entire tail. A deep copy must recursively copy
the rest of the list.

## Task

Fix [clonesharetail.go](clonesharetail.go) to deep-copy the tail.

Do **not** change the function signature or the tests.

## Examples

```go
Copy(1->2->3) // independent 1->2->3
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Shallow vs deep** | Reusing Next shares nodes. |
| 2 | **Recursive copy** | `Copy(head.Next)`. |
| 3 | **Independence** | No shared nodes. |

## Hint

Recurse on the tail: `return &Node{Val: head.Val, Next: Copy(head.Next)}`.

## Validate

```bash
make verify
```
