# io.Writer

## Intuition

`io.Writer` is the universal output seam. Because every write reports a count and an error, output code must thread both — and streaming keeps memory flat regardless of report size.

## Approach

1. Write the title line first and add its count to the total.
2. Return immediately when a write reports an error, along with the bytes written so far.
3. Loop the items, writing `"- " + item + "\n"` for each.
4. Return the total and `nil` at the end.

## Solution

```go
func WriteReport(w io.Writer, title string, items []string) (int, error) {
	total := 0

	n, err := io.WriteString(w, title+"\n")
	total += n
	if err != nil {
		return total, err
	}

	for _, item := range items {
		n, err = io.WriteString(w, "- "+item+"\n")
		total += n
		if err != nil {
			return total, err
		}
	}
	return total, nil
}
```

## Walkthrough

`failWriter{after: 1}` accepts the title (2 bytes) and rejects the next write. `WriteReport` returns `2, errBoom` — the count reflects what actually reached the sink.

## Pitfalls

- Ignoring the returned `n` and reporting `len(everything)` instead.
- Building the entire report with `+=` and writing once — correct output, but memory grows with the item count.
- Continuing after an error, which piles more failures onto a dead writer.
