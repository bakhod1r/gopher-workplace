# Formatting integers by hand

## Intuition

Extract digits from the right with `n%10`, map each to a character with
`byte('0'+d)`, divide by 10, and reverse:

```go
for n > 0 { buf = append(buf, byte('0'+n%10)); n /= 10 }
// reverse
```

## Approach

1. n==0 -> "0". 2. Track sign; work in the negative range (n = -n for positives) so math.MinInt won't overflow. 3. Extract digits via '0'-n%10 while n<0. 4. Append '-' if negative, then reverse.

## Solution

```go
func Format(n int) string {
	if n == 0 {
		return "0"
	}
	neg := n < 0
	if !neg {
		n = -n
	}
	var buf []byte
	for n < 0 {
		buf = append(buf, byte('0'-n%10))
		n /= 10
	}
	if neg {
		buf = append(buf, '-')
	}
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}
```

## Walkthrough

Format(-17): neg. digits from -17: 7,1 -> buf "71", append '-' -> "71-", reverse -> "-17".

## Pitfalls

- `Format(0)` must special-case to `"0"`, else the loop produces "".
- Taking `abs(n)` risks overflow for the most-negative int; a robust version
  works in the negative domain instead.
- Digits come out reversed — remember to flip the buffer.
