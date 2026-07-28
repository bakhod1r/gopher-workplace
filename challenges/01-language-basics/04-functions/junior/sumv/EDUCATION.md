# Variadic functions

## The idea

The final parameter written `...T` receives a `[]T` built from the trailing call arguments; passing `slice...` forwards an existing slice instead of copying element by element.

## Why it matters

Variadic APIs (`fmt.Println`, `append`) read naturally at the call site while still giving the body a plain slice to iterate.

## Watch out

- Inside the function the parameter is a real slice; a nil/empty one is valid.
- You can pass either loose args OR one spread slice, never both.
