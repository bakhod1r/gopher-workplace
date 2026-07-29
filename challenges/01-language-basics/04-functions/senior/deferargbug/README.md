# Defer Argument Timing

**Level:** senior
**Topic:** 01-language-basics → 04-functions · _defer_

## Context

`defer f(c)` evaluates `c` at the defer statement (when c==0), not at exit. To
read the final value, capture `c` in the closure body instead of passing it as
an argument.

## Task

Fix [deferargbug.go](deferargbug.go) so it records the FINAL counter value.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  FinalCount(5)
Output: 5
```

**Example 2:**

```
Input:  FinalCount(0)
Output: 0
```

**Example 3:**

```
Input:  FinalCount(3)
Output: 3
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Argument vs closure capture** | Arguments snapshot; body reads at run time. |
| 2 | **Defer evaluation time** | Args evaluate now, body runs at return. |
| 3 | **Named result** | `recorded` is set by the deferred call. |

## Hint

Drop the parameter and capture `c` in the closure body: `defer func(){ recorded = c }()`.

## Validate

```bash
make verify
```
