# Factory Method

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Task

Implement `Create` in [factorymeth.go](factorymeth.go):

1. Switch on `storeType`.
2. Return `MemStore{}` for `"mem"`.
3. Return `DiskStore{}` for `"disk"`.
4. Return `nil` otherwise.

## Validate

```bash
make verify
```
