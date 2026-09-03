# Naming The Sub-Benchmarks

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

One benchmark per input size is how you see a curve instead of a point. `b.Run` builds each sub-benchmark's full name by joining the parent name and the sub-name with `/` — that joined name is what `-bench` matches against and what benchstat groups by.

## Task

Implement `Names` in [benchsubrun.go](benchsubrun.go):

1. For each size produce `"<base>/size=<size>"`.
2. Keep the input order.
3. No sizes yields an empty, non-nil slice.

## Examples

**Example 1:**

```
Input:  Names("BenchmarkEncode", []int{1, 10})
Output: ["BenchmarkEncode/size=1" "BenchmarkEncode/size=10"]
```

**Example 2:**

```
Input:  Names("B", []int{100, 1})
Output: ["B/size=100" "B/size=1"]
```

**Example 3:**

```
Input:  Names("B", nil)
Output: [] (non-nil)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`b.Run` builds a path** | Parent and sub-name join with `/`, and `-bench` matches the whole path. |
| 2 | **Size in the name** | Encoding the parameter makes results comparable across runs. |
| 3 | **Non-nil empty results** | Callers ranging over the result should not have to nil-check. |

## Topics used again

`make` with a capacity hint, `fmt.Sprintf`.

## Hint

Preallocate with `make([]string, 0, len(sizes))` — that also gives you the non-nil empty case for free.

## Validate

```bash
make verify
```
