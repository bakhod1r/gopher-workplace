# Protection Proxy

## Intuition

A proxy has the real object's shape but not its job. Callers cannot tell the
difference; the proxy decides whether the real work happens at all. Variants
differ only in what that decision is: access control here, laziness in a virtual
proxy, network transport in a remote one.

## Approach

1. Evaluate the policy first.
2. Delegate only on success.
3. Return a substitute result otherwise.

## Solution

```go
func (p *Proxy) Do() string {
	if p.role != "admin" {
		return "denied"
	}
	return p.w.Do()
}
```

## Walkthrough

`p1` holds role `"user"`, so the guard returns immediately and `w.calls` stays
0 — this is the assertion that distinguishes a real proxy from one that merely
rewrites the return value. `p2` holds `"admin"`, so `p.w.Do()` runs, increments
`calls` to 1, and its `"done"` is passed through untouched.

Both proxies share one `*Worker`, so the counter is a single global witness to
how many calls got through.

## Pitfalls

- **Calling first, checking after.** `res := p.w.Do(); if ... return "denied"`
  passes the string assertions and fails the counter — the side effect already
  happened.
- **Comparing roles loosely.** A `strings.Contains` check would let
  `"not-admin"` through.
- **Nil worker with an admin role.** Real code validates the target in a
  constructor; here the test always supplies one.
