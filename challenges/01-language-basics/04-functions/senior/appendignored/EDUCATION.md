# append returns a new header

## The idea

`append` can reallocate and always returns the (possibly new) slice header; ignoring the return keeps the old, shorter header.

## Why it matters

`_ = append(...)` or a bare `append(...)` call silently builds nothing.

## Watch out

- Always write `s = append(s, ...)`.
- The returned len/cap differ from the input whenever growth happens.
