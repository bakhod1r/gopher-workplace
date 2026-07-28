# Binary search boundary updates

## The idea

After testing `mid`, the half that excludes it is `[mid+1, hi]` or `[lo, mid-1]`; reusing `mid` in a bound breaks the shrink invariant and can loop.

## Why it matters

Off-by-one bound updates are the canonical binary-search bug.

## Watch out

- `xs[mid] < target` ⇒ `lo = mid + 1`, never `lo = mid`.
- Every iteration must reduce `hi - lo`.
