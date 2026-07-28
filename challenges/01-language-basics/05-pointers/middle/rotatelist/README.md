# Rotate List Right

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

Rotating right by k: form a ring, then break it k%len nodes before the end.
Handle k >= len with the modulo.

## Task

Implement `Rotate` in [rotatelist.go](rotatelist.go).

Do **not** change the function signature or the tests.

## Examples

```go
Rotate(1->2->3->4->5, 2) // => 4->5->1->2->3
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Length and modulo** | Effective shift is `k % len`. |
| 2 | **Close the ring** | Link tail to head. |
| 3 | **Break at the right spot** | New tail is `len - k%len - 1` from head. |

## Hint

Compute length, link tail to head to form a ring, walk `len - k%len` steps to the new tail, set new head = newTail.Next, break the ring.

## Validate

```bash
make verify
```
