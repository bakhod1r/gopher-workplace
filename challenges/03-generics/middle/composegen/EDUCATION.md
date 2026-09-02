# Compose

## Intuition

The intermediate type `B` exists only inside the composition. Callers never name it, which is what makes composed pipelines readable.

## Approach

1. Return `func(a A) C { return g(f(a)) }`.

## Solution

```go
func Compose[A, B, C any](f func(A) B, g func(B) C) func(A) C {
	return func(a A) C {
		return g(f(a))
	}
}
```

## Walkthrough

`Compose(double, itoa)(2)` runs `double(2) = 4`, then `itoa(4) = "4"`.

## Pitfalls

- Applying `g` first — composition order is left to right here.
- Trying to write it with two type parameters, which forces `B` to equal `A` or `C`.
- Calling `f` and `g` eagerly instead of returning a closure.
