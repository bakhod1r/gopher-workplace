# Interface Map

**Level:** middle
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A command dispatcher maps names to command objects instead of running a long switch.

## Task

Implement the stub(s) in [ifacemap.go](ifacemap.go):

1. Implement `Register` and `Run` on `*Registry` (`Run` returns `ErrUnknown` for an unregistered name).
2. Implement `Names`, which returns the registered names in sorted order.
3. Registering an existing name replaces the command.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  r.Register("up", upper); r.Run("up", "hi")
Output: "HI", nil
```

**Example 2:**

```
Input:  r.Run("nope", "hi")
Output: "", ErrUnknown
```

**Example 3:**

```
Input:  r.Names() after registering "b" then "a"
Output: ["a", "b"]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **map of interface values** | Dispatch tables replace ever-growing switches. |
| 2 | **Comma-ok lookup** | Reused: distinguish an absent key from a zero value. |
| 3 | **sort.Strings** | Reused: map iteration order is random, so sort before returning. |

## Hint

Store `map[string]Command` and look up with the comma-ok form.

## Validate

```bash
make verify
```
