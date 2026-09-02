# Middleware Stack

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

Cross-cutting concerns — logging, auth, timing — wrap a handler instead of
living inside it. Each `Middleware` takes the next `Handler` and returns a new
one. A `Stack` is an ordered list of them, and its `Then` method folds the whole
list around a handler, keeping `s[0]` outermost.

## Task

Implement `Then` on `Stack` in [middleware.go](middleware.go):

1. Start from `next`.
2. Apply the middlewares in **reverse** order, each wrapping the result so far.
3. Return the fully wrapped handler.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Stack{mw1}.Then(h); h = "H:"+req
Output: "1(H:req)1"
```

**Example 2:**

```
Input:  Stack{mw1, mw2}.Then(h)
Output: "1(2(H:req)2)1"
```

**Example 3:**

```
Input:  Stack{}.Then(h)
Output: "H:req"   (no middleware, handler untouched)
```

_Explanation:_ folding over an empty stack returns `next` unchanged.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Method on a named slice type** | `type Stack []Middleware` can carry methods; the receiver *is* the slice. |
| 2 | **Reverse fold** | Wrapping innermost-last is what makes `s[0]` run first at request time. |
| 3 | **Closures over the accumulator** | Each wrap captures the handler built so far. |

## Hint

```go
h := next
for i := len(s) - 1; i >= 0; i-- {
    h = s[i](h)
}
return h
```

Forward order compiles and gives `2(1(...)1)2` — the exact reversal the test catches.

## Validate

```bash
make verify
```
