# Verbose Chain

## Intuition

A wrapped message repeats everything beneath it, so a single line makes the reader diff prefixes by eye. One line per layer makes the added context obvious.

## Approach

1. Return `""` for nil.
2. Walk the chain collecting messages.
3. Prefix every line after the first and join with newlines.

## Solution

```go
if err == nil {
	return ""
}
var lines []string
for e := err; e != nil; e = errors.Unwrap(e) {
	if len(lines) == 0 {
		lines = append(lines, e.Error())
		continue
	}
	lines = append(lines, "caused by: "+e.Error())
}
return strings.Join(lines, "\n")
```

## Walkthrough

A two-wrap chain yields three lines, each shorter than the last, showing exactly what each layer contributed.

## Pitfalls

- Prefixing the first line as well.
- Using `errors.Unwrap` on a joined error and silently stopping.
- Returning `"<nil>"` for a nil error instead of an empty string.
