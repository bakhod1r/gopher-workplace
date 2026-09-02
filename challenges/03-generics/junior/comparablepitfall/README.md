# What comparable Really Means

**Level:** junior  
**Topic:** 03-generics

## Context

Someone tried to use this helper with a struct holding a slice field and got a compile error they did not expect.

## Task

Implement the stub(s) in [comparablepitfall.go](comparablepitfall.go):

1. Implement `CountDistinct`, returning the number of distinct values.
2. Understand which types satisfy `comparable`: structs of comparable fields do; anything containing a slice, map, or function does not.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  CountDistinct([]int{1, 1, 2})
Output: 2
```

**Example 2:**

```
Input:  CountDistinct([]Point{{1,2},{1,2}})
Output: 1
```

**Example 3:**

```
Input:  CountDistinct([]int{})
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Comparable structs** | A struct is comparable when every field is — `Point{X, Y int}` qualifies. |
| 2 | **What is excluded** | Slices, maps, and functions are not comparable, so structs containing them are not either. |
| 3 | **A compile-time check** | The constraint is enforced at instantiation, not at run time. |

## Hint

Structs of comparable fields are themselves comparable — that is why `[]Point` works.

## Validate

```bash
make verify
```
