# Ordered Slice Delete

**Level:** senior
**Topic:** 01-language-basics → 04-functions · _loops_

## Context

Deleting index i means keeping `xs[:i]` and everything AFTER i, i.e. `xs[i+1:]`.
The bug appends `xs[i:]`, which re-includes the element meant to be removed.

## Task

Fix the delete in [deletebug.go](deletebug.go).

Do **not** change the function signature or the tests.

## Examples

```go
RemoveAt([10 20 30 40], 1) // => [10 30 40]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Ordered delete idiom** | `append(xs[:i], xs[i+1:]...)`. |
| 2 | **Off-by-one slice bound** | Skip index i by starting the tail at i+1. |
| 3 | **append with spread** | The tail is spread back onto the prefix. |

## Hint

The tail must start after i: `append(xs[:i], xs[i+1:]...)`.

## Validate

```bash
make verify
```
