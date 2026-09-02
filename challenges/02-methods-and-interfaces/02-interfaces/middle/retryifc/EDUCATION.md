# Retry Interface

## Intuition

Retry policy and the work being retried are separate concerns. The interface keeps the policy reusable, and error classification decides whether another attempt is even worth making.

## Approach

1. Each `Do` counts its call first, so the counter reflects real work.
2. `Flaky` fails while `Calls <= FailTimes`.
3. `Retry` loops `attempts` times, returning on success.
4. Return immediately when `errors.Is(err, ErrTemporary)` is false; otherwise fall out of the loop with the last error.

## Solution

```go
func (f *Flaky) Do() (string, error) {
	f.Calls++
	if f.Calls <= f.FailTimes {
		return "", ErrTemporary
	}
	return f.Value, nil
}

func Retry(op Op, attempts int) (string, error) {
	var err error
	for i := 0; i < attempts; i++ {
		var v string
		v, err = op.Do()
		if err == nil {
			return v, nil
		}
		if !errors.Is(err, ErrTemporary) {
			return "", err
		}
	}
	return "", err
}
```

## Walkthrough

`Flaky{FailTimes: 2}` with 3 attempts fails on calls 1 and 2 and succeeds on call 3. `Permanent` returns `ErrFatal` on call 1, so `Retry` exits without touching the remaining budget.

## Pitfalls

- Declaring `v, err := op.Do()` inside the loop with `:=`, which shadows the outer `err` and returns nil after exhaustion.
- Retrying fatal errors and burning the whole budget.
- Off-by-one: `for i := 0; i <= attempts` makes one extra call.
