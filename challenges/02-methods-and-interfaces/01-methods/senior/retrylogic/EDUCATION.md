# Retry Methods

## Intuition

A retry loop is a budget plus a stopping rule. The budget is `maxAttempts`; the
stopping rules are "it worked" and "the budget is gone". Everything else —
backoff, jitter, deciding which errors are retryable — layers on top of this
shape.

## Approach

1. Declare the error outside the loop so the last one survives.
2. Call the operation; return nil the moment it succeeds.
3. Fall out of the loop and return the error you kept.

## Solution

```go
func (c *Client) DoWithRetry(maxAttempts int) error {
	var err error
	for i := 0; i < maxAttempts; i++ {
		err = c.Do()
		if err == nil {
			return nil
		}
	}
	return err
}
```

## Walkthrough

With `FailInt: 2` and a budget of 3: attempt 1 fails, attempt 2 fails, attempt 3
succeeds and returns nil with `Attempts == 3`.

With `FailInt: 5` the loop uses all three attempts, each assigning `err`, and
returns the third failure. `Attempts == 3` proves the budget was respected — a
loop that kept going would show more.

## Pitfalls

- **Returning `err` inside the loop.** Turns three attempts into one; both
  `Attempts` assertions fail.
- **Declaring `err` inside the loop.** It goes out of scope, so the final return
  has nothing to report (or does not compile).
- **Looping `maxAttempts` times *after* a first call.** That is `maxAttempts+1`
  calls — an off-by-one that quietly doubles load on a struggling dependency.
- **Retrying everything.** A 400-class error will fail identically every time;
  real clients classify errors before retrying.

## Pair it with backoff

Retrying immediately is how a struggling service gets finished off. The
`backoff` puzzle in this subtopic supplies the delay this loop should wait
between attempts — retry budget and backoff policy are two halves of one
resilience story.
