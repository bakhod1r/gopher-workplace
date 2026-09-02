# Atomic Implementation Swap

## Intuition

`atomic.Pointer[T]` needs a concrete type, and an interface is not one. Wrapping it in a one-field struct gives you a pointer to swap, and the pointer store publishes the whole interface value — both words — atomically.

## Approach

1. `Set` builds a `&holder{p: p}` and stores it with one atomic write.
2. `Get` loads the holder and returns its policy, or nil.
3. `Allow` loads once, treats a nil holder or nil policy as fail-open, and dispatches otherwise.
4. Nothing on the request path takes a lock.

## Solution

```go
func (s *Strategy) Set(p Policy) {
	s.cur.Store(&holder{p: p})
}

func (s *Strategy) Get() Policy {
	h := s.cur.Load()
	if h == nil {
		return nil
	}
	return h.p
}

func (s *Strategy) Allow(key string) bool {
	h := s.cur.Load()
	if h == nil || h.p == nil {
		return true
	}
	return h.p.Allow(key)
}
```

## Walkthrough

Loading the holder pointer *once* into a local matters: two separate loads could straddle a swap and mix a nil check against one policy with a dispatch to another.

## Pitfalls

- Storing an interface in a plain field and swapping it without atomics — a two-word write that readers can tear.
- Calling `s.cur.Load()` twice in `Allow`, reintroducing the straddle window.
- Using `atomic.Value` with different dynamic types across `Set` calls, which panics at runtime.
