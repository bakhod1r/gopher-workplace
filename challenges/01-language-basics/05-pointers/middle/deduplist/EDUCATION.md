# In-place list compaction

## The idea

Relinking past equal successors removes duplicates without allocation; advance only when you didn't skip.

## Why it matters

Deduplicating sorted streams and lists uses adjacent-skip relinking.

## Watch out

- Don't advance `cur` when you just skipped — there may be more dups.
- Guard `cur.Next != nil` before comparing.
