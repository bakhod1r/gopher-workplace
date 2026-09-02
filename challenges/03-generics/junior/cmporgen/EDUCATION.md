# cmp.Or

## Intuition

`cmp.Or` is the stdlib version of the coalesce you built by hand, and it carries the same caveat: a deliberately empty value is indistinguishable from an absent one.

## Approach

1. Return `cmp.Or` of the three values in priority order.

## Solution

```go
func Display(nickname, username string) string {
	return cmp.Or(nickname, username, "anonymous")
}
```

## Walkthrough

`Display("", "user")` skips the empty nickname and returns `"user"` without evaluating the final fallback's role.

## Pitfalls

- Nesting `if` statements where one call suffices.
- Putting the fallback first, so it always wins.
- Expecting `cmp.Or` to distinguish "set to empty" from "unset".
