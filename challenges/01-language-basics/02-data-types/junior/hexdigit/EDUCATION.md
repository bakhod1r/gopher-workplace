# Character arithmetic with bytes

## The idea

ASCII character codes are contiguous within each run: `'0'..'9'` are 48..57,
`'a'..'f'` are 97..102. So you compute a digit character by offsetting from the
base:

```go
'0' + byte(n)      // n in 0..9 -> '0'..'9'
'a' + byte(n-10)   // n in 10..15 -> 'a'..'f'
```

## Why it matters

Encoders (hex, base32), parsers, and formatters lean on the contiguity of ASCII.
Character literals are just integer constants, so arithmetic on them is exact and
cheap.

## Watch out

- `'0'` is a rune constant (value 48); in a `byte` context it fits since it is
  ASCII.
- The digit/letter runs are contiguous *separately* — you cannot span the gap
  between `'9'` and `'a'` with one offset.
- Guard the input range, or you produce garbage bytes.

## Try it yourself

```go
'0' + 5      // 53 == '5'
'a' + 2      // 99 == 'c'
'9' + 1      // ':' , not 'a' — different runs
```
