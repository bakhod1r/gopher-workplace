# Required Field

## Intuition

"Empty" for a user-facing field means "nothing visible", not `len(s) == 0`. Normalising the input before the check collapses every blank shape into one case.

## Approach

1. Trim leading and trailing whitespace.
2. Return `ErrRequired` when the trimmed string is empty.
3. Return nil otherwise.

## Solution

```go
if strings.TrimSpace(s) == "" {
	return ErrRequired
}
return nil
```

## Walkthrough

`"\t\n"` trims to `""`, so the guard fires. `"  hi  "` trims to `"hi"` and passes.

## Pitfalls

- Checking `len(s) == 0` — spaces slip through.
- Trimming only spaces by hand, missing tabs and newlines.
- Mutating and returning the trimmed string; the contract returns only an error.
