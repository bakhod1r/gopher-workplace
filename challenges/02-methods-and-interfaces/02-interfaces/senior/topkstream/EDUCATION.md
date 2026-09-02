# Top-K Stream

## Intuition

Top-k does not need a sorted stream, only a sorted *window*. Once the window is full, most elements are rejected by a single comparison against its smallest member.

## Approach

1. Reject immediately when the window is full and `v` does not beat the last element.
2. Find the insertion position by scanning the descending window.
3. Grow the slice by one only while below `K`.
4. Shift right with `copy` and store `v` at the position.

## Solution

```go
func (t *TopK) Add(v int) {
	if t.K <= 0 {
		return
	}
	if len(t.vals) == t.K && v <= t.vals[len(t.vals)-1] {
		return
	}

	pos := len(t.vals)
	for i, existing := range t.vals {
		if v > existing {
			pos = i
			break
		}
	}

	if len(t.vals) < t.K {
		t.vals = append(t.vals, 0)
	}
	copy(t.vals[pos+1:], t.vals[pos:])
	t.vals[pos] = v
}

func (t *TopK) Result() []int { return t.vals }
```

## Walkthrough

`TestMemoryBoundedByK` streams 200k values and asserts both `len` and `cap` of the window stay at 10 — the capacity check is what proves nothing was buffered along the way.

## Pitfalls

- Appending everything and sorting at the end: correct, and O(n) memory.
- Growing past `K` before the shift, which silently keeps k+1 values.
- Using `append` to shift, which reallocates and breaks the capacity bound.
