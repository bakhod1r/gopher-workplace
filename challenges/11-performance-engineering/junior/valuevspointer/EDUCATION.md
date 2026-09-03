# Receivers Decide What Gets Copied

## Intuition

`func (c Counter)` means "give me my own Counter". `func (c *Counter)` means "give me the caller's".

## Approach

1. `Inc` and `Value` use the pointer receiver.
2. `IncCopy` mutates its own copy and hands it back.

## Solution

```go
func (c *Counter) Inc() {
	c.n++
}

func (c *Counter) Value() int {
	return c.n
}

func (c Counter) IncCopy() Counter {
	c.n++
	return c
}
```

## Walkthrough

`IncCopy` copies 520 bytes in and 520 bytes out per call; `Inc` passes one pointer. Both are allocation-free — copying happens on the stack — but the cycles are not equal.

## Pitfalls

- A value receiver on `Inc`, which silently discards every increment.
- Assuming a value receiver allocates; it copies, which is a different cost.
- Mixing receiver kinds without reason, so `Counter` satisfies an interface but `*Counter` behaves differently.
