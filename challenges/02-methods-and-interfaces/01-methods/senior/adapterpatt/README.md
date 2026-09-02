# Adapter Pattern

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A legacy service hands you numbers as strings. The rest of the codebase wants
`int`. Instead of sprinkling `strconv.Atoi` at every call site, wrap the legacy
type in an adapter whose method exposes the shape callers actually want.

## Task

Implement `GetIntData` on `*ModernAdapter` in [adapterpatt.go](adapterpatt.go):

1. Call `a.legacy.GetStringData()`.
2. Parse it with `strconv.Atoi`.
3. On a parse error return `0`; otherwise return the parsed value.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  (&ModernAdapter{}).GetIntData()   // legacy returns "123"
Output: 123
```

**Example 2:**

```
Input:  legacy data "0"
Output: 0
```

**Example 3:**

```
Input:  legacy data "abc"
Output: 0
```

_Explanation:_ `strconv.Atoi` fails, the adapter swallows the error and returns the zero value.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Adapter method** | The method — not the caller — owns the conversion, so the legacy type stays untouched. |
| 2 | **Embedded-by-field wrapping** | `legacy` is an ordinary field; the adapter delegates explicitly rather than promoting. |
| 3 | **Two-value returns** | `strconv.Atoi` returns `(int, error)`; ignoring the error silently is a deliberate choice here. |

## Hint

`n, err := strconv.Atoi(...)`. On `err != nil` return `0` — do not return `n`,
which is only guaranteed to be meaningful when `err == nil`.

## Validate

```bash
make verify
```
