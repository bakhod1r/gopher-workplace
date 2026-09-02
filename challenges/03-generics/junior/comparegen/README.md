# Compare

**Level:** junior  
**Topic:** 03-generics

## Context

A sort helper and a diff tool both want a three-way comparison rather than a boolean `less`.

## Task

Implement the stub(s) in [comparegen.go](comparegen.go):

1. Implement `Compare`, returning `-1`, `0`, or `1` according to the ordering of `a` and `b`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Compare(1, 2)
Output: -1
```

**Example 2:**

```
Input:  Compare(2, 2)
Output: 0
```

**Example 3:**

```
Input:  Compare("b", "a")
Output: 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`cmp.Ordered`** | The stdlib constraint for types supporting `<`, `<=`, `>`, `>=`. |
| 2 | **Three-way results** | A single `int` carries strictly more information than a `less` bool. |
| 3 | **`switch` without a subject** | Reused from language basics: `switch { case cond: }` reads better than nested ifs. |

## Hint

A bare `switch` with two cases and a default.

## Validate

```bash
make verify
```
