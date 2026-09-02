# Logger

## Intuition

Depending on a `Logger` interface instead of a concrete writer turns output into a test seam: production plugs in a real sink, tests plug in memory.

## Approach

1. `(*MemLogger).Log` appends to the slice field.
2. `Discard.Log` has an empty body.
3. `LogAll` ranges over `lines`, calling `l.Log`.

## Solution

```go
func (m *MemLogger) Log(line string) { m.Lines = append(m.Lines, line) }

func (d Discard) Log(line string) {}

func LogAll(l Logger, lines []string) {
	for _, line := range lines {
		l.Log(line)
	}
}
```

## Walkthrough

`LogAll(m, []string{"x","y","z"})` calls `Log` three times; each `append` grows `m.Lines`, so the last entry is `"z"`.

## Pitfalls

- Value receiver on `MemLogger.Log` — the appended slice is discarded with the copy.
- Passing `MemLogger{}` instead of `&MemLogger{}`; only the pointer is in the method set.
- Reversing the order by prepending instead of appending.
