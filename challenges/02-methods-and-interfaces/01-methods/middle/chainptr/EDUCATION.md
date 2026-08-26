# Method Chaining

## Intuition

Method chaining means each method returns the receiver so you can write
`a.B().C().D()` instead of three separate statements. In Go, this requires a
pointer receiver that returns `*T`.

## Approach

1. Assign into the map.
2. Return `c`.

## Solution

```go
func (c *Config) Set(key, value string) *Config {
	c.Data[key] = value
	return c
}
```

## Walkthrough

`NewConfig().Set("host", "localhost").Set("port", "8080")`:
1. `NewConfig()` → `&Config{Data: {}}`.
2. `.Set("host", "localhost")` → stores, returns same pointer.
3. `.Set("port", "8080")` → stores, returns same pointer.

## Pitfalls

- Returning a value (`Config` not `*Config`) breaks the chain — each call gets
  a copy, and mutations are lost.
- Not initializing the map → nil map panic on `c.Data[key] = value`.
