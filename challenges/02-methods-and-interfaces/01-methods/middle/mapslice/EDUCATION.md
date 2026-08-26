# Mapping Slice Types

## Intuition

Go lacks a generic `.Map()` method, so you write them manually. Since the length
is known exactly, it is best practice to pre-allocate the result slice.

## Approach

1. If `l == nil`, return `[]string{}`.
2. `make([]string, len(l))`.
3. Loop and convert.

## Solution

```go
func (l IntList) ToString() []string {
	if l == nil {
		return []string{}
	}
	res := make([]string, len(l))
	for i, v := range l {
		res[i] = fmt.Sprint(v)
	}
	return res
}
```

## Walkthrough

For `{1, 2}`:
- `res` = `["", ""]` (length 2).
- `res[0] = "1"`.
- `res[1] = "2"`.
- Returns `["1", "2"]`.

## Pitfalls

- Using `append` in a loop without `make` works but is slower due to reallocation.
- Using `make([]string, len(l))` and then `append` creates `["", "", "1", "2"]`.
  If you use `make` with length, use index assignment. If you use `make` with
  capacity (`make([]string, 0, len(l))`), use `append`.
