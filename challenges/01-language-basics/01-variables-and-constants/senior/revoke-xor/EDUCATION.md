# XOR toggles, AND-NOT clears

## Intuition

`^` between two values is **XOR**: it flips each bit where the operand is set.
That is the wrong tool for "remove a flag", because flipping an *absent* bit
turns it on:

```go
set := Read           // Write bit is 0
set ^ Write           // Read | Write  <- XOR ADDED Write!
```

To clear, use bit-clear `&^`:

```go
set &^ Write          // removes Write if present, no-op if absent
```

## Approach

1. XOR toggles bits, so revoking an absent bit sets it.
2. Use AND-NOT (`&^`) to clear only.

## Solution

```go
type Permission uint8

const (
	Read Permission = 1 << iota
	Write
	Execute
)

func Revoke(set, drop Permission) Permission {
	return set &^ drop
}
```

## Walkthrough

`Revoke(Read, Write)` with XOR would add Write; `&^` leaves an absent bit untouched, returning Read.

## Pitfalls

- XOR is for toggling and for cheap swaps/checksums, not for clearing.
- `&^` is idempotent for clearing; XOR is its own inverse (applying twice
  restores).
- Always test the "clear an absent bit" case — that is where the two diverge.
