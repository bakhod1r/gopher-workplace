# Context Up To The Cause

## Intuition

Nil and empty are two different answers here: one says the landmark was never found, the other says it was found with nothing above it. Collapsing them loses information the caller needs.

## Approach

1. Start with an empty non-nil slice.
2. Walk the chain, returning the collected messages when the target is reached.
3. Return nil when the walk ends without finding it.

## Solution

```go
out := []string{}
for e := err; e != nil; e = errors.Unwrap(e) {
	if e == target {
		return out
	}
	out = append(out, e.Error())
}
return nil
```

## Walkthrough

For a two-layer chain both wrapper messages are collected, and the loop stops before adding the target's own message.

## Pitfalls

- Starting with `var out []string`, so the found-immediately case returns nil.
- Including the target's message.
- Using `errors.Is` per link, which matches deeper and stops too early.
