# Field access through struct pointers

## The idea

Go implicitly dereferences a struct pointer for field selection, so `c.Count` reads/writes the pointee's field.

## Why it matters

Methods with pointer receivers mutate struct state exactly this way.

## Watch out

- No need to write `(*c).Count`; `c.Count` is equivalent.
- A value receiver would edit a copy, not the caller's struct.
