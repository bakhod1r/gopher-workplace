# Cloner

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A config service hands out copies of objects so callers cannot mutate the originals.

## Task

Implement the stub(s) in [cloner.go](cloner.go):

1. Implement `Clone` on `*Config` — return a new `*Config` whose `Tags` slice is independent of the original.
2. Implement `CopyOf`, which clones any `Cloner`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  c := &Config{Name: "db", Tags: []string{"a"}}; c.Clone().(*Config).Name
Output: "db"
```

**Example 2:**

```
Input:  d := c.Clone().(*Config); d.Tags[0] = "z"; c.Tags[0]
Output: "a"
```

**Example 3:**

```
Input:  CopyOf(&Config{Name: "x"}).(*Config).Name
Output: "x"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Interface with a self-returning method** | `Clone() Cloner` returns the interface, so callers assert back to the concrete type. |
| 2 | **Type assertion** | `v.(*Config)` recovers the concrete pointer from the interface value. |
| 3 | **Slice aliasing** | Reused from composite types: copying a struct copies the slice header, not the elements. |

## Hint

A struct copy shares the backing array. Build a fresh slice with `make` + `copy`.

## Validate

```bash
make verify
```
