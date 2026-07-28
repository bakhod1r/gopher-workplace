# Double indirection

## The idea

A `**int` lets a callee reseat the caller's pointer variable; `*pp = q` stores a new address into it.

## Why it matters

Insert/delete on linked lists and trees reseats parent pointers via double pointers.

## Watch out

- Reassigning a plain `*int` parameter only changes the local copy.
- `**int` reaches the caller's pointer.
