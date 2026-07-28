# Histogram Buckets

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Binning values into a histogram — latency buckets, size distributions.

## Task

Implement `Bucket(xs, size)`: count per `[i*size, (i+1)*size)` bin.

## Examples

```go
Bucket([]int{5,12,15,25,3}, 10) // => [2 2 1]
```

## Topics to Master

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
