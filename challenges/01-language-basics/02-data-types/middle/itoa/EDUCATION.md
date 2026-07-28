# Formatting integers by hand

## The idea

Extract digits from the right with `n%10`, map each to a character with
`byte('0'+d)`, divide by 10, and reverse:

```go
for n > 0 { buf = append(buf, byte('0'+n%10)); n /= 10 }
// reverse
```

## Why it matters

It mirrors `atoi` and is what `strconv.Itoa` does underneath. Handling the sign
and the zero case cleanly is the crux.

## Watch out

- `Format(0)` must special-case to `"0"`, else the loop produces "".
- Taking `abs(n)` risks overflow for the most-negative int; a robust version
  works in the negative domain instead.
- Digits come out reversed — remember to flip the buffer.
