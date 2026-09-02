# Auth Check

## Intuition

An interface is a contract of method names. Any type with those methods satisfies it automatically. `CanEnter` never learns which concrete type it got — it only calls `Allow`.

## Approach

1. `Token.Allow`: return `t.Value != "" && t.Value != "expired"`.
2. `Guest.Allow`: return `false`.
3. `CanEnter`: return `a.Allow()`.

## Solution

```go
func (t Token) Allow() bool {
	return t.Value != "" && t.Value != "expired"
}

func (g Guest) Allow() bool { return false }

func CanEnter(a Authenticator) bool { return a.Allow() }
```

## Walkthrough

`CanEnter(Token{Value: "abc"})` stores `Token` in the interface value. The call `a.Allow()` dispatches to `Token.Allow`, which sees a non-empty, non-expired value and returns `true`.

## Pitfalls

- Type-switching inside `CanEnter` instead of calling the method — that defeats the interface.
- Using a pointer receiver on `Token` — then `Token{}` (a value) no longer satisfies `Authenticator`.
- Forgetting the empty-string case.
