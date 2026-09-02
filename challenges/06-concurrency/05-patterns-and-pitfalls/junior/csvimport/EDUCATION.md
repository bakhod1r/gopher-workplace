# Chaining Three Stages

## Intuition

Pipelines scale by repetition, not by new concepts. A third stage is the same
shape as the second, and the shutdown still cascades from the source to the
sink because every stage closes its own output when its input dries up.

## Approach

1. Reader goroutine: send each row on `src`, `defer close(src)`.
2. Cleaning goroutine: range `src`, send `clean(row)` on `cleaned`, `defer close(cleaned)`.
3. Validation goroutine: range `cleaned`, send rows passing `valid` on `accepted`; the caller ranges over `accepted` and appends.

## Solution

```go
func ImportCSV(rows []string, clean func(string) string, valid func(string) bool) []string {
	src := make(chan string)
	go func() {
		defer close(src)
		for _, row := range rows {
			src <- row
		}
	}()

	cleaned := make(chan string)
	go func() {
		defer close(cleaned)
		for row := range src {
			cleaned <- clean(row)
		}
	}()

	accepted := make(chan string)
	go func() {
		defer close(accepted)
		for row := range cleaned {
			if valid(row) {
				accepted <- row
			}
		}
	}()

	var out []string
	for row := range accepted {
		out = append(out, row)
	}
	return out
}
```

## Walkthrough

For `" z", "", "a "`: cleaning yields `z`, ``, `a`; validation drops the empty
string and forwards `z` and `a`. When the reader closes `src`, each stage in
turn finishes its range and closes its own output, so the caller's loop ends
with `[z a]`.

## Pitfalls

- Letting a stage close a channel it did not create, which panics the stage still sending on it.
- Abandoning the final channel without draining it — the validation goroutine blocks on send forever.
- Collecting results inside a goroutine instead of the caller, which needs synchronisation the pipeline shape was meant to avoid.
