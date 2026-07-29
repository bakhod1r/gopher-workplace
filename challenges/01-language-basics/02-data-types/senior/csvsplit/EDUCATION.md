# Context-sensitive delimiters

## Intuition

A CSV comma is a field separator only when you are **not** inside a quoted field.
The parser carries an `inQuotes` flag; the split case must respect it:

```go
case c == ',' && !inQuotes:
```

## Approach

1. Bug: the comma case fires even inside a quoted field, wrongly splitting `"b,c"`.
2. Fix: guard the comma case with `&& !inQuotes` so only unquoted commas separate fields.
3. Quote toggling and "" un-escaping stay as-is.

## Solution

```go
func Split(line string) []string {
	var fields []string
	var cur []byte
	inQuotes := false
	for i := 0; i < len(line); i++ {
		c := line[i]
		switch {
		case c == '"':
			if inQuotes && i+1 < len(line) && line[i+1] == '"' {
				cur = append(cur, '"')
				i++
			} else {
				inQuotes = !inQuotes
			}
		case c == ',' && !inQuotes:
			fields = append(fields, string(cur))
			cur = cur[:0]
		default:
			cur = append(cur, c)
		}
	}
	fields = append(fields, string(cur))
	return fields
}
```

## Walkthrough

a,"b,c",d: inside quotes inQuotes=true so the middle comma appends to cur; only the two outer commas split.

## Pitfalls

- Toggle `inQuotes` on a lone quote; a doubled `""` inside quotes is a literal.
- The final field is flushed after the loop.
- Trailing/empty fields matter: `""` line yields one empty field.
