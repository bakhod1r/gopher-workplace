# Parallel Map

**Level:** staff
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A CPU-bound transform ran on one core while the box had sixteen. Parallelising it must actually use the cores, not just add goroutines.

## Task

Implement the stub(s) in [parallelmap.go](parallelmap.go):

1. Implement `Apply` on `SquareOp`.
2. Implement `MapParallel`, which applies the op across `GOMAXPROCS` workers over contiguous chunks, preserving index order.
3. Constraint: `-race` clean, results identical to the sequential version, and the workers must genuinely run concurrently — the test observes peak parallelism when `GOMAXPROCS > 1`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  MapParallel(op, [1 2 3])
Output: [1 4 9]
```

**Example 2:**

```
Input:  an input smaller than the worker count
Output: still correct
```

**Example 3:**

```
Input:  an empty input
Output: empty
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **GOMAXPROCS** | Parallelism is bounded by the number of Ps, not by goroutine count. |
| 2 | **Chunking** | Contiguous ranges give each worker cache-friendly work and disjoint output slots. |
| 3 | **Order preservation** | Reused: index-addressed writes need no lock and keep order. |

## Hint

Split `[0, n)` into `workers` contiguous chunks; each worker writes only its own range.

## Validate

```bash
make verify
```
