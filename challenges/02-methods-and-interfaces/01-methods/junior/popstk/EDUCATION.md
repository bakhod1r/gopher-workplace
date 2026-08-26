# Pop with Comma-OK

## Intuition

Go's comma-ok idiom (`v, ok := s.Pop()`) is preferred over panicking or
returning sentinel values. The caller can check `ok` and handle the empty case.

## Approach

1. If `len(s.Items) == 0`, return `0, false`.
2. Get the last element.
3. Reslice to remove it.
4. Return the value and `true`.

## Solution

```go
func (s *Stack) Pop() (int, bool) {
	n := len(s.Items)
	if n == 0 {
		return 0, false
	}
	v := s.Items[n-1]
	s.Items = s.Items[:n-1]
	return v, true
}
```

## Walkthrough

For `Stack{Items: []int{1, 2, 3}}`:
- `n` = 3.
- `v` = `s.Items[2]` = 3.
- `s.Items = s.Items[:2]` → `[1, 2]`.
- Returns `(3, true)`.

## Pitfalls

- Forgetting to reslice — the stack doesn't shrink.
- Index off-by-one: last element is at `n-1`, not `n`.
- Value receiver — the reslicing is lost to the caller.
