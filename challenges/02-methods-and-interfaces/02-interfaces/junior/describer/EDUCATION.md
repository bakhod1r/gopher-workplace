# Describer

## Intuition

`DescribeAll` is written once against the interface and works for every current and future implementer.

## Approach

1. Format each type with `fmt.Sprintf`.
2. Preallocate `make([]string, 0, len(ds))`.
3. Append `d.Describe()` in a `range` loop.

## Solution

```go
func (u User) Describe() string { return fmt.Sprintf("user %s", u.Name) }

func (s Server) Describe() string {
	return fmt.Sprintf("server %s:%d", s.Host, s.Port)
}

func DescribeAll(ds []Describer) []string {
	out := make([]string, 0, len(ds))
	for _, d := range ds {
		out = append(out, d.Describe())
	}
	return out
}
```

## Walkthrough

`Server{Host: "h", Port: 80}` formats `%s` as `h` and `%d` as `80`, giving `"server h:80"`.

## Pitfalls

- `make([]string, len(ds))` then `append` — you get leading empty strings.
- `%d` on a string or `%s` on an int — `go vet` catches it.
- Returning `nil` for an empty input is fine here, but the length must be 0.
