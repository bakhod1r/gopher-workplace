# Plugin Registry

**Level:** middle
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A build tool loads plugins that each declare a name and optionally a cleanup step.

## Task

Implement the stub(s) in [pluginreg.go](pluginreg.go):

1. Implement `Register` on `*Registry` — reject a duplicate name with `ErrDuplicate`.
2. Implement `RunAll`, which runs every plugin in registration order and returns the results.
3. Implement `CloseAll`, which calls `Close` only on the plugins that also implement `Closer`, returning how many were closed.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  r.Register(&Simple{N: "a"})
Output: nil
```

**Example 2:**

```
Input:  registering the same name twice
Output: ErrDuplicate
```

**Example 3:**

```
Input:  CloseAll with one closeable plugin of two
Output: 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Optional interfaces** | Assert to a second interface to discover extra capabilities. |
| 2 | **Registration order** | A slice preserves order; a map alone would not. |
| 3 | **Duplicate detection** | Reused: a set keyed by name. |

## Hint

`if c, ok := p.(Closer); ok { c.Close() }` — the classic optional-interface check.

## Validate

```bash
make verify
```
