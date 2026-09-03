# Call A Method By Name, Once Resolved

**Level:** staff
**Topic:** 10-advanced-topics / 03-reflection

## Context

A command dispatcher resolves handler methods by name on every request. `MethodByName` is a linear scan of the method table, and the table is not small.

## Task

Implement [methodcache.go](methodcache.go):

1. Call the named method on `v` with `arg` and return its int result.
2. Use the shared cache so a (type, name) pair is resolved at most once, as far as callers can tell.
3. Return `ErrMethod` for a missing method, a wrong signature, or a nil value.
4. A pointer-receiver method must not be reachable through the value type.
5. Correct under concurrent use.

Replace the stub body in [methodcache.go](methodcache.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  CallNamed(adder{2}, "Add", 3)
Output: 5, nil
```

**Example 2:**

```
Input:  CallNamed(adder{}, "Name", 1)
Output: ErrMethod
```

_Explanation:_ The result is a string.

**Example 3:**

```
Input:  CallNamed(ptrAdder{}, "Add", 1)
Output: ErrMethod
```

_Explanation:_ Pointer-receiver methods are not in the value's method set.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Method sets** | `T`'s method set excludes pointer-receiver methods; `*T`'s includes both. |
| 2 | **Method.Type includes the receiver** | Parameter 0 is the receiver, so a one-argument method has `NumIn() == 2`. |
| 3 | **Value.Method(i).Call** | The bound method value takes only the declared arguments. |
| 4 | **Negative caching** | Storing -1 for a failed lookup keeps the failure path as cheap as the success path. |

## Hint

The cache and the signature check are given. Resolve, guard, call, read the result.

## Validate

```bash
make verify
```
