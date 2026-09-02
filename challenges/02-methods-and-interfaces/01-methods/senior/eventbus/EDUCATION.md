# Event Bus

## Intuition

A bus inverts the call direction. Instead of the publisher naming who to call,
it names *what happened*; the bus owns the mapping from event to reactions. Add
a subscriber and no publishing code changes.

## Approach

1. `On`: append to the slice under the key — Go's zero values make the
   "first listener for this key" case identical to the rest.
2. `Emit`: read the slice and call each element with the payload.

## Solution

```go
func (b *Bus) On(eventType string, listener func(data string)) {
	b.listeners[eventType] = append(b.listeners[eventType], listener)
}

func (b *Bus) Emit(eventType string, data string) {
	for _, l := range b.listeners[eventType] {
		l(data)
	}
}
```

## Walkthrough

The first `On("user.login", ...)` reads a missing key, gets a nil slice, and
`append` allocates a one-element slice which is written back. The second `On`
appends to it. `Emit("user.login", "alice")` ranges over both and calls them, so
`got1` and `got2` are set in order.

`Emit("user.logout", "bob")` reads a missing key. The lookup yields a nil slice,
and `range` over nil runs zero times — no `ok` check needed.

## Pitfalls

- **Forgetting the write-back.** `append(b.listeners[k], l)` on its own throws
  the result away; maps do not hand out addressable values.
- **Guarding with `if _, ok := ...`.** Harmless but pointless — the zero value
  already does the right thing in both methods.
- **Emitting in a goroutine.** The tests assert synchronously; async delivery
  would also race on the captured variables.

## Reading a map value you cannot address

`b.listeners[k] = append(b.listeners[k], v)` is the canonical map-of-slice
idiom precisely because `b.listeners[k]` is not addressable — you cannot call a
pointer-receiver method on it, and you cannot append "in place".
