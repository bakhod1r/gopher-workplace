# Stream Filter

## Intuition

Predicates as values compose: `Not` wraps anything, and the pipeline stays one loop. Streaming means memory scales with the *result*, not the input.

## Approach

1. `Contains.Match` uses `strings.Contains`; `MinLen.Match` compares `len(line) >= m.N`.
2. `Not.Match` negates the inner predicate.
3. `FilterStream` loops `src.Next()` until `ok` is false, appending only matches.

## Solution

```go
func (c Contains) Match(line string) bool { return strings.Contains(line, c.Sub) }

func (m MinLen) Match(line string) bool { return len(line) >= m.N }

func (n Not) Match(line string) bool { return !n.Inner.Match(line) }

func FilterStream(src Source, p Predicate) []string {
	var out []string
	for {
		line, ok := src.Next()
		if !ok {
			return out
		}
		if p.Match(line) {
			out = append(out, line)
		}
	}
}
```

## Walkthrough

`Not{Inner: Not{...}}` composes because `Not` holds a `Predicate`, not a concrete type — the double negation matches again.

## Pitfalls

- Draining the source into a slice first, which defeats the streaming requirement.
- `len(line) > m.N`, which drops lines of exactly N bytes.
- Returning an empty non-nil slice where the test expects nil — here `nil` and length 0 both pass, but be deliberate.
