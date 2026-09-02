# Func Adapter

**Level:** middle
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

Most handlers are plain functions, but the router only accepts an interface. An adapter type bridges the two.

## Task

Implement the stub(s) in [funcadapt.go](funcadapt.go):

1. Implement `Handle` on `HandlerFunc` so a function value satisfies `Handler`.
2. Implement `Run`, which invokes a handler.
3. Implement `Chain`, which returns a handler applying every handler left to right.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Run(HandlerFunc(strings.ToUpper), "hi")
Output: "HI"
```

**Example 2:**

```
Input:  Chain(upper, exclaim).Handle("hi")
Output: "HI!"
```

**Example 3:**

```
Input:  Chain().Handle("hi")
Output: "hi"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Func adapter type** | `type HandlerFunc func(string) string` can carry methods — the `http.HandlerFunc` pattern. |
| 2 | **Interface satisfied by a function** | The method body simply calls the receiver. |
| 3 | **Closures** | Reused: `Chain` returns a handler that closes over the slice. |

## Hint

`func (f HandlerFunc) Handle(s string) string { return f(s) }` — the receiver *is* the function.

## Validate

```bash
make verify
```
