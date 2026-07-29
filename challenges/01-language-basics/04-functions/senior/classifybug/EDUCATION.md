# Handling the default case

## Intuition

A switch without a default (or a catch-all return) silently produces the zero value for unlisted inputs.

## Approach

1. A `switch` with no matching case and no `default` returns the zero value.
2. Add a `default: return "unknown"`.

## Solution

```go
func Class(code int) string {
	switch code / 100 {
	case 2:
		return "success"
	case 4:
		return "client"
	case 5:
		return "server"
	default:
		return "unknown"
	}
}
```

## Walkthrough

`Class(301)` matches no case and falls out returning "". A `default` branch supplies "unknown" for unmatched codes.

## Pitfalls

- Always define behaviour for inputs outside the enumerated cases.
- `default:` documents the catch-all intent.
