# Curry

## Intuition

Each returned closure captures its argument by reference to the variable, so the partially applied function stays valid long after the outer call returns.

## Approach

1. Return a closure taking `a` that returns a closure taking `b` and calling `f(a, b)`.

## Solution

```go
func Curry2[A, B, C any](f func(A, B) C) func(A) func(B) C {
	return func(a A) func(B) C {
		return func(b B) C {
			return f(a, b)
		}
	}
}
```

## Walkthrough

`plus1 := Curry2(add)(1)` captures `a = 1`; every later `plus1(x)` calls `add(1, x)`.

## Pitfalls

- Returning `func(A, B) C`, which is the original function.
- Calling `f` before both arguments are known.
- Swapping the argument order in the innermost call.
