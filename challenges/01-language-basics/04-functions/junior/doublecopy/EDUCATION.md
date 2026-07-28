# Call by value

## The idea

Go copies each argument into the parameter; mutating a non-pointer parameter cannot affect the caller.

## Why it matters

It makes functions easier to reason about — no hidden mutation of inputs unless you pass a pointer or slice.

## Watch out

- Multiply before dividing (`price*rate/100`) to avoid truncating rate/100 to 0.
- Slices/maps/pointers are the exceptions that DO share underlying data.
