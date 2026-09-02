# Every Case Is The Last Case

## Intuition

Because `c` is declared before the loop, `for _, c = range cases` overwrites one single variable. Every pointer appended is that same address, so after the loop they all read as the final case.

## Approach

1. Range over the indices.
2. Append the address of the element itself.
3. Return the slice.

## Solution

```go
func Pointers[T any](cases []Case[T]) []*Case[T] {
	out := make([]*Case[T], 0, len(cases))
	for i := range cases {
		out = append(out, &cases[i])
	}
	return out
}
```

## Walkthrough

For cases named a, b, c the result is three pointers to one variable holding `c`, so all three report "c".

## Pitfalls

- Assuming Go 1.22's per-iteration loop variables fix this — they only apply when the variable is declared by the `range` clause.
- Copying into a fresh local each pass and taking *its* address, which works but detaches the pointer from the caller's slice.
