# Say Hello

## Intuition

The greeting word varies; the sentence shape does not. Put the varying part behind the interface and write the shape once.

## Approach

1. `English.Hello` returns `"Hello"`.
2. `Uzbek.Hello` returns `"Salom"`.
3. `Greet` returns `g.Hello() + ", " + name`.

## Solution

```go
func (e English) Hello() string { return "Hello" }

func (u Uzbek) Hello() string { return "Salom" }

func Greet(g Greeter, name string) string { return g.Hello() + ", " + name }
```

## Walkthrough

`Greet(Uzbek{}, "Ali")` dispatches to `Uzbek.Hello` (`"Salom"`), then appends `", Ali"`.

## Pitfalls

- Missing the space after the comma.
- Hardcoding the greeting in `Greet` instead of calling the method.
- Special-casing an empty name — `"Hello, "` is the expected output.
