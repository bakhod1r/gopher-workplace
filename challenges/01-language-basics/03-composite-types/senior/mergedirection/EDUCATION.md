# Override order in a merge

## The idea

When copying two maps into one, the **last** write to a key wins. For "override
on top of base", copy the base first, then the overrides:

```go
for k, v := range base { out[k] = v }
for k, v := range over { out[k] = v } // over wins
```

## Why it matters

Layered configuration (defaults < file < env < flags) depends entirely on copy
order. Reversing it makes defaults silently clobber user overrides — a real,
confusing config bug.

## Watch out

- Order of the two loops is the whole behavior.
- Within one map, iteration order is random but irrelevant here.
- Document precedence explicitly; it's easy to get backwards.
