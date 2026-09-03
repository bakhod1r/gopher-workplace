# Keys That Print The Same

## Intuition

The set needs a key that means "this exact value". `T` is constrained to
`comparable`, which is precisely the promise that `T` itself can be that key.
Turning the value into text first throws away everything the constraint bought
you.

## Approach

1. Allocate the result with the input's length as capacity.
2. Keep a `map[T]struct{}` of values already emitted.
3. Skip an element that is already in the set; otherwise record and append it.

## Solution

```go
func Distinct[T comparable](vals []T) []T {
	out := make([]T, 0, len(vals))
	seen := make(map[T]struct{}, len(vals))
	for _, v := range vals {
		if _, dup := seen[v]; dup {
			continue
		}
		seen[v] = struct{}{}
		out = append(out, v)
	}
	return out
}
```

## Walkthrough

With `map[string]struct{}` and `fmt.Sprint(v)` as the key, `Distinct([]any{1, "1"})`
computes the key `"1"` twice and drops the string. The same happens for two
distinct struct types with identical fields: `{1 2}` renders identically for both,
so the second is discarded even though `point{1,2} == pair{1,2}` does not even
compile. Keying on `v` directly compares the dynamic type as well as the value,
which is what the documented rule asks for — and it removes a `fmt.Sprint` call,
with its reflection and allocation, from every iteration.

## Pitfalls

- Reaching for a string, hash, or `fmt` key when the type parameter is already `comparable`.
- Assuming `any` equality ignores the dynamic type; it does not.
- Keying on a hash without a fallback comparison, which trades one collision class for another.
- Building the set but forgetting that the *first* occurrence is the one to keep.
