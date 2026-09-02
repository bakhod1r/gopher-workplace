# Counting a Stream

## Intuition

Anything you can compute incrementally should be computed as the values
arrive. Counting matches needs one integer of state, whatever the log's
length — this is what makes channel consumers scale to unbounded input.

## Approach

1. Start `n` at 0.
2. `range` over `lines`.
3. Increment `n` when `line == want`.
4. Return `n`.

## Solution

```go
func CountStatus(lines <-chan string, want string) int {
	n := 0
	for line := range lines {
		if line == want {
			n++
		}
	}
	return n
}
```

## Walkthrough

For `"200"`, `"500"`, `"200"` with `want == "200"`: matches on the first
and third lines bring `n` to 2; `"500"` is skipped.

## Pitfalls

- Collecting the log into a slice and counting afterwards wastes memory for no gain.
- Comparing with `strings.EqualFold` would make the match case-insensitive — not what is asked.
- The loop still requires the tailer to close, or the analyser never reports.
