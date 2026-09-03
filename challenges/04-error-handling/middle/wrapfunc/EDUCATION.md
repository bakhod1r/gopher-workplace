# Decorate A Step

## Intuition

A decorator moves a cross-cutting concern out of every implementation. The step knows how to fail; the wrapper knows what the step is called.

## Approach

1. Return a closure over `name` and `f`.
2. Inside it, call `f` and return nil on success.
3. Otherwise wrap with `%w`.

## Solution

```go
return func() error {
	if err := f(); err != nil {
		return fmt.Errorf("%s: %w", name, err)
	}
	return nil
}
```

## Walkthrough

Calling the decorated step twice runs `f` twice — the closure holds the function, not a cached result.

## Pitfalls

- Calling `f` while building the decorator.
- Wrapping the nil result, turning success into failure.
- Returning `f` unchanged when the name is empty.
