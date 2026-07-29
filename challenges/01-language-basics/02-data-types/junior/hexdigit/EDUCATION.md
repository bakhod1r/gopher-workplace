# Character arithmetic with bytes

## Intuition

ASCII character codes are contiguous within each run: `'0'..'9'` are 48..57,
`'a'..'f'` are 97..102. So you compute a digit character by offsetting from the
base:

```go
'0' + byte(n)      // n in 0..9 -> '0'..'9'
'a' + byte(n-10)   // n in 10..15 -> 'a'..'f'
```

## Approach

1. For 0..9 return '0'+byte(n).
2. For 10..15 return 'a'+byte(n-10).
3. Otherwise return '?'.

## Solution

```go
func Digit(n int) byte {
	switch {
	case n >= 0 && n <= 9:
		return '0' + byte(n)
	case n >= 10 && n <= 15:
		return 'a' + byte(n-10)
	default:
		return '?'
	}
}
```

## Walkthrough

Digit(10): in 10..15 branch, 'a'(97)+byte(0)=97='a'.

## Pitfalls

- `'0'` is a rune constant (value 48); in a `byte` context it fits since it is
  ASCII.
- The digit/letter runs are contiguous *separately* — you cannot span the gap
  between `'9'` and `'a'` with one offset.
- Guard the input range, or you produce garbage bytes.
