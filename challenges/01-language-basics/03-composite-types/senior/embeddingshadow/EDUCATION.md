# Promotion and shadowing

## Intuition

Embedding promotes a field: `e.ID` normally means `e.Base.ID`. But if the outer
struct declares its **own** `ID`, that one wins — the promoted field is shadowed
and must be qualified:

```go
return e.Base.ID // reach the embedded field explicitly
```

## Approach

1. Bug: e.ID resolves to the outer shadowing field, not the embedded Base.ID. 2. When an outer struct declares a field with the same name as an embedded field, the outer one wins for the unqualified selector. 3. Fix: qualify the path: return e.Base.ID.

## Solution

```go
type Base struct {
	ID int
}

type Entity struct {
	Base
	ID int // shadows Base.ID
}

func BaseID(e Entity) int {
	return e.Base.ID
}
```

## Walkthrough

Entity{Base:{ID:7}, ID:99}: e.ID -> 99 (outer). e.Base.ID -> 7 (embedded). Task wants the embedded value, so use e.Base.ID.

## Pitfalls

- A shallower field shadows a deeper (promoted) one.
- Ambiguous promotions (two embeds with the same field) are a compile error until
  qualified.
- Methods promote and shadow by the same rules.
