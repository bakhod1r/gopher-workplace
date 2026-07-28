# Raw string literals

## The idea

Go has two string literal forms. Interpreted literals use double quotes and
process escape sequences (`\n`, `\t`, `\\`). **Raw** literals use backticks and
take every character verbatim — no escapes, so a backslash is just a backslash:

```go
`C:\Users\temp` // 13 characters, three literal backslashes
"C:\\Users\\temp" // same string, doubled backslashes
```

## Why it matters

Regexes, Windows paths, and embedded snippets are far more readable as raw
strings — you avoid a forest of doubled backslashes. Raw literals may also span
multiple lines.

## Watch out

- The only thing a raw literal cannot contain is a backtick.
- Carriage returns (`\r`) are stripped, but newlines are kept, so raw multi-line
  strings embed real line breaks.
- Because nothing is escaped, `\n` inside backticks is two characters, not a
  newline.

## Try it yourself

```go
`a\tb`      // 4 chars: a \ t b
"a\tb"      // 3 chars: a <tab> b
`line1
line2`      // contains a real newline
```
