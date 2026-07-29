# Byte Units

**Level:** junior
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

A storage dashboard needs the binary size constants `KiB`, `MiB`, and `GiB`.
Declare them with an `iota` shift block so the values stay DRY instead of
hand-typed magic numbers.

## Task

Implement the `const` block in [byteunits.go](byteunits.go) so that:

1. `KiB == 1024`, `MiB == 1024*1024`, `GiB == 1024*1024*1024`.
2. The values stay derived from `iota` — do not hand-type each number.
3. `Bytes(n)` returns `n` whole KiB in bytes.

Do **not** change the function signature or the tests.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  KiB
Output: 1024
```

**Example 2:**

```
Input:  MiB
Output: 1048576 (1024*1024)
```

**Example 3:**

```
Input:  Bytes(2)
Output: 2048
```

_Explanation:_ 2 whole KiB in bytes.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **`iota`** | Resets to 0 at each `const` block and increments by 1 per line. |
| 2 | **Shift by iota** | `1 << (10 * iota)` gives 1024, 1048576, … *once iota starts at 1*. |
| 3 | **Skipping a value** | A leading `_ = iota` line burns index 0 so the first named constant lands on iota == 1. |

## Hint

`iota` starts at `0` in a `const` block. If `KiB` sits on `iota == 0` its shift
is `1 << 0 == 1`, and every later value inherits the same off-by-one. Burn the
zero index with a leading `_ = iota` so the first real constant lands on `1`.

## Validate

```bash
make verify
```
