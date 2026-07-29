# Raw string literals

## Intuition

Go has two string literal forms. Interpreted literals use double quotes and
process escape sequences (`\n`, `\t`, `\\`). **Raw** literals use backticks and
take every character verbatim — no escapes, so a backslash is just a backslash:

```go
`C:\Users\temp` // 13 characters, three literal backslashes
"C:\\Users\\temp" // same string, doubled backslashes
```

## Approach

1. Return a backtick-delimited raw string literal.
2. Inside backticks backslashes are literal, so no doubling is needed.

## Solution

```go
func TempPath() string {
	return `C:\Users\temp\log.txt`
}
```

## Walkthrough

`C:\Users\temp\log.txt` is taken character-for-character, giving the path with single backslashes.

## Pitfalls

- The only thing a raw literal cannot contain is a backtick.
- Carriage returns (`\r`) are stripped, but newlines are kept, so raw multi-line
  strings embed real line breaks.
- Because nothing is escaped, `\n` inside backticks is two characters, not a
  newline.
