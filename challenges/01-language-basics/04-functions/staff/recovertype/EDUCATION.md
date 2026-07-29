# Type-asserting a recovered value

## Intuition

`recover()` yields `any`; extracting a specific type needs a (comma-ok) assertion — asserting the wrong type silently misses the value.

## Approach

1. Only an `error` panic should be captured into `err`.
2. The bug type-asserts to `string` and drops it; assert to `error` and assign.

## Solution

```go
func Call(f func()) (err error) {
	defer func() {
		r := recover()
		if e, ok := r.(error); ok {
			err = e
		}
	}()
	f()
	return
}
```

## Walkthrough

Asserting `r.(string)` ignores an error panic, leaving `err` nil. Asserting `r.(error)` captures `errBoom`, while a non-error panic like "x" leaves `err` nil.

## Pitfalls

- Recovered value is `any`; use `r.(error)` (comma-ok) to get an error.
- A plain assertion of a wrong type would panic — always use comma-ok here.
