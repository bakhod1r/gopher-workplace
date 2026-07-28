# Split vs Fields

## The idea

`strings.Split(s, " ")` cuts on **every** single space, so consecutive spaces
produce empty strings. `strings.Fields(s)` splits on runs of *any* whitespace and
drops the empties:

```go
len(strings.Fields(s)) // correct word count
```

## Why it matters

Tokenizing user text, log lines, and CLI input needs whitespace-run semantics.
Using `Split` with a single-space separator is a frequent, silent off-by-many
bug.

## Watch out

- `Fields` treats tabs and newlines as separators too — usually what you want.
- `Split("", " ")` returns `[""]` (length 1); `Fields("")` returns length 0.
- If you need a custom separator set, `strings.FieldsFunc` takes a predicate.
