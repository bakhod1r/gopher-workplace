# Look-ahead loop bounds

## The idea

When the body reads `xs[i+1]`, the loop must terminate at `len(xs)-1`; the classic `< len(xs)` bound over-runs by one.

## Why it matters

Adjacent-pair scans (diffs, transitions, merges) all need this reduced bound.

## Watch out

- Reading `xs[i+1]` ⇒ stop at `len(xs)-1`.
- With 0 or 1 elements the loop must not run.
