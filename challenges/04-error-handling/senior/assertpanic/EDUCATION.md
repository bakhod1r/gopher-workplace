# Did It Panic

## Intuition

`recover` returning nil used to be ambiguous: no panic, or a panic with a nil payload? Modern Go removes the ambiguity by substituting a real value for `panic(nil)`.

## Approach

1. Defer a closure that recovers.
2. Set both named results when the recovered value is non-nil.
3. Call `f`.

## Solution

```go
defer func() {
	if r := recover(); r != nil {
		value = r
		panicked = true
	}
}()
f()
return nil, false
```

## Walkthrough

`panic(nil)` is recovered as a `*runtime.PanicNilError`, so the helper still reports `true` with a non-nil payload.

## Pitfalls

- Returning `recover() != nil` from the function body, where recover is always nil.
- Assuming a nil recovered value means no panic.
- Letting the panic escape by recovering in the wrong scope.
