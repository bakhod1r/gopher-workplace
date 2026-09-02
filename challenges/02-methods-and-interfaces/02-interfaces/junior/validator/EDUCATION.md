# Validator

## Intuition

Each rule is a small value with one method. The runner composes them in order and never grows when new rules appear.

## Approach

1. `NotEmpty.Validate` returns `ErrEmpty` for `""`, else `nil`.
2. `MaxLen.Validate` returns `ErrTooLong` when `len(s) > m.N`.
3. `ValidateAll` returns the first non-nil error, `nil` at the end.

## Solution

```go
func (n NotEmpty) Validate(s string) error {
	if s == "" {
		return ErrEmpty
	}
	return nil
}

func (m MaxLen) Validate(s string) error {
	if len(s) > m.N {
		return ErrTooLong
	}
	return nil
}

func ValidateAll(vs []Validator, s string) error {
	for _, v := range vs {
		if err := v.Validate(s); err != nil {
			return err
		}
	}
	return nil
}
```

## Walkthrough

`ValidateAll(rules, "abc")`: `NotEmpty` passes, `MaxLen{N: 2}` sees `len("abc") == 3 > 2` and returns `ErrTooLong`, which propagates unchanged so `errors.Is` matches.

## Pitfalls

- `>=` in `MaxLen` — a string of exactly N bytes is valid.
- Wrapping the sentinel with a new `errors.New`, which breaks `errors.Is`.
- Continuing the loop and returning the last error instead of the first.
