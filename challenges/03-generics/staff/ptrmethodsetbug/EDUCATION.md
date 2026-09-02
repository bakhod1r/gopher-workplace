# The Reset That Resets A Copy

## Intuition

The pointer constraint gets the mutating method into scope, and then the loop hands it the address of the loop variable rather than the address of the element. The mutation lands on a copy that is discarded at the end of the iteration.

## Approach

1. Range over the indices, not the values.
2. Convert `&s[i]` to `PT` and call `Reset` on it.

## Solution

```go
func ResetAll[T any, PT Resettable[T]](s []T) {
	for i := range s {
		PT(&s[i]).Reset()
	}
}
```

## Walkthrough

`ResetAll` over `[{3} {4}]` zeroes two temporaries and leaves the slice untouched, with no error anywhere.

## Pitfalls

- Ranging over values and reassigning `s[i] = v` afterwards — it works, but it copies twice and hides the intent.
- Declaring the constraint as `interface{ Reset() }` on `T`, which no value type with a pointer-receiver method satisfies.
