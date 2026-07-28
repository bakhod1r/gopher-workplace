# Defer mutating the return value

## The idea

Because deferred functions execute after the return value is assigned, a deferred closure over a named result can transform it (wrap errors, add context).

## Why it matters

Idiomatic error decoration (`defer func(){ err = wrap(err) }()`) depends on this timing.

## Watch out

- Only NAMED returns are visible to the deferred closure; a bare `return x` on an unnamed result can't be adjusted.
- The mutation happens before the caller resumes.
