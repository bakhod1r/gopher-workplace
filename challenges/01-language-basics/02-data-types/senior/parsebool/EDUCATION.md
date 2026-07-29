# Recognized vs default

## Intuition

Return two bools: the parsed value and whether the input was recognized. That
lets the caller distinguish an explicit `false` from an unknown string that
should keep a default. Missing an accepted form (`off`) silently turns it into
"unknown".

## Approach

1. Bug: the false case omitted "off", so "off" fell through to unrecognized.
2. Fix: add "off" to the false case list.
3. Input is lowercased and trimmed first.

## Solution

```go
import "strings"

func Parse(s string) (bool, bool) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "true", "1", "yes", "on":
		return true, true
	case "false", "0", "no", "off":
		return false, true
	}
	return false, false
}
```

## Walkthrough

"off": ToLower/Trim -> "off" matches false case -> (false,true).

## Pitfalls

- Normalize (`ToLower` + `TrimSpace`) before matching.
- Keep truthy and falsey sets symmetric (`on`/`off`, `yes`/`no`).
- Unknown input returns `ok=false`, not `false,true`.
