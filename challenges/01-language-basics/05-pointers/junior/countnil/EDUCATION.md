# Testing pointers for nil

## Intuition

The zero value of any pointer is nil; comparing with `== nil` detects unset references before use.

## Approach

1. Range over the slice.
2. Increment a counter when `p == nil`.
3. Return the count.

## Solution

```go
func CountNil(ps []*int) int {
	c := 0
	for _, p := range ps {
		if p == nil {
			c++
		}
	}
	return c
}
```

## Walkthrough

`[]*int{&a, nil, nil}`: the loop sees one real pointer and two nils, counting `2`.

## Pitfalls

- Only `== nil` distinguishes an unset pointer.
- Dereferencing the nil ones would panic.
