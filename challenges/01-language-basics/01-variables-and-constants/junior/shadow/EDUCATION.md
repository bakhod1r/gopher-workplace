# Scope and shadowing

## Intuition

A variable exists from its declaration to the end of the **block** that contains
it — the nearest enclosing pair of braces. Inside a nested block you may declare
a new variable with the same name. It does not overwrite the outer one; it
*hides* it for the rest of that block. That is shadowing.

```go
sum := 0
for _, n := range nums {
	sum := sum + n     // NEW sum, born and dead each iteration
}
// sum is still 0
```

The inner `sum` is initialised from the outer one, then thrown away when the
iteration ends. The accumulator never moves.

## `:=` declares, `=` assigns

That single character is the whole difference:

```go
sum := sum + n   // declare a new sum in this block
sum = sum + n    // assign to the existing sum
sum += n         // same, shorter
```

`:=` requires at least one new name on the left. That rule is what makes the
following legal and confusing:

```go
v, err := f()    // both new
w, err := g()    // w is new, err is REUSED — not shadowed, same variable
```

Reuse only happens in the *same* block. Inside a new block, `err` would be a
fresh, shadowed variable — the reason "the error I set is nil outside the `if`"
is such a common bug.

## Where blocks come from

Every `{}` opens a block, and some statements open one implicitly:

```go
if v := compute(); v > 0 {   // v lives in the if, including its else
	...
} else {
	...                      // v visible here
}
// v gone here

for i := 0; i < n; i++ {     // i belongs to the for
}
// i gone here
```

The `if`/`for`/`switch` init statement is scoped to the whole statement — a
useful way to keep a temporary from leaking into the function.

## Why Go allows it at all

Shadowing keeps short names usable in small scopes without fear of collision,
and it makes blocks self-contained. The cost is exactly the bug above, which is
why `go vet -shadow` (and linters like `govet`'s shadow check) exist.

## Approach

1. Accumulate into a `sum` declared before the loop.
2. Add each positive value.
3. Avoid re-declaring `sum` inside the loop (which would shadow it).

## Solution

```go
func Tally(nums []int) int {
	sum := 0
	for _, n := range nums {
		if n > 0 {
			sum += n
		}
	}
	return sum
}
```

## Walkthrough

For `[1 -2 3]`, 1 and 3 are positive, so `sum` reaches 4; negatives are skipped.

## Pitfalls

- Unused *variables* are a compile error, but a shadowed variable that you do
  assign is "used" — so the compiler will not save you here.
- The inner declaration may have a different type: `x := "1"` inside a block
  where the outer `x` is an `int` compiles fine.
- Package-level names can be shadowed too, including built-ins. `len := 3` is
  legal and disables `len(...)` for that block.
