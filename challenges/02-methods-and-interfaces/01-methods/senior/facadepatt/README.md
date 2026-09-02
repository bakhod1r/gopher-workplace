# Facade Pattern

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

Two subsystems each expose their own call. Callers should not have to know
about both, or in what order to use them. A facade offers one method that
orchestrates the pieces behind it.

## Task

Implement `Operation` on `*Facade` in [facadepatt.go](facadepatt.go):

1. Call `f.s1.Op1()` and `f.s2.Op2()`.
2. Return the two results joined with `"+"`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  (&Facade{}).Operation()
Output: "1+2"
```

**Example 2:**

```
Input:  Op1() = "1"
Output: contributes the "1" before the separator
```

**Example 3:**

```
Input:  Op2() = "2"
Output: contributes the "2" after the separator
```

_Explanation:_ the facade's only job is ordering and joining the subsystem calls.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Zero-value usability** | `Sub1`/`Sub2` are empty structs, so `&Facade{}` is ready to use without a constructor. |
| 2 | **Value receivers via a field** | `Op1` has a value receiver; calling it through `f.s1` needs no pointer. |
| 3 | **Orchestration in one method** | The subsystems stay unaware of each other. |

## Hint

One line: `return f.s1.Op1() + "+" + f.s2.Op2()`.

## Validate

```bash
make verify
```
