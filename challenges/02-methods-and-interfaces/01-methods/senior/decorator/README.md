# Decorator Pattern

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A decorator adds behaviour around an existing object without editing it. It
holds the component, calls the component's method, and does something extra on
the way in or out.

## Task

Implement `DoWork` on `*Decorator` in [decorator.go](decorator.go):

1. Call `d.Comp.DoWork()`.
2. Return that result wrapped in square brackets.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  (&Decorator{Comp: &Component{}}).DoWork()
Output: "[work]"
```

**Example 2:**

```
Input:  the inner component returns "work"
Output: "[" + "work" + "]"
```

**Example 3:**

```
Input:  a decorator wrapping another decorator's result "[work]"
Output: "[[work]]"
```

_Explanation:_ the wrapping composes, which is the whole point of a decorator.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Same method name, outer type** | `Decorator.DoWork` shadows nothing — it is a distinct method on a distinct type. |
| 2 | **Delegation through a field** | `Comp` is a plain pointer field; the decorator must call it explicitly. |
| 3 | **String concatenation** | `"[" + s + "]"` is enough; no `fmt` needed. |

## Hint

One line: `return "[" + d.Comp.DoWork() + "]"`.

## Validate

```bash
make verify
```
