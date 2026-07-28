# Non-destructive BST insertion

## The idea

Insertion must recurse to a nil slot and add a node there; overwriting the root each call collapses the tree to one node.

## Why it matters

Overwriting instead of descending is a gross but easy-to-miss insert bug.

## Watch out

- `root = &Tree{Val: v}` discards everything already inserted.
- Recurse to the correct empty child instead.
