# Building pipelines from function slices

## The idea

Threading a value through a slice of functions gives a data-flow pipeline with the order fixed by slice order.

## Why it matters

Middleware chains and transform stages generalise this.

## Watch out

- Empty `fns` must return x unchanged (identity).
- Order is fns[0] first — opposite of mathematical Compose.
