# Taking a Prefix Without Leaking

## Intuition

Every unbuffered send needs a matching receive. Walking away from a stream
leaves its producer parked on a send that will never complete. Draining is the
simplest fix at this level: read the rest and throw it away.

## Approach

1. Handle `n <= 0` by draining and returning nil.
2. Range over `hits`, appending until `len(out) == n`, then `break`.
3. Drain the remainder with `for range hits {}` and return the prefix.

## Solution

```go
func TopResults(hits <-chan string, n int) []string {
	if n <= 0 {
		for range hits {
		}
		return nil
	}

	var out []string
	for hit := range hits {
		out = append(out, hit)
		if len(out) == n {
			break
		}
	}

	for range hits {
	}
	return out
}
```

## Walkthrough

For five hits with `n = 3`: the first loop appends a, b, c and breaks. The
producer is at that moment blocked trying to send d. The drain loop receives
d and e, the producer's own loop ends, it closes the channel, and the drain
loop finishes with it.

## Pitfalls

- Returning straight after the `break` — the producer leaks, one goroutine per search.
- Draining before taking, which consumes the hits you wanted.
- Forgetting the `n <= 0` guard: `len(out) == n` never matches, so the loop takes everything.
