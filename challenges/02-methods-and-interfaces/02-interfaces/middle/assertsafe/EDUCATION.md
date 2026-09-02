# Safe Assertion

## Intuition

The bus stays type-agnostic; each handler decides for itself. The comma-ok assertion is the gate, and rejecting must leave state untouched.

## Approach

1. Assert the payload to the handler's type with comma-ok.
2. Return `false` immediately when the assertion fails.
3. Only then update state and return `true`.
4. `Dispatch` offers the payload to all handlers and counts the `true` results.

## Solution

```go
func (h *IntHandler) Handle(payload any) bool {
	n, ok := payload.(int)
	if !ok {
		return false
	}
	h.Sum += n
	return true
}

func (h *TextHandler) Handle(payload any) bool {
	s, ok := payload.(string)
	if !ok {
		return false
	}
	h.Seen = append(h.Seen, s)
	return true
}

func Dispatch(hs []Handler, payload any) int {
	n := 0
	for _, h := range hs {
		if h.Handle(payload) {
			n++
		}
	}
	return n
}
```

## Walkthrough

`Dispatch(hs, 1.5)`: both assertions fail, both handlers return `false`, the count is 0 and no state changes — which the final assertion in the test verifies.

## Pitfalls

- Mutating before checking `ok`, which adds a zero value on every rejected payload.
- Stopping the dispatch loop at the first acceptance.
- Using a value receiver, so accepted payloads never accumulate.
