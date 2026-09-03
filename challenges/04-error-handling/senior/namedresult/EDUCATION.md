# Annotate On The Way Out

## Intuition

Annotating at every `return` scales badly: each new exit path is a chance to forget. A deferred rewrite of the named error covers them all, including the ones added later.

## Approach

1. Defer a closure that checks the named `err`.
2. Rewrite it with `%w` when non-nil.
3. Return `f()` directly.

## Solution

```go
defer func() {
	if err != nil {
		err = fmt.Errorf("%s: %w", op, err)
	}
}()
return f()
```

## Walkthrough

The deferred closure runs after `return f()` assigns the named results, so it sees the error and rewrites it before the caller does.

## Pitfalls

- Using unnamed results, so the defer cannot change what is returned.
- Wrapping unconditionally, converting nil into an error.
- Wrapping at each return statement and missing one.
