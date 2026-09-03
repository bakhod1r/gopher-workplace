# Did It Panic

**Level:** senior
**Topic:** 04-error-handling

## Context

A test helper checks that a function panics as documented, and hands the payload back so the caller can assert on it.

## Task

Implement `Panicked` in [assertpanic.go](assertpanic.go):

1. Return the recovered value and `true` when `f` panics.
2. Return `nil, false` when `f` returns normally.
3. Never let the panic escape.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Panicked(func() { panic("x") })
Output: "x", true
```

**Example 2:**

```
Input:  Panicked(func() {})
Output: nil, false
```

**Example 3:**

```
Input:  Panicked(func() { panic(nil) })
Output: a non-nil value, true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **recover returns any** | The payload is returned untouched. |
| 2 | **panic(nil)** | Go 1.21+ substitutes a *runtime.PanicNilError. |
| 3 | **Named results in a helper** | The deferred closure reports the outcome. |

## Hint

Go no longer lets `panic(nil)` be recovered as nil — the runtime substitutes a real value.

## Validate

```bash
make verify
```
