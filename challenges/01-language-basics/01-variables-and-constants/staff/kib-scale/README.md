# Kibi vs Kilo

**Level:** staff
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

A kibibyte (KiB) is `1024`, not `1000`. The decimal factor silently
under-reports every size by ~2.4% — the kind of unit bug that corrupts capacity
math at scale.

## Task

Fix the single line between the markers in [scale.go](scale.go) so `KiB` is the
binary factor.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  KiB
Output: 1024
```

**Example 2:**

```
Input:  Bytes(2)
Output: 2048
```

**Example 3:**

```
Input:  Bytes(1)
Output: 1024
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Binary vs decimal units** | KiB=1024, kB=1000. |
| 2 | **Named constant** | One source of truth for the factor. |
| 3 | **Compile-time fold** | `n * KiB` folds the constant in. |

## Hint

`const KiB = 1024` (or `1 << 10`).

## Validate

```bash
make verify
```
