# Method Expressions

## Intuition

A **method expression** `T.Method` produces a function where the receiver
becomes an explicit first parameter. Unlike a method value (`v.Method`), no
receiver is bound — you pass it each time.

| Form | Type | Receiver |
|------|------|----------|
| `a.Add` | `func(int) int` | bound to `a` |
| `Adder.Add` | `func(Adder, int) int` | explicit first arg |

## Approach

1. Call `fn(a, n)`.

## Solution

```go
func CallExpr(fn func(Adder, int) int, a Adder, n int) int {
	return fn(a, n)
}
```

## Walkthrough

`CallExpr(Adder.Add, Adder{10}, 5)`:
- `fn` is `Adder.Add`.
- `fn(Adder{10}, 5)` → `Adder{10}.Add(5)` → `15`.

## Pitfalls

- Confusing method values (`a.Add`) with method expressions (`Adder.Add`) —
  they have different signatures.
- For pointer receiver methods, the expression is `(*T).Method` with
  `func(*T, args) ret`.
