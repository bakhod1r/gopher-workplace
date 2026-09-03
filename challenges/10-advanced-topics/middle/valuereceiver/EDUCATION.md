# The Method That Copies Its Receiver

## Intuition

Receivers follow the same rules as parameters: a value receiver is a copy. That is fine for small types and quietly expensive for large ones — and it also means the method cannot see later writes.

## Approach

1. Return `c.Read` and `c.Write`.

## Solution

```go
// Config is a deliberately large settings block.
type Config struct {
	Read  int
	Write int
	Pad   [512]byte
}

// Timeouts returns the read and write timeouts from c.
//
// The receiver is a pointer because Config is large: a value receiver would
// copy the whole struct on every call.
//
// Examples:
//
// 	(&Config{Read: 1, Write: 2}).Timeouts() => 1, 2
func (c *Config) Timeouts() (read, write int) {
	return c.Read, c.Write
}
```

## Walkthrough

With a pointer receiver the call passes one word. With a value receiver it would push 528 bytes onto the stack, and the copy would be a snapshot rather than a live view.

## Pitfalls

- Dereferencing into a local `cfg := *c` — that reintroduces the copy you avoided.
- Nil receivers: a pointer receiver may be nil, so guard if the API allows it.
