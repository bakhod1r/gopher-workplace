# Unique Failures

## Intuition

Repetition in an error report is noise that hides the one failure that is different. Deduplicating by message keeps the distinct signal while a map preserves the membership test in constant time.

## Approach

1. Track seen messages in a map.
2. Skip nil entries and already-seen messages.
3. Append the first occurrence of each message.

## Solution

```go
seen := make(map[string]bool)
var out []error
for _, err := range errs {
	if err == nil {
		continue
	}
	msg := err.Error()
	if seen[msg] {
		continue
	}
	seen[msg] = true
	out = append(out, err)
}
return out
```

## Walkthrough

`sameMsg` is a different error value with the same text as `ErrA`, so it is dropped — deduplication is by message, not identity.

## Pitfalls

- Deduplicating by error value, which keeps identical messages from distinct values.
- Ranging the map to build the output, losing order.
- Returning an empty non-nil slice when nothing failed.
