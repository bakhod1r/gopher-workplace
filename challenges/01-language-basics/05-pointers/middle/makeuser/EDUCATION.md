# Constructing struct pointers

## Intuition

`&T{...}` allocates and initialises in one expression; it's the idiomatic constructor, equivalent to `new(T)` followed by field sets.

## Approach

1. Build a composite literal `&User{...}`.
2. Return its address; the struct escapes to the heap and outlives the call.

## Solution

```go
type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}
```

## Walkthrough

`NewUser("ann", 30)` allocates a `User`, fills both fields, and returns a pointer the caller keeps.

## Pitfalls

- `&User{...}` is clearer than `new(User)` + assignments.
- Each call allocates a distinct instance.
