# List Length Counts a Phantom

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

The counter must start at 0; the loop already visits every node. Starting at 1
counts a node that isn't there (and makes nil report 1).

## Task

Fix the counter initialisation in [lenoffbyone.go](lenoffbyone.go).

Do **not** change the function signature or the tests.

## Examples

```go
Length(nil) // => 0
Length(1->2->3) // => 3
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Counter start** | Begin at 0. |
| 2 | **Loop visits all** | Each node increments once. |
| 3 | **Empty case** | nil must be 0. |

## Hint

Start at zero: `count := 0`.

## Validate

```bash
make verify
```
