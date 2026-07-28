# Mixing fixed and variadic parameters

## The idea

Only the last parameter may be variadic; fixed parameters bind positionally before it.

## Why it matters

Separator-style APIs stay ergonomic while the body still gets a plain slice of the trailing arguments.

## Watch out

- Prepending `sep` unconditionally leaks a leading separator.
- With zero parts the loop never runs and `""` is correct.
