# Render Numbers Without Boxing Them

**Level:** middle
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A metrics exporter builds every line with `fmt.Sprintf`. The formatter's variadic `any` parameter forces each number onto the heap, once per metric, once per scrape.

## Task

Implement [appendints.go](appendints.go):

1. Append the decimal rendering of each value to `dst`, separated by a single space.
2. Return the extended slice.
3. With enough capacity in `dst`, the call must allocate nothing.

Replace the stub body in [appendints.go](appendints.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  AppendInts(nil, []int{1,2,3})
Output: "1 2 3"
```

**Example 2:**

```
Input:  AppendInts([]byte("x:"), []int{-4})
Output: "x:-4"
```

_Explanation:_ dst is extended, not replaced.

**Example 3:**

```
Input:  AppendInts([]byte("keep"), nil)
Output: "keep"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Interface boxing** | Storing an int in an `any` needs a heap word to point at. |
| 2 | **strconv.Append*** | Renders straight into a byte slice with no intermediate string. |
| 3 | **Caller-owned buffers** | An `Append`-style API lets the caller keep the memory. |

## Hint

`strconv` has an `Append` twin for every `Format` function.

## Validate

```bash
make verify
```
