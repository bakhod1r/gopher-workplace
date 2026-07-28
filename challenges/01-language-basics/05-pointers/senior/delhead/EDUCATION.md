# Returning the new head

## The idea

Deleting the head has no predecessor to relink, so the function must return the successor for the caller to adopt.

## Why it matters

Head-deletion bugs leave stale nodes at the front of lists/queues.

## Watch out

- Returning the original `head` keeps the deleted node; RETURN head.Next.
- The caller must use the returned head.
