# Forgotten Spread Operator

**Level:** senior
**Topic:** 01-language-basics → 04-functions · _variadic_

## Context

To pass a slice's elements to a variadic function you spread it with `xs...`.
The bug passes `len(xs)` (a single count) instead, so the total is meaningless.

## Task

Fix [spreadmissing.go](spreadmissing.go) to forward the slice elements.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Total([1 2 3 4])
Output: 10
```

**Example 2:**

```
Input:  Total(nil)
Output: 0
```

**Example 3:**

```
Input:  Total([5])
Output: 5
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Spread operator** | `sum(xs...)` passes each element. |
| 2 | **Variadic call forms** | Loose args or one spread slice. |
| 3 | **Not the length** | `len(xs)` is a count, not the data. |

## Hint

Spread the slice: `return sum(xs...)`.

## Validate

```bash
make verify
```
