# Map Keys

**Level:** junior  
**Topic:** 03-generics

## Context

A config auditor lists which settings were provided, without caring about their values or types.

## Task

Implement the stub(s) in [keysgen.go](keysgen.go):

1. Implement `Keys`, returning a slice of the map's keys.
2. Any order is acceptable; the tests sort before comparing.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Keys(map[string]int{"a": 1})
Output: []string{"a"}
```

**Example 2:**

```
Input:  Keys(map[int]bool{1: true, 2: false})
Output: []int{1, 2} (any order)
```

**Example 3:**

```
Input:  Keys(map[string]int{})
Output: []string{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Two constrained parameters** | `K comparable` is required by the map type; `V any` because values are never touched. |
| 2 | **Map iteration order** | Reused from language basics: Go randomises map range order deliberately. |
| 3 | **Maps with type parameters** | A map key type must be comparable, so `K` needs the `comparable` constraint. |

## Hint

`for k := range m` yields keys only — no second variable needed.

## Validate

```bash
make verify
```
