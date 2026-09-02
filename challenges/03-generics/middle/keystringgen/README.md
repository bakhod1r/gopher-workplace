# String-Like Map Keys

**Level:** middle  
**Topic:** 03-generics

## Context

Headers arrive keyed by a named `HeaderName` type with inconsistent casing, and the downstream API wants plain lowercase strings.

## Task

Implement the stub(s) in [keystringgen.go](keystringgen.go):

1. Implement `Normalize`, returning a map keyed by lowercase plain strings.
2. Return an empty (non-nil) map for an empty or nil input.
3. Later keys that collide after lowercasing may overwrite earlier ones.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Normalize(map[HeaderName]int{"A": 1})
Output: map[string]int{"a": 1}
```

**Example 2:**

```
Input:  keys differing only in case
Output: collapse to one entry
```

**Example 3:**

```
Input:  Normalize(nil)
Output: {}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Constrained key parameters** | `K ~string` is both comparable and convertible to `string`. |
| 2 | **Changing the key type** | The result's key type is concrete, so only the input side is generic. |
| 3 | **Collisions after normalisation** | Case folding can merge keys — document it rather than hiding it. |

## Hint

`string(k)` converts; `strings.ToLower` normalises.

## Validate

```bash
make verify
```
