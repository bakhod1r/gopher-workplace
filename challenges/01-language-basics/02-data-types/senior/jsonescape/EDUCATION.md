# Escaping the escape character

## The idea

In JSON, the backslash introduces every escape, so a literal backslash must
itself be written as `\\`. If you escape quotes and newlines but forget the
backslash, any input containing `\` yields invalid JSON:

```go
case '\\':
	b.WriteString(`\\`) // raw literal: two backslash characters
```

## Why it matters

Serialization must escape the delimiter **and** the escape character. Missing the
backslash is a real injection/parse bug — Windows paths, regexes, and LaTeX in
user data all contain backslashes.

## Watch out

- Escape the backslash *before* other rules conceptually; order in a `switch`
  doesn't matter since each rune matches once.
- A raw literal `` `\\` `` is two characters; `"\\\\"` is the interpreted
  equivalent.
- Prefer `encoding/json` in production; hand-escaping is for understanding.
