# Immediately-Invoked Init

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

An immediately-invoked function literal must be CALLED with a trailing `()` to
produce its result. The bug stores the function itself and returns nil; add the
call and return the built map.

## Task

Fix [iifeinit.go](iifeinit.go) so the table is built and returned.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  BuildTable(4)
Output: {0:0, 1:1, 2:4, 3:9}
```

**Example 2:**

```
Input:  BuildTable(1)
Output: {0:0}
```

**Example 3:**

```
Input:  BuildTable(0)
Output: {}
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Immediately-invoked function** | Append `()` to call it now. |
| 2 | **Function value vs call** | Without `()` you hold the func, not its result. |
| 3 | **Scoped initialisation** | IIFEs bundle setup logic into one expression. |

## Hint

Invoke the literal and return its result: `return func() map[int]int { ... }()` (note the trailing `()`).

## Validate

```bash
make verify
```
