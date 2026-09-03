# One Buffer, Many Records

## Intuition

Allocation happens when you need new memory. Rendering into the same array every time means you need new memory exactly once — the first time the record is bigger than anything before it.

## Approach

1. Reset the buffer to length zero.
2. Append each pair, bounded by the shorter slice.
3. `Clone` copies the current contents.

## Solution

```go
func (e *Encoder) Encode(names, values []string) []byte {
	e.buf = e.buf[:0]
	for i := 0; i < min(len(names), len(values)); i++ {
		e.buf = append(e.buf, names[i]...)
		e.buf = append(e.buf, '=')
		e.buf = append(e.buf, values[i]...)
		e.buf = append(e.buf, ';')
	}
	return e.buf
}

func (e *Encoder) Clone() []byte {
	return slices.Clone(e.buf)
}
```

## Walkthrough

`e.buf[:0]` on a nil slice is legal and yields an empty nil slice, so the zero `Encoder` works without a constructor; the first appends allocate, and every later call reuses that array.

## Pitfalls

- Forgetting the reset, so record two reads `a=1;b=2;`.
- Handing the returned slice to something that stores it — it will be overwritten under them.
- Returning `e.buf` from `Clone`, which is exactly the aliasing the method exists to avoid.
