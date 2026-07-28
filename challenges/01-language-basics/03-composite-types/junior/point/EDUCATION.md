# Structs

## The idea

A struct groups named fields into one value:

```go
type Point struct { X, Y int }
p := Point{X: 1, Y: 2}
```

Fields are accessed with dot notation; the whole struct copies on assignment.

## Why it matters

Structs model records and domain types. Comparable structs (all comparable
fields) support `==` and can be map keys.

## Watch out

- Struct assignment copies all fields (value semantics).
- Use a pointer receiver / pointer to mutate a struct in a function.
- Field order affects memory layout (padding), not equality.
