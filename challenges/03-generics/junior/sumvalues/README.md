# Sum Map Values

**Level:** junior  
**Topic:** 03-generics

## Context

A metrics map holds one counter per label. The overall total ignores the labels entirely.

## Task

Implement the stub(s) in [sumvalues.go](sumvalues.go):

1. Implement `SumValues`, returning the sum of the map's values.
2. Return the zero value for an empty map.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SumValues(map[string]int{"a": 1, "b": 2})
Output: 3
```

**Example 2:**

```
Input:  SumValues(map[int]float64{1: 0.5, 2: 0.5})
Output: 1
```

**Example 3:**

```
Input:  SumValues(map[string]int{})
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Two constraints, two jobs** | `K comparable` satisfies the map key rule; `V Number` enables `+`. |
| 2 | **Constraints permit operations** | A type parameter can only do what every type in its set can do — that is why `+` needs a numeric constraint. |
| 3 | **Order independence** | Addition is commutative, so randomised map order does not change the result. |

## Hint

Two type parameters with different constraints — pick each for the operation it enables.

## Validate

```bash
make verify
```
