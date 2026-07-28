# Named return values

## The idea

Declaring result names lets you assign them anywhere in the body and use a bare `return`; they default to zero values.

## Why it matters

They shine with defer-based error handling and make guard clauses concise.

## Watch out

- A bare `return` on the zero path yields (0,false) automatically.
- Overusing named returns hurts readability; reserve for defer/guard cases.
