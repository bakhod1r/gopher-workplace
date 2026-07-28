# Run-length encoding

## The idea

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

## Why it matters

RLE is a real compression primitive (bitmaps, fax, simple protocols) and a clean
exercise in run-scanning with a `strings.Builder`.

## Watch out

- Build with a `Builder`, not `+=`.
- Decoding must handle multi-digit counts (`a12`).
- RLE only helps when runs are long; it can expand random data.
