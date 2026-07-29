# Drop While

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

DropWhile finds the first element failing the predicate and returns everything
from there — the complement of TakeWhile.

## Task

Implement `DropWhile` in [dropwhile.go](dropwhile.go). Return a fresh slice (copy the tail).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  DropWhile([2 4 5 6], even)
Output: [5 6]
```

**Example 2:**

```
Input:  DropWhile([1 2], even)
Output: [1 2]
```

**Example 3:**

```
Input:  DropWhile([2 4], even)
Output: []
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Find boundary** | Advance while pred holds. |
| 2 | **Copy the tail** | Return a new slice to avoid aliasing. |
| 3 | **Complements TakeWhile** | Drops the prefix TakeWhile keeps. |

## Hint

Advance `i` while `i < len(xs) && pred(xs[i])`; return a copy of `xs[i:]`.

## Validate

```bash
make verify
```
