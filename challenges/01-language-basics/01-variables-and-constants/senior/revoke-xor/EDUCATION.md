# XOR toggles, AND-NOT clears

## The idea

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

## Why it matters

`Revoke` with `^` grants permissions to users who never had them — a security
bug that only shows when the target bit was already 0. Tests that only revoke
*present* flags miss it entirely.

## Watch out

- XOR is for toggling and for cheap swaps/checksums, not for clearing.
- `&^` is idempotent for clearing; XOR is its own inverse (applying twice
  restores).
- Always test the "clear an absent bit" case — that is where the two diverge.

## Try it yourself

```go
b := 0b0001
b ^ 0b0010  // 0b0011 (added)
b &^ 0b0010 // 0b0001 (unchanged)
b ^ 0b0001  // 0b0000 (removed, since present)
```
