# Statement ordering in pointer surgery

## The idea

Overwriting `cur.Next` before saving it loses the rest of the list; capture the successor first.

## Why it matters

Order-of-assignment bugs are common in in-place list/tree rewiring.

## Watch out

- Read `next := cur.Next` before `cur.Next = prev`.
- Reversed order truncates the list after the first node.
