# Template Method

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

The *order* of the steps is fixed; the steps themselves vary. `Template.Run`
owns the algorithm's skeleton and defers each step to whatever `Step`
implementation it was given.

## Task

Implement `Run` on `*Template` in [templatemeth.go](templatemeth.go):

1. Call `t.impl.DoStep1()` and `t.impl.DoStep2()`, in that order.
2. Return the two results joined with `"-"`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Template{impl: MyTask{}}.Run()
Output: "a-b"
```

**Example 2:**

```
Input:  DoStep1() = "a"
Output: appears before the separator
```

**Example 3:**

```
Input:  an implementation whose steps return "x" and "y"
Output: "x-y"
```

_Explanation:_ `Run` never changes — only the injected `Step` does.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Interface field** | `impl Step` is the substitution point; Go composes instead of inheriting. |
| 2 | **Fixed skeleton** | `Run` encodes order and joining; subclasses cannot reorder it. |
| 3 | **Value receiver satisfies the interface** | `MyTask{}` (not `&MyTask{}`) is assignable to `Step`. |

## Hint

One line: `return t.impl.DoStep1() + "-" + t.impl.DoStep2()`.

## Validate

```bash
make verify
```
