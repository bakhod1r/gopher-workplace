# Builder Pattern

## Intuition

A builder is a mutable draft of an immutable result. Each setter writes one
field of the draft and then hands the draft back, so the caller can keep
talking to it. The chain is not magic — it is just "every method returns the
same pointer it was called on".

## Approach

1. Write the argument into the embedded `req`.
2. Return the receiver so the expression's type is still `*RequestBuilder`.
3. `Build` copies the finished draft out as a value.

## Solution

```go
func (b *RequestBuilder) URL(u string) *RequestBuilder {
	b.req.URL = u
	return b
}

func (b *RequestBuilder) Auth(t string) *RequestBuilder {
	b.req.Auth = t
	return b
}
```

## Walkthrough

`NewBuilder()` returns a `*RequestBuilder` whose `req` is the zero `Request`.
`.Method("GET")` writes `req.Method` and evaluates to the same pointer, so
`.URL("/api")` is a method call on that pointer, and so on. `Build()` returns
`b.req` **by value**, so later mutations of the builder cannot reach the
`Request` the caller already holds.

## Pitfalls

- **Forgetting `return b`.** The method then returns the zero `*RequestBuilder`
  (`nil`) or does not compile, and the chain panics or breaks.
- **Value receiver.** `func (b RequestBuilder) URL(...) *RequestBuilder` mutates
  a copy; `Build` would see an empty request.
- **Returning a fresh builder each step.** Legal, but then earlier fields must
  be copied forward — easy to lose one.
