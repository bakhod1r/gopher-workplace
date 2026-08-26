# Merge Object

**Level:** middle
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A config loader merges overrides from CLI flags. Only non-zero fields are
applied — zero means "not set".

## Task

Implement `Merge` on `*Config` in [mergeobj.go](mergeobj.go):

1. For each field in `other`: if non-zero, copy to `c`.
2. Zero-value fields in `other` are skipped.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  c := Config{"localhost", 8080, false}; c.Merge(Config{Port: 9090})
Output: c == Config{"localhost", 9090, false}
```

**Example 2:**

```
Input:  c := Config{"h", 80, false}; c.Merge(Config{Debug: true})
Output: c == Config{"h", 80, true}
```

**Example 3:**

```
Input:  c := Config{"h", 80, true}; c.Merge(Config{})
Output: c == Config{"h", 80, true} (no change)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Pointer receiver** | Mutation must persist. |
| 2 | **Zero values** | `""` for string, `0` for int, `false` for bool. |
| 3 | **Conditional update** | Check non-zero before overwriting. |

## Hint

`if other.Host != "" { c.Host = other.Host }` — and similarly for other fields.

## Validate

```bash
make verify
```
