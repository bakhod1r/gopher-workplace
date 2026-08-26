# Methods vs Functions

## Intuition

In Go, a **method** is a function with a special first parameter called the
*receiver*. It lets you associate behaviour with a type — `Person.Greet()` reads
better than `Greet(person)` and scopes the function to the type.

## Approach

1. Define a `Person` struct with a `Name` field.
2. Add a value receiver `(p Person)` before the function name.
3. Return the formatted greeting using string concatenation.

## Solution

```go
func (p Person) Greet() string {
	return "Hello, " + p.Name + "!"
}
```

## Walkthrough

For `Person{Name: "Alice"}`:
- `p.Name` is `"Alice"`.
- `"Hello, " + "Alice" + "!"` produces `"Hello, Alice!"`.

## Pitfalls

- Forgetting the receiver `(p Person)` makes it a standalone function, not a method.
- Using `fmt.Sprintf` works but is slower than plain `+` for simple concatenation.
- A value receiver copies the struct — fine here since we only read `Name`.
