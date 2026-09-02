# Notifier

## Intuition

Fan-out delivery only needs one method. Failure is data (a bool), so one bad channel does not stop the others.

## Approach

1. `(*Email).Notify` appends to `Sent`, then returns `true`.
2. `Broken.Notify` returns `false`.
3. `Broadcast` loops over all notifiers, incrementing on success, and returns the count.

## Solution

```go
func (e *Email) Notify(msg string) bool {
	e.Sent = append(e.Sent, msg)
	return true
}

func (b Broken) Notify(msg string) bool { return false }

func Broadcast(ns []Notifier, msg string) int {
	ok := 0
	for _, n := range ns {
		if n.Notify(msg) {
			ok++
		}
	}
	return ok
}
```

## Walkthrough

With `{a, Broken{}, b}` the loop calls all three: `a` succeeds, `Broken` fails, `b` succeeds — result 2, and `b.Sent` proves the loop did not stop at the failure.

## Pitfalls

- Returning early when a notifier fails — later channels never get the message.
- Value receiver on `Email.Notify`, so `Sent` stays empty.
- Counting calls instead of successes.
