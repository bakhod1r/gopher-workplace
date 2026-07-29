# Defer Overwrites Result

**Level:** senior
**Topic:** 01-language-basics → 04-functions · _defer_

## Context

A deferred closure runs AFTER the return value is set and can overwrite it.
Here the defer zeroes `result` just before the caller receives it. Remove (or
correct) the deferred mutation.

## Task

Fix [deferwipe.go](deferwipe.go) so `Compute` returns `a*b`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Compute(6, 7)
Output: 42
```

**Example 2:**

```
Input:  Compute(0, 5)
Output: 0
```

**Example 3:**

```
Input:  Compute(3, 3)
Output: 9
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Defer runs post-return** | It sees and can change the named result. |
| 2 | **Unintended mutation** | Zeroing `result` discards the computation. |
| 3 | **Order of effects** | Body sets result, then defer overwrites it. |

## Hint

Remove the deferred reset (it serves no purpose), or change it so it doesn't clobber `result`.

## Validate

```bash
make verify
```
