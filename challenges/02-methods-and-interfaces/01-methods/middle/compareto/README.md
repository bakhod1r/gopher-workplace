# Compare To

**Level:** middle
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A package manager compares versions to decide which is newer.

## Task

Implement `Compare` on `Version` in [compareto.go](compareto.go):

1. Return `-1` if `v < other`, `0` if equal, `+1` if greater.
2. Compare `Major` first, then `Minor`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Version{1, 0}.Compare(Version{2, 0})
Output: -1
```

**Example 2:**

```
Input:  Version{1, 2}.Compare(Version{1, 2})
Output: 0
```

**Example 3:**

```
Input:  Version{2, 0}.Compare(Version{1, 9})
Output: 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Value receiver** | Both versions are read-only. |
| 2 | **Multi-key comparison** | Compare primary key first, then secondary. |
| 3 | **Defined types** | `Version` is a struct with comparison semantics. |

## Hint

Compare `Major` first; if equal, compare `Minor`.

## Validate

```bash
make verify
```
