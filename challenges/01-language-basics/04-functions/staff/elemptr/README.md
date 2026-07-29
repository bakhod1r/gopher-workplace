# Pointer to Element Before Grow

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

A full slice reallocates on append; a pointer taken beforehand still points at
the OLD array. Writing through it leaves the returned (new) slice unchanged.

## Task

Fix [elemptr.go](elemptr.go) so the result's first element is 42.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  BumpHead(9)
Output: [42 9]
```

**Example 2:**

```
Input:  write lands on element 0
Output: true
```

**Example 3:**

```
Input:  pointer re-taken after grow
Output: true
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Reallocation on growth** | append past cap copies to a new array. |
| 2 | **Stale element pointer** | Old `&xs[0]` no longer aliases the slice. |
| 3 | **Re-address after append** | Re-take the pointer or index directly. |

## Hint

Re-take the address after the append: `p = &xs[0]; *p = 42` (or write `xs[0] = 42`).

## Validate

```bash
make verify
```
