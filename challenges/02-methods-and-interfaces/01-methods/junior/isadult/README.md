# Is Adult

**Level:** junior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

An age-gated service checks whether a user is old enough to proceed. The check
is a method on the `Person` struct.

## Task

Implement `IsAdult` on `Person` in [isadult.go](isadult.go):

1. Return `true` if `Age >= 18`.
2. Return `false` otherwise.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Person{"Alice", 25}.IsAdult()
Output: true
```

**Example 2:**

```
Input:  Person{"Bob", 17}.IsAdult()
Output: false
```

**Example 3:**

```
Input:  Person{"Carol", 18}.IsAdult()
Output: true
```

_Explanation:_ Exactly 18 counts as adult.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Methods** | `IsAdult` is bound to `Person` via receiver. |
| 2 | **Value receiver** | Read-only — just checks `Age`. |
| 3 | **Boolean return** | Return the comparison directly. |

## Hint

`return p.Age >= 18` — a single expression, no `if` needed.

## Validate

```bash
make verify
```
