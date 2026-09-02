# Func Adapter

## Intuition

Any named type can have methods — including a function type. That one trick lets a bare function satisfy an interface, which is exactly how `http.HandlerFunc` works.

## Approach

1. `Handle` calls `f(s)`; the receiver is the function itself.
2. `Run` returns `h.Handle(s)`.
3. `Chain` returns a `HandlerFunc` closure that folds `s` through each handler in order.

## Solution

```go
func (f HandlerFunc) Handle(s string) string { return f(s) }

func Run(h Handler, s string) string { return h.Handle(s) }

func Chain(hs ...Handler) Handler {
	return HandlerFunc(func(s string) string {
		for _, h := range hs {
			s = h.Handle(s)
		}
		return s
	})
}
```

## Walkthrough

`Chain(exclaim, upper).Handle("hi")` gives `"hi!"` then `"HI!"` — the same output as the other order here, but the intermediate values differ, which matters for non-commutative handlers.

## Pitfalls

- `return HandlerFunc(f)(s)` — an infinite loop if you call `Handle` again instead of `f`.
- Returning a plain `func` from `Chain`; the result must be converted to `HandlerFunc` to satisfy `Handler`.
- Applying handlers right to left.
