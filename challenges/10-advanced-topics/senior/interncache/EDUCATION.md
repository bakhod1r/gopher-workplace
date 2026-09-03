# Store One Copy Of Each Repeated String

## Intuition

Interning trades a map lookup for an allocation, which pays off when repeats dominate. The subtlety is that the lookup can be free — a string view built to hash and compare never escapes — while the stored copy cannot be.

## Approach

1. Return `""` for an empty input.
2. Build a borrowed string view and look it up; return the stored string on a hit.
3. On a miss, copy with `string(b)`, store it as both key and value, and return it.

## Solution

```go
import "unsafe"

// Pool holds one canonical string per distinct byte sequence.
type Pool struct {
	m map[string]string
}

// Len reports how many distinct strings the pool holds.
func (p *Pool) Len() int { return len(p.m) }

// Intern returns a string with b's contents, reusing a previously stored
// one when the same bytes have been seen before.
//
// Repeated values then share one allocation instead of one each, and a
// repeat lookup must not allocate at all.
//
// Examples:
//
// 	p.Intern([]byte("a")) twice => the same string, one allocation
func (p *Pool) Intern(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	view := unsafe.String(unsafe.SliceData(b), len(b))
	if s, ok := p.m[view]; ok {
		return s
	}
	owned := string(b)
	if p.m == nil {
		p.m = make(map[string]string)
	}
	p.m[owned] = owned
	return owned
}
```

## Walkthrough

The first `Intern` of a token copies it once and stores it. Every later call with the same bytes hashes the borrowed view, finds the entry, and returns the stored string — no allocation at all.

## Pitfalls

- Storing the borrowed view as the key, which the reuse test exists to catch.
- Storing the key and the value as two separate copies.
- Letting the pool grow without bound; a real interner needs an eviction rule.
