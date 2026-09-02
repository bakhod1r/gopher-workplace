# Max Of Two

**Level:** junior  
**Topic:** 03-generics

## Context

A rate limiter picks the stricter of two limits. The same comparison serves ints, floats, and version strings.

## Task

Implement the stub(s) in [maxof.go](maxof.go):

1. Implement `Max`, returning the larger of `a` and `b`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Max(2, 5)
Output: 5
```

**Example 2:**

```
Input:  Max("a", "b")
Output: "b"
```

**Example 3:**

```
Input:  Max(3.5, 1.5)
Output: 3.5
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`cmp.Ordered`** | The stdlib constraint for types supporting `<`, `<=`, `>`, `>=`. |
| 2 | **Constraints permit operations** | A type parameter can only do what every type in its set can do — that is why `+` needs a numeric constraint. |
| 3 | **Importing `cmp`** | `cmp.Ordered` comes from the stdlib `cmp` package — no need to hand-write the union. |

## Hint

`comparable` gives you `==` only; ordering needs `cmp.Ordered`.

## Validate

```bash
make verify
```
