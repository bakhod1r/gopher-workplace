# An Address Is Not A Reference

**Level:** senior
**Topic:** 10-advanced-topics / 04-unsafe-package

## Context

A struct walker computes field addresses as integers. It works for months, then starts returning zeros on a build with a different inliner — the kind of bug that is never reproduced on demand.

## Task

Fix the single planted bug in [uintptrhold.go](uintptrhold.go):

1. Return `p.B`, reached through the field's offset rather than the selector.
2. Fix the single bug: the arithmetic must never leave `unsafe.Pointer`.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  SecondWord(&Pair{A:1, B:2})
Output: 2
```

**Example 2:**

```
Input:  SecondWord(&Pair{})
Output: 0
```

**Example 3:**

```
Input:  2000 iterations with GC
Output: every read correct
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **uintptr is not a pointer** | The collector does not see it, so the object it names can be freed or moved. |
| 2 | **unsafe.Add** | The supported way to offset a pointer; the result stays a pointer. |
| 3 | **The valid uintptr pattern** | Converting to `uintptr` and back is only defined within a single expression. |
| 4 | **runtime.KeepAlive** | Needed when an object's last use is through an address the collector cannot follow. |

## Hint

Two statements is one statement too many. What is the pointer during the second one?

## Validate

```bash
make verify
```
