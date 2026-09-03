# Bounding a Receive With a Deadline

## Intuition

A bare `<-scrapes` is an unbounded promise: it waits as long as the slowest
target takes. `select` turns that into a race between "a sample arrived" and
"the budget is gone", and whichever fires first decides the outcome.

The trap is *where* the timer lives. `time.After(budget)` evaluated inside the
loop creates a fresh timer per iteration, so the function's real bound becomes
"budget between consecutive samples" — a target that dribbles one sample every
19ms keeps a 20ms budget alive forever. Hoisting it out of the loop makes the
budget cover the whole collection.

## Approach

1. Return `[]Sample{}, true` early when `want <= 0`.
2. Create `deadline := time.After(budget)` once.
3. Loop while `len(collected) < want`, selecting on `scrapes` and `deadline`.
4. On a receive: if `!ok`, the pool closed — return what you have with `false`; else append.
5. On the deadline: return what you have with `false`.
6. After the loop, return the full set with `true`.

## Solution

```go
// CollectScrapes gathers up to want samples from the scrape channel, giving up
// when the budget expires or when the scrape pool closes the channel early.
// The bool reports whether the full set arrived; on a short read the samples
// gathered so far are still returned so the scrape can be recorded as partial.
//
// The timeout is a budget for the whole collection, not per sample.
//
// Examples:
//
//	CollectScrapes(chan 3 samples, 3, 5s) => 3 samples, true
//	CollectScrapes(chan 1 sample closed, 3, 5s) => 1 sample, false
//	CollectScrapes(silent chan, 2, 20ms) => no samples, false
func CollectScrapes(scrapes <-chan Sample, want int, budget time.Duration) ([]Sample, bool) {
	collected := []Sample{}
	if want <= 0 {
		return collected, true
	}

	deadline := time.After(budget)
	for len(collected) < want {
		select {
		case s, ok := <-scrapes:
			if !ok {
				return collected, false
			}
			collected = append(collected, s)
		case <-deadline:
			return collected, false
		}
	}
	return collected, true
}
```

## Walkthrough

- With three samples already queued and a five-second budget, the receive arm is
  ready on every pass and the deadline never fires: three samples, `true`.
- Asking for two of three queued samples stops at two and leaves the third in
  the channel — the collector takes only what it asked for.
- A closed channel makes `ok` false, so one sample comes back with `false`
  rather than the loop spinning on a closed channel forever.
- A silent target leaves only the deadline arm ready; after 20ms the function
  returns the empty slice and `false`.
- `want <= 0` never touches the channel at all.

## Pitfalls

- `case <-time.After(budget)` written inside the `select` resets the budget on
  every iteration.
- Dropping the comma-ok turns a closed channel into an infinite stream of zero
  `Sample` values, and the function returns `want` bogus samples with `true`.
- Returning `nil` instead of `[]Sample{}` on the empty paths makes callers
  distinguish nil from empty for no reason — and fails a `DeepEqual` test.
- Do not add a `default:` arm here: it would turn the blocking receive into a
  busy loop that burns a core while waiting for the budget.
