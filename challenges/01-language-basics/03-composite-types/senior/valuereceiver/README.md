# Receiver Doesn't Mutate

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`Credit` has a **value** receiver, so `w.Balance += amount` updates a copy that's
discarded — the wallet never changes.

## Task

Fix the receiver between the markers in
[valuereceiver.go](valuereceiver.go) to a pointer.

## Examples

```go
w := &Wallet{100}; w.Credit(50) // Balance must become 150
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Value receiver** | Operates on a copy. |
| 2 | **Pointer receiver** | Mutates the original. |
| 3 | **Silent no-op** | Value-receiver mutation compiles but is lost. |

## Hint

`func (w *Wallet) Credit(amount int)`.

## Validate

```bash
make verify
```
