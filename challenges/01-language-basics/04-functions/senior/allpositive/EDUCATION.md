# Early return for all/any checks

## The idea

"All" returns false at the first counterexample then true after the loop; "any" is the mirror. Returning true mid-loop answers the wrong question.

## Why it matters

Inverted quantifier logic passes on all-true inputs and fails only on mixed data — easy to ship.

## Watch out

- "All P" ⇒ fail fast on `!P`, succeed after the loop.
- "Any P" ⇒ succeed fast on `P`, fail after the loop.
