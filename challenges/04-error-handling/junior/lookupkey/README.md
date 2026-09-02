# Map Lookup

**Level:** junior
**Topic:** 04-error-handling

## Context

A pricing table is keyed by SKU. A missing SKU is a real failure, not a zero price.

## Task

Implement `Lookup` in [lookupkey.go](lookupkey.go):

1. Return the value stored under `key` and a nil error.
2. Return `0` and `ErrNotFound` when the key is absent.
3. Treat a nil map as containing nothing.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Lookup(map[string]int{"a": 1}, "a")
Output: 1, nil
```

**Example 2:**

```
Input:  Lookup(map[string]int{"a": 1}, "z")
Output: 0, ErrNotFound
```

**Example 3:**

```
Input:  Lookup(nil, "a")
Output: 0, ErrNotFound
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Comma-ok idiom** | `v, ok := m[k]` distinguishes missing from zero. |
| 2 | **Zero value trap** | A missing key yields 0, which is also a legal price. |
| 3 | **Nil map reads** | Reading from a nil map is safe and always misses. |

## Hint

A stored value of 0 and a missing key look identical without the second return value.

## Validate

```bash
make verify
```
