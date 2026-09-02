# Min Of Two

**Level:** junior  
**Topic:** 03-generics

## Context

A scheduler takes the earliest of two deadlines, whatever unit the caller uses to express them.

## Task

Implement the stub(s) in [minof.go](minof.go):

1. Implement `Min`, returning the smaller of `a` and `b`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Min(2, 5)
Output: 2
```

**Example 2:**

```
Input:  Min("b", "a")
Output: "a"
```

**Example 3:**

```
Input:  Min(3.5, 1.5)
Output: 1.5
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`cmp.Ordered`** | The stdlib constraint for types supporting `<`, `<=`, `>`, `>=`. |
| 2 | **Constraints permit operations** | A type parameter can only do what every type in its set can do — that is why `+` needs a numeric constraint. |
| 3 | **Mirror functions** | `Min` and `Max` differ only in the comparison direction. |

## Hint

One `if`, comparison flipped from `Max`.

## Validate

```bash
make verify
```
