# Histogram Buckets

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Binning values into a histogram — latency buckets, size distributions.

## Task

Implement `Bucket(xs, size)`: count per `[i*size, (i+1)*size)` bin.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [5,12,15,25,3], size=10
Output: [2,2,1]
```

_Explanation:_ bins [0,10)=2, [10,20)=2, [20,30)=1

**Example 2:**

```
Input:  nil, size=10
Output: []
```

**Example 3:**

```
Input:  [0], size=5
Output: [1]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Bin index** | `v / size`. |
| 2 | **Size the result** | Enough bins for the max value. |
| 3 | **Allocate + count** | `make` then `bins[v/size]++`. |

## Hint

Find `maxBin = max(v)/size`; `make([]int, maxBin+1)`; `bins[v/size]++`.

## Validate

```bash
make verify
```
