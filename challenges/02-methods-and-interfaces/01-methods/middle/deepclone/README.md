# Deep Clone

**Level:** middle
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A configuration system needs to fork configs. A shallow copy would share the
`Tags` slice — mutations to the fork would corrupt the original.

## Task

Implement `Clone` on `Config` in [deepclone.go](deepclone.go):

1. Copy all fields.
2. Deep-copy the `Tags` slice so mutations are isolated.
3. If `Tags` is nil, the clone's Tags should also be nil.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Config{"app", ["v1","prod"]}.Clone()
Output: Config{"app", ["v1","prod"]} (independent copy)
```

**Example 2:**

```
c := Config{"app", ["v1"]}; c2 := c.Clone()
c2.Tags[0] = "v2"
c.Tags[0] still == "v1"
```

**Example 3:**

```
Input:  Config{"x", nil}.Clone()
Output: Config{"x", nil}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Value receiver** | `Config` is copied — but slices share backing array. |
| 2 | **Shallow vs deep copy** | Value receiver copies the struct but not slice data. |
| 3 | **Slice cloning** | `copy(dst, src)` or `append([]T(nil), src...)`. |

## Hint

The value receiver already copies `Name`. For `Tags`, use
`copy` into a new slice or `append([]string(nil), c.Tags...)`.

## Validate

```bash
make verify
```
