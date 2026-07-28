# Node deletion by relinking

## The idea

Removing a node re-points its predecessor's Next to its successor; deleting the head yields a new head.

## Why it matters

Unlinking nodes is the core of list/queue/LRU mutation.

## Watch out

- Handle the head separately — it has no predecessor.
- Only the first match is removed.
