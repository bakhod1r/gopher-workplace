# Flat Map

**Level:** middle  
**Topic:** 03-generics

## Context

Each order expands into several shipment lines. The report wants one flat list of lines.

## Task

Implement the stub(s) in [flatmapgen.go](flatmapgen.go):

1. Implement `FlatMap`, applying `f` to each element and concatenating the results in order.
2. An element mapping to an empty slice contributes nothing.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  FlatMap([]int{1,2}, dup)
Output: []int{1,1,2,2}
```

**Example 2:**

```
Input:  FlatMap([]int{1}, none)
Output: []int{}
```

**Example 3:**

```
Input:  FlatMap([]int{}, dup)
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Map plus Flatten** | One pass does the work of both, with no intermediate `[][]U`. |
| 2 | **Variadic append** | `append(out, f(v)...)` splices each result in place. |
| 3 | **Higher-order generic functions** | A type parameter may appear in a function-typed parameter or return value. |

## Hint

`append(out, f(v)...)` inside the loop.

## Validate

```bash
make verify
```
