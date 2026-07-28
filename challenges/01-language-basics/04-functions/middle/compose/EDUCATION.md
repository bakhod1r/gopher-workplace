# Composing functions

## The idea

A closure capturing two functions yields their composition; order matters — `f(g(x))` applies `g` first.

## Why it matters

Pipelines, decorators, and transformers chain small functions this way.

## Watch out

- `Compose(f,g)` is f-after-g, not g-after-f.
- Both captured functions live as long as the returned closure.
