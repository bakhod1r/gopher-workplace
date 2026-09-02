# The Sum That Allocates Per Element

**Level:** staff  
**Topic:** 03-generics

## Context

A metrics aggregator sums a few million samples per scrape. The numbers are right, but the process spends most of its time in the garbage collector and the heap profile is dominated by this one function.

## Task

Fix the single planted bug in [anyboxhotloopbug.go](anyboxhotloopbug.go):

1. Find and fix the single bug so the sum is computed without boxing each element.
2. The result must not change for any input.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Sum([]int{1,2,3})
Output: 6
```

**Example 2:**

```
Input:  Sum([]float64{1.5,2.5})
Output: 4
```

**Example 3:**

```
Input:  Sum([]int{})
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Boxing into `any`** | Storing a non-pointer value in an interface copies it to the heap; in a loop that is one allocation per element. |
| 2 | **Generics are not `any`** | A type parameter is resolved at compile time — routing values through `interface{}` throws that away. |
| 3 | **Allocation as a test target** | `testing.AllocsPerRun` turns "is it fast" into a deterministic assertion. |

## Hint

Count the heap allocations for a slice of a thousand large integers.

## Validate

```bash
make verify
```
