# Reading the Request Budget

## Intuition

`Deadline()` is a read-only query on the context chain. The `ok` half distinguishes "the deadline happens to be the zero time" from "there is no deadline", the same comma-ok shape as a map lookup. `WithCancel` and `WithValue` add no deadline, so children hand back the ancestor's answer unchanged.

## Approach

1. Return `ctx.Deadline()` directly.

## Solution

```go
// Budget reports the deadline the caller imposed on this request and whether
// one exists at all. The object-storage client calls it before an upload to
// decide between a single PUT and a slower multipart transfer.
//
// Examples:
//
//	Budget(context.Background())            => zero time, false
//	Budget(ctx with deadline t)             => t, true
//	Budget(context.WithValue(bg, k, v))     => zero time, false
func Budget(ctx context.Context) (time.Time, bool) {
	return ctx.Deadline()
}
```

## Walkthrough

- `context.Background()` has no deadline, so it returns `(time.Time{}, false)`.
- `WithDeadline(bg, at)` stores `at` and returns `(at, true)`.
- `WithCancel(withDeadline)` adds cancellation but no new deadline, so it reports its parent's `at` — the upload client sees the same budget however many layers wrapped the context.

## Pitfalls

- Ignoring the boolean and testing `got.IsZero()` conflates "no deadline" with a zero-valued one.
- Compare `time.Time` values with `Equal`, not `==`, because of the monotonic-clock and location fields.
