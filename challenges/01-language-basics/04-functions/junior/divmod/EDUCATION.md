# Integer division and modulo

## The idea

For integers `/` yields the truncated quotient and `%` the remainder; together they satisfy `a == (a/b)*b + a%b`.

## Why it matters

Splitting a value into whole units and leftover (seconds into minutes+seconds, cents into dollars+cents) is a one-liner with both operators.

## Watch out

- Division truncates toward zero, so `-7/2 == -3` and `-7%2 == -1`.
- Dividing by zero panics; the task guarantees `b != 0`.
