# Escaping quotes

## The idea

Inside a double-quoted string, a literal double quote must be escaped as `\"`,
because an unescaped `"` would end the string:

```go
"\"" // a string of length 1: one quote character
```

## Why it matters

Generating JSON-ish output, CSV fields, or shell arguments by hand requires
embedding quotes. Knowing `\"` is one character keeps lengths and comparisons
correct.

## Watch out

- `len("\"")` is 1.
- A raw literal sidesteps escaping quotes but **cannot contain a backtick**, so
  it cannot hold a backtick the way `"..."` holds a quote.
- The standard `strconv.Quote` adds quotes *and* escapes contents — different
  job from simple wrapping.

## Try it yourself

```go
"\""          // one quote char
`"`           // also one quote char (raw)
len("\"hi\"") // 4
```
