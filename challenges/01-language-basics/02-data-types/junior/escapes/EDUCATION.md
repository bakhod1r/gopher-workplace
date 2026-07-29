# Interpreted string literals

## Intuition

Double-quoted strings are **interpreted**: backslash escape sequences become
single characters. `\t` is one tab byte, `\n` one newline byte, `\\` one
backslash, `\"` a quote:

```go
"a\tb\n" // 4 bytes: 'a', tab, 'b', newline
```

## Approach

1. Concatenate name + "\t" + value + "\n".
2. \t and \n are single bytes inside an interpreted (double-quoted) literal.

## Solution

```go
func Row(name, value string) string {
	return name + "\t" + value + "\n"
}
```

## Walkthrough

Row("id","42"): "id" + tab + "42" + newline = "id\t42\n".

## Pitfalls

- `len("\t")` is 1, not 2 — the escape is a single byte.
- Unknown escapes are a compile error (`\q` is invalid).
- `é` / `\U0001F600` embed Unicode code points; `\x41` embeds a raw byte.
- For text full of backslashes, a raw `` `...` `` literal is clearer.
