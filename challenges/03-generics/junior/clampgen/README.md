# Clamp

**Level:** junior  
**Topic:** 03-generics

## Context

A slider widget must never report a value outside its configured bounds, whether the value is an int step or a float ratio.

## Task

Implement the stub(s) in [clampgen.go](clampgen.go):

1. Implement `Clamp`, returning `lo` when `v` is below the range, `hi` when above, and `v` otherwise.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Clamp(5, 0, 3)
Output: 3
```

**Example 2:**

```
Input:  Clamp(-1, 0, 3)
Output: 0
```

**Example 3:**

```
Input:  Clamp(2, 0, 3)
Output: 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`cmp.Ordered`** | The stdlib constraint for types supporting `<`, `<=`, `>`, `>=`. |
| 2 | **Ordered guards** | Two comparisons fully define the clamp — no arithmetic needed. |
| 3 | **Constraints permit operations** | A type parameter can only do what every type in its set can do — that is why `+` needs a numeric constraint. |

## Hint

Check the low bound first, then the high bound, then fall through.

## Validate

```bash
make verify
```
