# Run-length encoding

## Intuition

Scan consecutive equal bytes, counting the run; when the byte changes (or the
string ends), emit the character and its count:

```go
for i := 0; i < len(s); {
	j := i
	for j < len(s) && s[j] == s[i] { j++ }
	b.WriteByte(s[i]); b.WriteString(strconv.Itoa(j - i))
	i = j
}
```

## Approach

1. Scan s with index i.
2. Extend j while s[j]==s[i] to find run end.
3. Write the char then the run length (strconv.Itoa).
4. Set i=j and continue.
5. Return builder string.

## Solution

```go
import (
	"strconv"
	"strings"
)

func Encode(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); {
		j := i
		for j < len(s) && s[j] == s[i] {
			j++
		}
		b.WriteByte(s[i])
		b.WriteString(strconv.Itoa(j - i))
		i = j
	}
	return b.String()
}
```

## Walkthrough

"aaab": i=0 run of a length 3 -> "a3"; i=3 run of b length 1 -> "b1"; result "a3b1".

## Pitfalls

- Build with a `Builder`, not `+=`.
- Decoding must handle multi-digit counts (`a12`).
- RLE only helps when runs are long; it can expand random data.
