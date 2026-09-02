# Chain of Responsibility

## Intuition

Every handler answers one question: *is this mine?* If yes, it produces the
result and the chain stops. If no, it forwards — and forwarding is identical for
every link, so it lives once in `BaseHandler`.

Embedding is what makes that sharing free: `H20` declares `BaseHandler` with no
field name, so `h.Next(...)` and `h.SetNext(...)` resolve to the embedded
value's methods.

## Approach

1. Test the link's own condition.
2. Return its own answer on a match.
3. Delegate to the promoted `Next` otherwise.

## Solution

```go
func (h *H20) Handle(req int) string {
	if req == 20 {
		return "twenty"
	}
	return h.Next(req)
}
```

## Walkthrough

`h10.SetNext(h20)` writes `h20` into `h10`'s embedded `BaseHandler.next`.

- `h10.Handle(10)` matches immediately: `"ten"`.
- `h10.Handle(20)` does not match, so `h10.Next(20)` sees a non-nil `next` and
  calls `h20.Handle(20)`, which matches: `"twenty"`.
- `h10.Handle(30)` falls through to `h20.Handle(30)`, which also misses, and
  `h20.next` is nil, so `BaseHandler.Next` returns `"unhandled"`.

## Pitfalls

- **Calling `h.next.Handle(req)` directly.** That is a nil-pointer panic on the
  last link; `Next` exists to guard it.
- **Returning `"unhandled"` yourself.** The fallback belongs to the base so the
  chain length stays a runtime detail.
- **Value receiver on `Handle`.** `*H20` would no longer satisfy `Handler`
  consistently with `SetNext`, which needs a pointer to mutate `next`.
