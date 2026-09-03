# Longest Match That Returns The Shortest

## Intuition

The scan visits every prefix of `s` in increasing length. Returning on the first hit therefore yields the *shortest* registered match — exactly backwards. Recording the hit and letting the loop run gives the longest one.

## Approach

1. Track a best value and a found flag.
2. Walk every prefix length from 0 to len(s).
3. Overwrite the best on each hit, since later hits are longer.
4. Return the best after the loop.

## Solution

```go
func (p *Prefixes[V]) Longest(s string) (V, bool) {
	var best V
	found := false
	for i := 0; i <= len(s); i++ {
		if v, ok := p.m[s[:i]]; ok {
			best = v
			found = true
		}
	}
	return best, found
}

func (p *Prefixes[V]) Add(prefix string, v V) {
	if p.m == nil {
		p.m = make(map[string]V)
	}
	p.m[prefix] = v
}

func (p *Prefixes[V]) Len() int {
	return len(p.m)
}
```

## Walkthrough

With `""`, `/a` and `/a/b` registered, `Longest("/a/b/c")` matches `""` at i=0 and returns the catch-all, never reaching `/a/b`.

## Pitfalls

- Walking from the longest prefix down and returning on the first hit — correct, but the loop bounds are easy to get wrong.
- Starting the loop at `i = 1`, which silently disables the catch-all route.
- Slicing by bytes and assuming that is safe for multi-byte prefixes; it is, but only because prefixes align on the same boundaries.
