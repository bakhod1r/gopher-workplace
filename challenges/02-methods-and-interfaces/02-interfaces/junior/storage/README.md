# Storage

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

An application persists key/value pairs to whichever backend is configured.

## Task

Implement the stub(s) in [storage.go](storage.go):

1. Implement `Put` and `Get` on `*MemStore` (`Get` returns the value and whether the key existed).
2. Implement `Copy`, which copies every listed key from one store to another and returns how many moved.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  s := NewMemStore(); s.Put("a", "1"); s.Get("a")
Output: "1", true
```

**Example 2:**

```
Input:  s.Get("missing")
Output: "", false
```

**Example 3:**

```
Input:  Copy(src, dst, []string{"a", "zz"})
Output: 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Store behind an interface** | Backends swap without touching callers. |
| 2 | **Comma-ok map read** | Reused from composite types: `v, ok := m[k]`. |
| 3 | **Constructor function** | Reused: `NewMemStore` initialises the map so writes do not panic. |

## Hint

Writing to a nil map panics — initialise it in `NewMemStore`.

## Validate

```bash
make verify
```
