# Immediately-invoked function expressions

## The idea

Appending `()` to a function literal runs it at once, yielding a value — handy for scoped, complex initialisation; omitting the call leaves you with the function itself.

## Why it matters

Forgetting the invoking `()` is a real bug that compiles but returns/stores a func instead of a value.

## Watch out

- `x := func(){...}` stores a func; `x := func(){...}()` stores its result.
- IIFEs keep setup logic local without a named helper.
