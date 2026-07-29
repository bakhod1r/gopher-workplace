# Deref or Zero

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _nil-pointer-dereference_

## Context

A very common helper: safely turn an optional pointer into a value, defaulting
to the zero value on nil.

## Task

Implement `DerefOrZero` in [orzero.go](orzero.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  DerefOrZero(nil)
Output: 0
```

_Explanation:_ Nil → zero, no dereference.

**Example 2:**

```
Input:  n := 8; DerefOrZero(&n)
Output: 8
```

**Example 3:**

```
Input:  n := 0; DerefOrZero(&n)
Output: 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Nil guard** | Check before deref. |
| 2 | **Zero default** | Return 0 on nil. |
| 3 | **Optional pointer** | Pointer models 'maybe'. |

## Hint

`if p == nil { return 0 }; return *p`.

## Validate

```bash
make verify
```
