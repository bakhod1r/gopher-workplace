# Take While

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

TakeWhile consumes a prefix, breaking out of the loop as soon as the predicate
fails — order-sensitive, unlike Filter.

## Task

Implement `TakeWhile` in [takewhile.go](takewhile.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  TakeWhile([2 4 5 6], even)
Output: [2 4]
```

**Example 2:**

```
Input:  TakeWhile([1 2], even)
Output: []
```

**Example 3:**

```
Input:  TakeWhile([2 4], even)
Output: [2 4]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Break on first failure** | Stop, don't skip. |
| 2 | **Prefix semantics** | Only the leading run is taken. |
| 3 | **Predicate arg** | `pred func(int) bool`. |

## Hint

Range `xs`; `if !pred(v) { break }; out = append(out, v)`.

## Validate

```bash
make verify
```
