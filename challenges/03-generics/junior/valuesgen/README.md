# Map Values

**Level:** junior  
**Topic:** 03-generics

## Context

A metrics exporter sums the recorded counters; the label keys are irrelevant to the aggregation step.

## Task

Implement the stub(s) in [valuesgen.go](valuesgen.go):

1. Implement `Values`, returning a slice of the map's values.
2. Any order is acceptable; the tests sort before comparing.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Values(map[string]int{"a": 1})
Output: []int{1}
```

**Example 2:**

```
Input:  Values(map[string]int{"a": 1, "b": 2})
Output: []int{1, 2} (any order)
```

**Example 3:**

```
Input:  Values(map[string]int{})
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Maps with type parameters** | A map key type must be comparable, so `K` needs the `comparable` constraint. |
| 2 | **Blank identifier in `range`** | `for _, v := range m` discards the key you do not need. |
| 3 | **Type parameters** | `[T any]` declares a type parameter; the caller (or inference) picks `T`. |

## Hint

The result element type is `V`, so allocate `[]V`.

## Validate

```bash
make verify
```
