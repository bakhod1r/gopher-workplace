# Deleting from a slice, order-preserving

## The idea

Removing index `i` while keeping order joins the parts before and after it:

```go
return append(xs[:i], xs[i+1:]...)
```

The `...` spreads the tail as individual arguments.

## Why it matters

The standard "delete keeping order" idiom (`slices.Delete` generalizes it). If
order doesn't matter, swapping the last element in is O(1).

## Watch out

- It **overwrites** `xs`'s backing array — copy first if the caller's slice
  matters.
- Guard `i` in `[0, len)` or the slice expressions panic.
- The leftover tail element lingers in the backing array (a leak for pointers).
