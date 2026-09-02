# Weather Importer

## Intuition

The element type changes nothing about the concurrency: preallocate, write by
index, wait. That shape is type-agnostic, which is why it is worth learning
once.

## Approach

1. Allocate `out := make([]float64, len(readings))`.
2. Launch one goroutine per reading, passing `i` and the value.
3. Write the converted value to `out[i]`.
4. `wg.Wait()` before returning.

## Solution

```go
// Package weatherimporter — Gopher Workplace challenge.
package weatherimporter

import (
	"sync"
)

// ToCelsius converts every Fahrenheit reading to Celsius.
//
// Examples:
//
//	ToCelsius([]float64{32, 212})  => [0 100]
//	ToCelsius([]float64{-40})      => [-40]
//	ToCelsius(nil)                 => []
func ToCelsius(readings []float64) []float64 {
	out := make([]float64, len(readings))
	var wg sync.WaitGroup
	for i, f := range readings {
		wg.Add(1)
		go func(i int, f float64) {
			defer wg.Done()
			out[i] = (f - 32) * 5 / 9
		}(i, f)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

- `32 °F` converts to `(32-32) * 5 / 9 = 0 °C`.
- `212 °F` converts to `180 * 5 / 9 = 100 °C`.
- `-40` is the fixed point of the two scales and converts to `-40`.

## Pitfalls

- Writing `f - 32 * 5 / 9` — precedence applies the multiplication first.
- Using integer arithmetic anywhere in the expression and truncating the result.
- Converting `readings` in place, mutating the caller's batch.
