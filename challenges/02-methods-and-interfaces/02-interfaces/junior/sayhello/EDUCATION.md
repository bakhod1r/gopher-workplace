# Implicit Interface Satisfaction

## Intuition

In Go, there is no `implements` keyword. A type satisfies an interface simply
by having all the methods the interface requires.

## Solution

```go
func (e English) Greet() string { return "Hello!" }
func (u Uzbek) Greet() string { return "Salom!" }
```

## Walkthrough

`SayHello(English{})` calls `English{}.Greet()` via the `Greeter` interface.
The compiler checks at compile time that `English` has a `Greet() string` method.
