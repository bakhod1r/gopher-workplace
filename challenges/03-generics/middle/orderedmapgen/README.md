# Insertion-Ordered Map

**Level:** middle  
**Topic:** 03-generics

## Context

A config renderer must emit fields in the order the file declared them, which a plain Go map cannot do.

## Task

Implement the stub(s) in [orderedmapgen.go](orderedmapgen.go):

1. Implement `NewOrdered`, `Set`, `Get`, and `Keys`.
2. Re-setting an existing key updates the value but keeps its original position.
3. `Keys` returns a copy in insertion order.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Set(a,1); Set(b,2); Keys()
Output: [a b]
```

**Example 2:**

```
Input:  Set(a,1); Set(a,2); Keys()
Output: [a]
```

**Example 3:**

```
Input:  Get(missing)
Output: zero, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Ordered iteration** | Go maps randomise range order; a side list is the standard fix. |
| 2 | **Update versus insert** | Only a genuinely new key extends the order list. |
| 3 | **Defensive copies** | Handing out internal storage lets callers corrupt the structure. |

## Hint

Append to `keys` only when the key is not already present.

## Validate

```bash
make verify
```
