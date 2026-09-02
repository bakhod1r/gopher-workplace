# Factory Method

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

Callers want *a* `Store` without knowing which concrete type they get. A factory
method centralizes that decision: one place maps a string key to a concrete
implementation, and everything downstream talks to the interface.

## Task

Implement `Create` on `StoreFactory` in [factorymeth.go](factorymeth.go):

1. `"mem"` returns a `MemStore`.
2. `"disk"` returns a `DiskStore`.
3. Anything else returns `nil`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  StoreFactory{}.Create("mem")
Output: MemStore{} (as a Store)
```

**Example 2:**

```
Input:  StoreFactory{}.Create("disk")
Output: DiskStore{} (as a Store)
```

**Example 3:**

```
Input:  StoreFactory{}.Create("s3")
Output: nil
```

_Explanation:_ unknown keys are not an error here — the contract says nil.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Method returning an interface** | The static type is `Store`; the dynamic type is chosen at runtime. |
| 2 | **Value receivers satisfy interfaces** | `MemStore` has `Save` on a value receiver, so `MemStore{}` itself is a `Store`. |
| 3 | **Type assertion in tests** | `f.Create("mem").(MemStore)` recovers the dynamic type — that is what is being checked. |

## Hint

A `switch storeType` with three arms. Return `MemStore{}`, not `&MemStore{}` —
the test asserts on the value type.

## Validate

```bash
make verify
```
