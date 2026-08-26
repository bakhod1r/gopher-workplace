# Build String

**Level:** middle
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A report generator builds lines by chaining `.Add()` calls, then `.Build()`
joins everything with a separator.

## Task

Implement `Add` on `*Builder` in [buildstr.go](buildstr.go):

1. Append `part` to `b.parts`.
2. Return `b` for chaining.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  NewBuilder(", ").Add("a").Add("b").Add("c").Build()
Output: "a, b, c"
```

**Example 2:**

```
Input:  NewBuilder("-").Add("2024").Add("01").Add("01").Build()
Output: "2024-01-01"
```

**Example 3:**

```
Input:  NewBuilder("|").Build()
Output: ""
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Builder pattern** | Accumulate state, produce result at the end. |
| 2 | **Method chaining** | `Add` returns `*Builder` for fluent API. |
| 3 | **strings.Join** | Joins a `[]string` with a separator. |

## Hint

`b.parts = append(b.parts, part); return b`.

## Validate

```bash
make verify
```
