# Skipping a matching prefix

## The idea

DropWhile locates the first non-matching element and returns the remainder; copying the tail avoids sharing the caller's backing array.

## Why it matters

Trimming leading whitespace/zeros/headers is DropWhile in disguise.

## Watch out

- Returning `xs[i:]` directly would alias the caller's array; copy it.
- If everything matches, the result is empty.
