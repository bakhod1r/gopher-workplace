# Method Value

**Level:** middle
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A plugin system registers callbacks. Instead of wrapping method calls in
anonymous functions, pass the method value directly — cleaner and idiomatic.

## Task

Implement `ApplyMethod` in [methodval.go](methodval.go):

1. Return `g.Greet` as a bound method value (type `func(string) string`).
2. The returned function must capture the receiver's state at creation time.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  fn := ApplyMethod(Greeter{"Hello"}); fn("Alice")
Output: "Hello, Alice!"
```

**Example 2:**

```
Input:  fn := ApplyMethod(Greeter{"Hi"}); fn("Bob")
Output: "Hi, Bob!"
```

**Example 3:**

```
Input:  fn := ApplyMethod(Greeter{"Hey"}); fn("")
Output: "Hey, !"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Method values** | `g.Greet` (without parens) produces a `func(string) string` bound to `g`. |
| 2 | **Closure capture** | The method value captures a *copy* of the value receiver at creation time. |
| 3 | **Value receiver** | `Greeter` is copied when captured — later mutations to `g` don't affect the bound function. |

## Hint

A method value is just `g.Greet` — no parentheses, no wrapper function needed.

## Validate

```bash
make verify
```
