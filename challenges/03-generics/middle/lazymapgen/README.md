# Usable Zero Value

**Level:** middle  
**Topic:** 03-generics

## Context

Every user of the earlier map-backed types had to remember a constructor. Making the zero value work removes a whole class of nil-map panics.

## Task

Implement the stub(s) in [lazymapgen.go](lazymapgen.go):

1. Implement `Set`, `Get`, and `Len` so that `var s Store[string,int]` works immediately.
2. Allocate the map lazily inside `Set`.
3. `Get` and `Len` must be safe on a nil map.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  var s Store[string,int]; s.Set(a,1)
Output: no panic
```

**Example 2:**

```
Input:  Get on a fresh store
Output: zero, false
```

**Example 3:**

```
Input:  Len on a fresh store
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Usable zero values** | Idiomatic Go: `sync.Mutex` and `bytes.Buffer` work without constructors. |
| 2 | **Reads of nil maps are safe** | Only writes panic — which is why only `Set` needs the guard. |
| 3 | **Fewer constructors** | Lazy allocation removes the "forgot to call New" failure mode. |

## Hint

Only `Set` needs the nil check; reading a nil map is already safe.

## Validate

```bash
make verify
```
