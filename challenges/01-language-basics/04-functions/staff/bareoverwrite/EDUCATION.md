# Bare return with deferred adjustment

## The idea

A bare `return` returns whatever the named results currently hold; a deferred closure then transforms them. Forgetting to assign leaves the defer operating on zero.

## Why it matters

Understanding that defers act on the named result AFTER the return statement assigns it explains many defer/return interactions.

## Watch out

- A bare `return` on an unassigned named result yields its zero value.
- The deferred edit acts on that value — assign it first.
