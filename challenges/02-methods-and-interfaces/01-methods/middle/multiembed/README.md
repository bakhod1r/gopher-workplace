# Multiple Embedding

**Level:** middle
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

If a struct embeds two types that have the same method name, the compiler
doesn't know which one to promote. You must resolve the ambiguity.

## Task

Implement `Name` on `Collision` in [multiembed.go](multiembed.go):

1. Explicitly return the result of `B`'s `Name()` method.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Collision{}.Name()
Output: "B"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Embedding collision** | Ambiguous promoted methods cause compile errors. |
| 2 | **Explicit resolution** | Define the method on the outer struct and route it to the specific embedded field: `c.B.Method()`. |

## Hint

`return c.B.Name()`.

## Validate

```bash
make verify
```
