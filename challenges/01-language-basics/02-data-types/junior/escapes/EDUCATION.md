# Interpreted string literals

## The idea

Double-quoted strings are **interpreted**: backslash escape sequences become
single characters. `\t` is one tab byte, `\n` one newline byte, `\\` one
backslash, `\"` a quote:

```go
"a\tb\n" // 4 bytes: 'a', tab, 'b', newline
```

## Why it matters

Formats like TSV/CSV, log lines, and protocols need exact control characters.
Escapes put them in a literal without embedding invisible bytes in your source.

## Watch out

- `len("\t")` is 1, not 2 — the escape is a single byte.
- Unknown escapes are a compile error (`\q` is invalid).
- `é` / `\U0001F600` embed Unicode code points; `\x41` embeds a raw byte.
- For text full of backslashes, a raw `` `...` `` literal is clearer.

## Try it yourself

```go
len("\n")      // 1
"tab\there"    // tab is one char
"é"       // "é"
```
