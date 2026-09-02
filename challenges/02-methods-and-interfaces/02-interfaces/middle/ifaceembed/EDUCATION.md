# Embedded Interface

## Intuition

Embedding an interface gives the wrapper the whole method set for free. You override only the methods you want to instrument; everything else forwards automatically.

## Approach

1. `Gauge.Value` returns `g.N`.
2. `(*CountingSource).Value` increments `Calls`, then returns `c.Source.Value()`.
3. `ReadTwice` returns two calls to `s.Value()`.

## Solution

```go
func (g Gauge) Value() int { return g.N }

func (c *CountingSource) Value() int {
	c.Calls++
	return c.Source.Value()
}

func ReadTwice(s Source) (int, int) {
	return s.Value(), s.Value()
}
```

## Walkthrough

Wrapping a counter in a counter works because the outer `Value` calls the *embedded* `Source.Value`, which is the inner counter — each layer records exactly one call.

## Pitfalls

- Calling `c.Value()` inside the override — infinite recursion and a stack overflow.
- Leaving `Source` nil: the promoted method panics with a nil interface call.
- A value receiver on the override, so `Calls` never advances.
