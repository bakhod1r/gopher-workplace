# Guarding multi-step pointer advances

## The idea

Moving two steps requires both the current and next pointers to be non-nil; a single guard leaves the second hop unsafe.

## Why it matters

Missing the second guard is the classic slow/fast nil-panic bug.

## Watch out

- `fast.Next.Next` needs `fast.Next != nil` too.
- Check every dereference in a compound advance.
