# Ring-based list rotation

## The idea

Closing the list into a ring and breaking it at the right offset rotates in O(n); the modulo handles k larger than the length.

## Why it matters

Rotation appears in scheduling, buffers, and circular queues.

## Watch out

- Reduce k with `k % len` (guard len 0).
- Breaking the ring one node early sets the new head.
