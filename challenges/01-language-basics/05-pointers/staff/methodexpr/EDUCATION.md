# Method expressions

## The idea

`T.Method` (or `(*T).Method`) yields a function with the receiver as an explicit first parameter, unlike a method value which binds a specific receiver.

## Why it matters

Method expressions enable generic dispatch where the receiver varies per call.

## Watch out

- `(*Counter).Add` has signature `func(*Counter, int)`.
- A method VALUE (`c.Add`) binds one receiver; an EXPRESSION takes it as an argument.
