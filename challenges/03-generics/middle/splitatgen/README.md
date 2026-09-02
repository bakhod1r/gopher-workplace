# Split At

**Level:** middle  
**Topic:** 03-generics

## Context

A paginator cuts a result set at an offset that arrives from a URL and cannot be trusted.

## Task

Implement the stub(s) in [splitatgen.go](splitatgen.go):

1. Implement `SplitAt`, returning the elements before and from index `i`.
2. Clamp `i` into `[0, len(s)]` rather than panicking.
3. Both halves must be independent of the input.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SplitAt([]int{1,2,3}, 1)
Output: []int{1}, []int{2,3}
```

**Example 2:**

```
Input:  SplitAt([]int{1,2}, 9)
Output: []int{1,2}, []int{}
```

**Example 3:**

```
Input:  SplitAt([]int{1,2}, -1)
Output: []int{}, []int{1,2}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Clamping untrusted input** | Two comparisons turn a panicking cut into a total function. |
| 2 | **No aliasing** | Return fresh slices; sub-slices of the input share its backing array. |
| 3 | **Independent halves** | Copying stops an append into one half from touching the other. |

## Hint

Clamp first, then copy — two `make`+`copy` pairs.

## Validate

```bash
make verify
```
