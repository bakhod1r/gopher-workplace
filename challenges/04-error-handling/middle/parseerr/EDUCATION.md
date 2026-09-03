# Error With A Line Number

## Intuition

A message can always be reconstructed from data, but data cannot be reliably parsed back out of a message. Store the line in a field and format it only at the edge.

## Approach

1. Format `Line` and `Msg` in `Error`.
2. Declare `var pe *ParseError` in `LineOf`.
3. Use `errors.As` and return the field.

## Solution

```go
// Error:
return fmt.Sprintf("line %d: %s", e.Line, e.Msg)

// LineOf:
var pe *ParseError
if errors.As(err, &pe) {
	return pe.Line, true
}
return 0, false
```

## Walkthrough

The wrapped case reads `"config: line 9: eof"`, yet `errors.As` still hands back the struct so `Line` is exactly 9.

## Pitfalls

- Parsing the line back out of the message string.
- Type-asserting instead of using `errors.As`, which fails under wrapping.
- Returning `pe.Line` without checking the boolean, dereferencing nil.
