# Allocation-Free Failure

## Intuition

Returning an existing sentinel copies an interface header and nothing else. Every `fmt.Errorf` on a hot failure path is a heap allocation the caller usually discards immediately.

## Approach

1. Return `ErrEmpty` for an empty string.
2. Return nil otherwise.

## Solution

```go
if s == "" {
	return ErrEmpty
}
return nil
```

## Walkthrough

`AllocsPerRun` reports 0 for both paths because no new value is constructed — the sentinel already exists.

## Pitfalls

- Wrapping the sentinel with `fmt.Errorf`, which allocates on every call.
- Building the error with `errors.New` inside the function.
- Returning a struct error by value that escapes to the heap.
