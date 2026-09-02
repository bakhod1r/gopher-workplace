# Method Values

## Intuition

`x.M` is not a lazy reference to `x` — it evaluates the receiver right there and stores it in the closure. With a value receiver you get a frozen copy; with a pointer receiver you get the pointer, so later writes are visible through it.

## Approach

1. Give `Counter` pointer-receiver `Get`/`Set` and `ValCounter` a value-receiver `Get`.
2. `BindValue` returns `v.Get` — the copy is made at that moment.
3. `BindPointer` returns `c.Get`, capturing the pointer.
4. `GetExpr` returns `ValCounter.Get`, the method expression form.

## Solution

```go
func (c *Counter) Get() int { return c.N }

func (c *Counter) Set(n int) { c.N = n }

func (v ValCounter) Get() int { return v.N }

func BindValue(v ValCounter) func() int { return v.Get }

func BindPointer(c *Counter) func() int { return c.Get }

func GetExpr() func(ValCounter) int { return ValCounter.Get }
```

## Walkthrough

`TestCallbackRegistrationPitfall` binds three callbacks from one mutated variable. Each captures its own snapshot, so they return 1, 2, 3 — a pointer receiver would have made all three return 3.

## Pitfalls

- Expecting `v.Get` to track later writes to `v` — it never re-reads the variable.
- Registering callbacks from a value receiver in a loop and being surprised the values are frozen (or, with pointers, all identical).
- Confusing a method value (`x.M`, receiver bound) with a method expression (`T.M`, receiver as first argument).
