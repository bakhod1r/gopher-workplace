# Discarded Remainder

**Level:** junior
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

A pagination helper already computes both halves of one division: how many whole
pages fit, and what is left over. Two thin views sit on top of it — one wants
only the page count, the other only the remainder.

Go has no "just give me the first one" rule. A call that returns two values must
be received in full, or it does not compile. The blank identifier `_` is how you
accept a value and immediately throw it away.

## Task

Implement `Pages` and `Leftover` in [discard.go](discard.go) so that:

1. Both **call `Split`** rather than recomputing the arithmetic.
2. `Pages` keeps the page count and discards the remainder.
3. `Leftover` keeps the remainder and discards the page count.
4. Both stay correct for `size == 0`, which `Split` already handles.

Do **not** change the function signatures or the tests.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Pages(10, 3)
Output: 3
```

**Example 2:**

```
Input:  Leftover(10, 3)
Output: 1
```

**Example 3:**

```
Input:  Pages(2, 5)
Output: 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Blank identifier** | `_` is a write-only sink: assigning to it discards the value, and it never occupies a name. |
| 2 | **Multi-value assignment** | A two-value call must be received by exactly two operands — `x := Split(...)` is a compile error, not a shortcut. |
| 3 | **Unused variables are errors** | Naming the value you do not want (`pages, rest := ...` then ignoring `rest`) fails to compile; `_` is the escape hatch. |
| 4 | **Named results** | `Split` names its results `(pages, rest int)` — documentation for the caller; the names do not change how the call is received. |

## Hint

`pages := Split(n, size)` will not compile: the call yields two values. Write
both operands and park the unwanted one in `_`:

```go
pages, _ := Split(n, size)
_, rest := Split(n, size)
```

Do not be tempted to declare `rest` and leave it unused — Go rejects unused
local variables, which is exactly why `_` exists.

## Validate

```bash
make verify
```
