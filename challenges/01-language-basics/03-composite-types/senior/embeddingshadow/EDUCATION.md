# Promotion and shadowing

## The idea

Embedding promotes a field: `e.ID` normally means `e.Base.ID`. But if the outer
struct declares its **own** `ID`, that one wins — the promoted field is shadowed
and must be qualified:

```go
return e.Base.ID // reach the embedded field explicitly
```

## Why it matters

Name collisions between embedded and outer fields are silent — `e.ID` compiles and
returns the outer value. Understanding promotion vs shadowing prevents reading the
wrong field.

## Watch out

- A shallower field shadows a deeper (promoted) one.
- Ambiguous promotions (two embeds with the same field) are a compile error until
  qualified.
- Methods promote and shadow by the same rules.
