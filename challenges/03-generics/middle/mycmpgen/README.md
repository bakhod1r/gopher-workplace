# Writing Your Own Ordered

**Level:** middle  
**Topic:** 03-generics

## Context

A package predating the `cmp` import wrote its own ordering constraint. Understanding what it must contain explains why the stdlib version looks the way it does.

## Task

Implement the stub(s) in [mycmpgen.go](mycmpgen.go):

1. Implement `Largest` against the hand-written `Ordered` constraint.
2. Compare the constraint with `cmp.Ordered`: same shape, and the same NaN caveat.
3. On a tie keep the earlier element.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Largest([]int{1, 9, 3})
Output: 9, true
```

**Example 2:**

```
Input:  Largest([]string{"a", "c"})
Output: "c", true
```

**Example 3:**

```
Input:  Largest([]float64{})
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **What `Ordered` must list** | Every integer kind, both floats, and `~string` — complex types have no `<`. |
| 2 | **Why `~` is required** | Without it, named types like `Celsius` would be excluded. |
| 3 | **Prefer the stdlib** | `cmp.Ordered` is the same set, maintained centrally — hand-rolling it is now a code smell. |

## Hint

The body is an ordinary scan; the lesson is in what the constraint has to enumerate.

## Validate

```bash
make verify
```
