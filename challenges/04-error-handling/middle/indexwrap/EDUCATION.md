# Annotate With Position

## Intuition

Annotation is most useful when it is specific: not "parse failed" but "record 3 failed to parse". `%w` lets the specific message and the matchable cause coexist in one error.

## Approach

1. Return nil for a nil error.
2. Format with `%d` for the index and `%w` for the cause.

## Solution

```go
if err == nil {
	return nil
}
return fmt.Errorf("record %d: %w", i, err)
```

## Walkthrough

`AtIndex(0, ErrParse)` produces `"record 0: parse failed"` — index zero is a real position and is not special-cased.

## Pitfalls

- Skipping the nil guard and annotating a success.
- Using `%v` for the cause, which breaks `errors.Is`.
- Treating index 0 as "no index".
