# Match Through Wrapping

## Intuition

Once errors are wrapped, identity comparison only sees the outermost layer. `errors.Is` unwraps repeatedly and compares at every level, which is why sentinels survive being annotated.

## Approach

1. Call `errors.Is` with the sentinel as the target.
2. Return its result directly.

## Solution

```go
return errors.Is(err, ErrNotFound)
```

## Walkthrough

For the twice-wrapped case, `errors.Is` compares the outer error, unwraps, compares again, unwraps once more, and matches `ErrNotFound`.

## Pitfalls

- Writing `err == ErrNotFound`, which fails as soon as anyone wraps it.
- Comparing `err.Error() == "not found"` — string matching breaks on any prefix.
- Guarding nil by hand; `errors.Is(nil, target)` is already false.
