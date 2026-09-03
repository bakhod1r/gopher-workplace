# Wrap A Library Error

## Intuition

Library errors already carry precise information. Replacing them with your own message throws that away; wrapping them keeps both the library's detail and your context.

## Approach

1. Call `strconv.Atoi`.
2. Return `0` and a `%w` wrapper on failure, quoting the input with `%q`.
3. Return the value and nil otherwise.

## Solution

```go
n, err := strconv.Atoi(s)
if err != nil {
	return 0, fmt.Errorf("parse %q: %w", s, err)
}
return n, nil
```

## Walkthrough

`strconv.Atoi("x")` returns a `*strconv.NumError` wrapping `strconv.ErrSyntax`; the `%w` wrapper adds a layer and `errors.Is` still reaches the sentinel.

## Pitfalls

- Returning `errors.New("bad number")`, discarding the cause.
- Using `%v` for the cause and breaking `errors.Is`.
- Returning the partial value `n` alongside the error.
