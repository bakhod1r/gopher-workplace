# Count Fields Without Splitting Anything

## Intuition

Splitting materialises every field so you can measure them. But the field count is just the separator count plus one, and the payload size is the length minus the separators — both are running totals over a single scan.

## Approach

1. Return 0, 0 for an empty line.
2. Start the field count at 1.
3. Scan the bytes: a separator increments the field count, anything else increments the size.

## Solution

```go
// CountFields returns how many sep-separated fields line holds and how
// many bytes those fields occupy in total, excluding the separators.
//
// Splitting builds a string and a slice of headers that are thrown away
// immediately. One scan over the bytes answers both questions.
//
// Examples:
//
// 	CountFields([]byte("ab,c"), ',') => 2, 3
func CountFields(line []byte, sep byte) (fields, size int) {
	if len(line) == 0 {
		return 0, 0
	}
	fields = 1
	for _, b := range line {
		if b == sep {
			fields++
			continue
		}
		size++
	}
	return fields, size
}
```

## Walkthrough

"ab,c" has one separator, so two fields, and three non-separator bytes. The split version would allocate a 4-byte string and a 2-element header slice to reach the same numbers.

## Pitfalls

- Starting the field count at 0, which is off by one for every non-empty line.
- Counting the separators in the size total.
