# Reaching the predecessor for deletion

## The idea

Unlinking a node needs a handle on the node before it; walking all the way to the target leaves nothing to relink.

## Why it matters

Off-by-one traversal is the classic remove-at-index list bug.

## Watch out

- To delete index i, stop `prev` at i-1.
- Then `prev.Next = prev.Next.Next` skips the target.
