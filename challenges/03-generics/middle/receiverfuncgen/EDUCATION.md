# Functions That Act Like Methods

## Intuition

Taking the receiver as the first parameter is the standard workaround for the missing generic method — the call site reads nearly the same.

## Approach

1. `Update`: assign `f(b.value)` back into the box.
2. `Convert`: build a `*Box[U]` from `f(b.value)`.
3. `Get`: return the field.

## Solution

```go
func Update[T any](b *Box[T], f func(T) T) {
	b.value = f(b.value)
}

func Convert[T, U any](b *Box[T], f func(T) U) *Box[U] {
	return &Box[U]{value: f(b.value)}
}

func (b *Box[T]) Get() T {
	return b.value
}
```

## Walkthrough

`Convert(b, itoa)` produces a `*Box[string]` while `b` keeps holding its original value.

## Pitfalls

- Trying to declare `Convert` as a method with its own `U`.
- Making `Update` return a new box, which contradicts "in place".
- Mutating the source box inside `Convert`.
