# Optional

## Intuition

Inference works from arguments. `Some(5)` has one; `None()` has none, so the type argument must be supplied at the call site.

## Approach

1. `Some`: set both fields.
2. `None`: return the zero `Optional[T]`.
3. `Or`: return `def` when not present, else the value.

## Solution

```go
func Some[T any](v T) Optional[T] {
	return Optional[T]{value: v, present: true}
}

func None[T any]() Optional[T] {
	return Optional[T]{}
}

func (o Optional[T]) Or(def T) T {
	if !o.present {
		return def
	}
	return o.value
}
```

## Walkthrough

`Some(0).Or(9)` returns `0`: the value is present, so the fallback is not used.

## Pitfalls

- Writing `None()` and expecting inference to work.
- Comparing the stored value to zero instead of reading `present`.
- Giving `Or` a pointer receiver for no reason — it never mutates.
