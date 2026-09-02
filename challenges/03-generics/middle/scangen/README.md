# Scan

**Level:** middle  
**Topic:** 03-generics

## Context

A balance chart needs the running total after every transaction, not just the final figure.

## Task

Implement the stub(s) in [scangen.go](scangen.go):

1. Implement `Scan`, returning the accumulator's value after each element.
2. The result has exactly `len(s)` elements — `init` itself is not included.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Scan([]int{1,2,3}, 0, add)
Output: []int{1,3,6}
```

**Example 2:**

```
Input:  Scan([]int{}, 5, add)
Output: []int{}
```

**Example 3:**

```
Input:  Scan([]int{2}, 1, mul)
Output: []int{2}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Scan versus Reduce** | `Reduce` keeps the last value; `Scan` keeps them all. |
| 2 | **Result length** | One output per input — deciding whether to include `init` is an API choice; here it is excluded. |
| 3 | **Higher-order generic functions** | A type parameter may appear in a function-typed parameter or return value. |

## Hint

Same fold as `Reduce`, but append after every step.

## Validate

```bash
make verify
```
