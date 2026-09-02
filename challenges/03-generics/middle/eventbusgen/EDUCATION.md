# Event Bus

## Intuition

A map keyed by an ever-increasing id makes unsubscription safe: with slice indexes, removing one handler would silently renumber the others.

## Approach

1. `Subscribe`: bump the counter, allocate the map when nil, store and return the id.
2. `Unsubscribe`: delete when present.
3. `Publish`: call every handler and count.

## Solution

```go
func (b *Bus[T]) Subscribe(h func(T)) int {
	b.nextID++
	if b.handlers == nil {
		b.handlers = make(map[int]func(T))
	}
	b.handlers[b.nextID] = h
	return b.nextID
}

func (b *Bus[T]) Unsubscribe(id int) bool {
	if _, ok := b.handlers[id]; !ok {
		return false
	}
	delete(b.handlers, id)
	return true
}

func (b *Bus[T]) Publish(v T) int {
	n := 0
	for _, h := range b.handlers {
		h(v)
		n++
	}
	return n
}
```

## Walkthrough

Unsubscribing an id removes exactly that handler; the remaining ids keep working.

## Pitfalls

- Storing handlers in a slice and invalidating ids on removal.
- Panicking on a nil map because the zero value was never allocated.
- Reusing ids after deletion, so a stale id silently removes a new handler.
