# `iota` and constant blocks

## Intuition

Inside a `const (...)` block, `iota` is the index of the current *line*. It
resets to 0 at the start of every block and increments by one per ConstSpec:

```go
const (
	A = iota   // 0
	B = iota   // 1
	C = iota   // 2
)
```

Writing `= iota` on every line is noise, so Go lets you omit it: **a constant
line with no expression repeats the previous line's expression**, with `iota`
now holding the new index.

```go
const (
	A = iota   // 0
	B          // 1  — repeats "= iota"
	C          // 2
)
```

## Building scales

Because the repeated expression is evaluated fresh each line, any formula
involving `iota` becomes a sequence. The binary storage units are the canonical
example:

```go
const (
	_  = iota             // burn index 0
	KiB = 1 << (10 * iota) // 1 << 10 = 1024
	MiB                    // 1 << 20 = 1048576
	GiB                    // 1 << 30 = 1073741824
)
```

Three things to see there:

1. `_ = iota` **skips** index 0 without naming anything, so `KiB` lands on
   `iota == 1`. Without it, `KiB` would be `1 << 0` = 1.
2. `MiB` and `GiB` carry no expression — they repeat `1 << (10 * iota)`.
3. `1 << (10 * iota)` is a *shift*: multiply by 1024 per step. These are all
   powers of two, which is exactly what "binary prefix" means.

## Why derive instead of typing the numbers

`GiB = 1073741824` compiles fine and is wrong-looking to review, easy to mistype,
and says nothing about the pattern. The derived form documents the rule, and the
compiler does the arithmetic — constant expressions cost nothing at run time.

## Approach

1. Start the `const` block with `_ = iota` to burn slot 0.
2. Define `KiB = 1 << (10 * iota)`; the next lines repeat the expression for MiB and GiB.
3. `Bytes(n)` is `n * KiB`.

## Solution

```go
const (
	_   = iota
	KiB = 1 << (10 * iota)
	MiB
	GiB
)

func Bytes(n int) int {
	return n * KiB
}
```

## Walkthrough

At `KiB` iota is 1, so `1 << 10 = 1024`; MiB gets iota 2 → `1 << 20`; GiB iota 3 → `1 << 30`.

## Pitfalls

- `iota` counts **lines in the block**, not named constants. A line declaring
  three names (`A, B, C = iota, iota, iota`) still advances it by one.
- A blank line or comment does not advance `iota`; only ConstSpecs do.
- `iota` is only meaningful inside `const`. There is no `var` equivalent.
- Reset is per block: a second `const (...)` starts at 0 again.
- Skipping with `_` is idiomatic when index 0 would be a meaningless or unsafe
  value — for units, or for an enum where the zero value should mean "unset".

## KiB is not KB

`KiB` is 1024 bytes (binary); `KB` is 1000 bytes (decimal). Disk vendors use the
decimal one, operating systems usually the binary one — the reason a "500 GB"
drive shows up as ~465 GiB.
