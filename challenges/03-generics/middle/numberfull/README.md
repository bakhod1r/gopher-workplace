# The Whole Numeric Set

**Level:** middle  
**Topic:** 03-generics

## Context

A metrics layer stores counters as `uint64`, deltas as `int`, and ratios as `float64`. One summing helper must cover all three.

## Task

Implement the stub(s) in [numberfull.go](numberfull.go):

1. Implement `Total`, summing every element.
2. Study the `Number` constraint: it deliberately includes unsigned types, so nothing in the body may assume negatives exist.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Total([]int{1, -2})
Output: -1
```

**Example 2:**

```
Input:  Total([]uint{1, 2})
Output: 3
```

**Example 3:**

```
Input:  Total([]float64{0.5, 0.5})
Output: 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Type sets** | A constraint's type set is the exact list of types an instantiation may use. |
| 2 | **Unsigned members restrict you** | With `~uint` in the set, `v < 0` is dead code and `-v` wraps. |
| 3 | **Operations follow the set** | A type parameter supports only what every member of its set supports. |

## Hint

Addition is safe for every member; comparison against zero is not meaningful for all of them.

## Validate

```bash
make verify
```
