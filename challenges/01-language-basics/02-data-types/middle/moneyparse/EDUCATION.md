# Money as integer minor units

## The idea

Represent currency as an integer number of cents, never a float. Parse the
dollars and the (optional) 1–2 fraction digits separately, pad the fraction to
two places, and combine: `dollars*100 + cents`.

## Why it matters

Floats cannot represent `0.10` exactly, so summing float dollars drifts.
Real payment and accounting systems parse and store integer minor units. This is
the parsing half of that discipline.

## Watch out

- `"3.5"` means 350 cents — pad a single fraction digit to two.
- Reject 3+ fraction digits rather than silently rounding.
- No `.` at all is valid ("7" → 700); a trailing `.` is not.
