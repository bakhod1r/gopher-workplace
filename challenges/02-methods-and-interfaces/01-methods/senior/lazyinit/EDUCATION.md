# Lazy Initialization

## Intuition

"Compute on first use" needs a way to say *not yet*. For a `string` there is no
spare value — `""` is a perfectly good result. Promoting the field to `*string`
buys exactly one extra state, `nil`, and that state means "unset".

## Approach

1. Check the pointer.
2. On nil, run the initializer into a local and point the field at it.
3. Dereference and return.

## Solution

```go
func (l *LazyString) String() string {
	if l.val == nil {
		v := l.init()
		l.val = &v
	}
	return *l.val
}
```

## Walkthrough

`New` leaves `val` nil, so the test's early `calls != 0` check passes — nothing
has run. The first `String()` sees nil, calls `init` (counter → 1), and stores
the address of the local `v`. Escape analysis promotes `v` to the heap because
its address outlives the call. The second `String()` finds a non-nil pointer and
returns `*l.val` without touching `init`.

## Pitfalls

- **`&l.init()`.** Not addressable — a compile error. The local is mandatory.
- **Declaring `v` outside the `if`.** It is then re-initialized on every call and
  the sentinel logic gets confusing; keep the scope tight.
- **Value receiver.** The stored pointer is written to a copy, so `init` runs on
  every call.
- **Concurrency.** This is deliberately unsynchronized; two goroutines can both
  see nil. Use `sync.Once` when that matters (see the `onceinit` puzzle).
