# Between

**Level:** junior  
**Topic:** 03-generics

## Context

A validation rule accepts values inside an inclusive range — ages, prices, and version strings all use it.

## Task

Implement the stub(s) in [betweengen.go](betweengen.go):

1. Implement `Between`, reporting whether `v` lies in the inclusive range `[lo, hi]`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Between(2, 1, 3)
Output: true
```

**Example 2:**

```
Input:  Between(1, 1, 3)
Output: true
```

**Example 3:**

```
Input:  Between(4, 1, 3)
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`cmp.Ordered`** | The stdlib constraint for types supporting `<`, `<=`, `>`, `>=`. |
| 2 | **Inclusive bounds** | `>=` and `<=` include the endpoints; `>` and `<` would exclude them. |
| 3 | **Short-circuit `&&`** | Reused from language basics: the second comparison is skipped when the first fails. |

## Hint

One expression: `v >= lo && v <= hi`.

## Validate

```bash
make verify
```
