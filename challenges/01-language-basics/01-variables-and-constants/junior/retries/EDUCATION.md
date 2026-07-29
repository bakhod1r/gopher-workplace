# Rune literals vs integer literals

## Intuition

`3` and `'3'` are both integer constants in Go. They are not the same number.

```go
3     // untyped integer constant, value 3
'3'   // untyped rune constant, value 51 — the Unicode code point of the digit
```

Single quotes mean **character**, and a character is stored as its code point.
`'0'` is 48, `'A'` is 65, `'a'` is 97. The digit characters run consecutively
from 48, which is why `'7' - '0'` gives 7 — the classic char-to-digit trick.

## Why the mistake survives compilation

An untyped rune constant is still an integer constant, so it goes wherever an
integer is expected:

```go
const MaxRetries = '3'      // 51
attempts := 1 + MaxRetries  // 52 — compiles cleanly
```

No conversion, no warning. The program simply retries 51 times. Nothing about
the syntax is wrong; only the meaning is.

## The three quoting forms

| Literal | Type | Meaning |
|---------|------|---------|
| `'3'` | rune (int32) | one character, value 51 |
| `"3"` | string | a one-character string |
| `3` | integer | the number three |

Double quotes never produce a number: `1 + "3"` does not compile, which is why
the string version of this mistake is caught immediately and the rune version is
not.

## Untyped constants inherit kinds

An untyped constant carries a *kind* — integer, rune, float, string, bool — and
that kind propagates through constant arithmetic:

```go
const a = '3'      // rune kind
const b = a + 1    // still rune kind, value 52
var n int = b      // fine: the value fits an int
fmt.Printf("%c", b) // '4' — printing as a character reveals the kind
```

The kind decides how the value *reads*, not whether it compiles.

## Approach

1. Declare `MaxRetries = 3` (the number, not the rune `'3'`).
2. `Budget` is `1 + MaxRetries`.

## Solution

```go
const MaxRetries = 3

func Budget() int {
	return 1 + MaxRetries
}
```

## Walkthrough

The first try plus `MaxRetries` retries gives 4 total attempts.

## Pitfalls

- A rune is an alias for `int32`; `byte` is an alias for `uint8`. `'é'` needs 2
  bytes in UTF-8 but is one rune, value 233.
- `%c` prints a number as a character, `%d` as a number, `%q` as a quoted
  character. Reach for `%c`/`%q` when a value looks 48 too big.
- Digit conversion is `'7' - '0'`, never `int('7')` — the latter gives 55.
