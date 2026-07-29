# Method expressions

## Intuition

`T.Method` (or `(*T).Method`) yields a function with the receiver as an explicit first parameter, unlike a method value which binds a specific receiver.

## Approach

1. A **method expression** `(*Counter).Add` is a function whose first parameter is the receiver.
2. The bug ignores the passed receiver and mutates a captured local.
3. Return `(*Counter).Add`.

## Solution

```go
type Counter struct{ N int }

func (c *Counter) Add(d int) { c.N += d }

func AdderExpr() func(*Counter, int) {
	return (*Counter).Add
}
```

## Walkthrough

The bug's closure calls `c.Add` on its own hidden `c`, so `f(c, 5)` never touches the caller's counter. The method expression applies to whatever receiver you pass.

## Pitfalls

- `(*Counter).Add` has signature `func(*Counter, int)`.
- A method VALUE (`c.Add`) binds one receiver; an EXPRESSION takes it as an argument.
