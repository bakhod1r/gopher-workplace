# Tracking the new head after reversal

## The idea

The loop's final `prev` is the reversed list's head; the original head has become the tail with a nil Next.

## Why it matters

Returning the wrong end is a subtle reversal bug that yields a single node.

## Watch out

- After reversal, `head.Next == nil` (it's the tail).
- Return `prev`, the last node advanced.
