# Encode Any Struct, Resolve It Once

**Level:** staff
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A form encoder walks the struct's field table on every call and builds a string per field. It is the slowest and the most allocating step in a request that does nothing else interesting.

## Task

Implement [encodecache.go](encodecache.go):

1. Append `name=value` for each exported string field, separated by `&`.
2. Use the cached layout so the field table is walked once per type.
3. Append into `dst`; with room, the call must allocate nothing.
4. Return `ErrKind` for anything that is not a struct.

Replace the stub body in [encodecache.go](encodecache.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Encode(nil, user{Name:"ann", Email:"a@b"})
Output: "Name=ann&Email=a@b"
```

**Example 2:**

```
Input:  Encode([]byte("pre:"), user{Name:"x"})
Output: "pre:Name=x&Email="
```

_Explanation:_ dst is extended.

**Example 3:**

```
Input:  16 goroutines x 200 encodes
Output: every result correct
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Cache the type, not the value** | Field indices are per type and never change. |
| 2 | **Append everything** | Writing into the caller's buffer removes the per-field string. |
| 3 | **sync.Map for a read-mostly cache** | Loads after the first write take no lock. |
| 4 | **Value.String without boxing** | It reads a string field directly, unlike `Interface()`. |

## Hint

The cache is written for you. Validate, fetch the layout, append.

## Validate

```bash
make verify
```
