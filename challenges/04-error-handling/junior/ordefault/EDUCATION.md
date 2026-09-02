# Value Or Default

## Intuition

A default is a decision that the failure does not matter. It is correct only when the error carries no information the caller needs — and the branch must be chosen by the error, never by the value.

## Approach

1. Return `def` when `err != nil`.
2. Return `v` otherwise.

## Solution

```go
if err != nil {
	return def
}
return v
```

## Walkthrough

`OrDefault("", nil, "80")` returns `""` — the lookup succeeded and found an empty setting, so the default does not apply.

## Pitfalls

- Falling back when `v == ""`, overriding a deliberate empty value.
- Returning the failed call's `v` (`"junk"`) instead of the default.
- Using this pattern where the error genuinely needed handling.
