# Deep-copying a linked list

## The idea

Copying only the head while reusing Next aliases the whole tail; recurse to allocate a fresh node per element.

## Why it matters

Shallow list copies that share the tail corrupt the source on edit.

## Watch out

- `Next: head.Next` shares the rest of the list.
- Use `Next: Copy(head.Next)` for a deep copy.
