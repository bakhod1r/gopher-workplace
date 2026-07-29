# Escaping quotes

## Intuition

Inside a double-quoted string, a literal double quote must be escaped as `\"`,
because an unescaped `"` would end the string:

```go
"\"" // a string of length 1: one quote character
```

## Approach

1. Build "\"" + s + "\"".
2. The \" escape places a literal double-quote byte inside the interpreted literal.

## Solution

```go
func Wrap(s string) string {
	return "\"" + s + "\""
}
```

## Walkthrough

Wrap("hello"): quote + hello + quote = a 7-byte string beginning and ending with a double quote.

## Pitfalls

- `len("\"")` is 1.
- A raw literal sidesteps escaping quotes but **cannot contain a backtick**, so
  it cannot hold a backtick the way `"..."` holds a quote.
- The standard `strconv.Quote` adds quotes *and* escapes contents — different
  job from simple wrapping.
