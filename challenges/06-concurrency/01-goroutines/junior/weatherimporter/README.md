# Weather Importer

**Level:** junior
**Topic:** 06-concurrency → 01-goroutines

## Context

A weather ingest pipeline receives station readings in Fahrenheit and stores
them in Celsius. Readings are independent values in a batch, so the conversion
fans out across goroutines and the batch keeps its station order.

## Task

Implement `ToCelsius` in [weatherimporter.go](weatherimporter.go) so that:

1. Return a new slice the same length as `readings`; do not modify the input.
2. Reading `i` becomes `(readings[i] - 32) * 5 / 9`.
3. Convert each reading in its own goroutine, joined with a `sync.WaitGroup`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ToCelsius([]float64{32, 212})
Output: [0 100]
```

**Example 2:**

```
Input:  ToCelsius([]float64{-40})
Output: [-40]
```

**Example 3:**

```
Input:  ToCelsius(nil)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`go` statement** | `go f(x)` starts a goroutine; the caller keeps running and does not wait. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before each launch, `defer wg.Done()` inside, `wg.Wait()` in the parent. |
| 3 | **Loop-variable capture** | Pass the index and the element in as goroutine parameters instead of reading the loop variable. |
| 4 | **Any element type works** | `make([]float64, n)` zero-fills before the fan-out; the pattern does not care about the type. |

## Hint

Keep the arithmetic in `float64` throughout — subtract first, then scale.

## Validate

```bash
make verify
```
