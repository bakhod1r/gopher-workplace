# Middleware

## Intuition

A middleware is a function from handler to handler. Composition happens at build time; at request time the layers are just nested calls.

## Approach

1. `Handle` calls the receiver function.
2. Each middleware returns `func(next Handler) Handler` producing a `HandlerFunc` closure that does its work and calls `next.Handle`.
3. `Apply` loops from the last middleware to the first, rewrapping `h` each time.

## Solution

```go
func (f HandlerFunc) Handle(s string) string { return f(s) }

func WithCount(n *int) Middleware {
	return func(next Handler) Handler {
		return HandlerFunc(func(s string) string {
			*n++
			return next.Handle(s)
		})
	}
}

func WithPrefix(p string) Middleware {
	return func(next Handler) Handler {
		return HandlerFunc(func(s string) string {
			return p + next.Handle(s)
		})
	}
}

func Apply(h Handler, ms ...Middleware) Handler {
	for i := len(ms) - 1; i >= 0; i-- {
		h = ms[i](h)
	}
	return h
}
```

## Walkthrough

`Apply(echo, a, b)` first wraps with `b`, then with `a`. At call time `a` runs first and prefixes `"a:"` onto whatever `b` produced — `"a:b:x"`.

## Pitfalls

- Wrapping front to back, which reverses the visible order to `"b:a:x"`.
- Prefixing before calling `next`, which is fine here but changes semantics for middlewares that transform the input.
- Capturing the loop variable instead of the parameter when building closures.
