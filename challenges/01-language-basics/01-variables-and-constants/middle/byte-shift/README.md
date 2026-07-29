# Binary Byte Units

**Level:** middle
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

Storage sizes are powers of 1024. `iota` plus a left shift generates KB..TB in
one block — the classic Go idiom.

## Task

In [units.go](units.go):

1. Define `KB, MB, GB, TB` as `1 << (10 * iota)` (skip the `iota==0` slot with
   `_`).
2. Implement `Humanize(n)` returning the largest unit dividing `n` evenly and
   its symbol; `(n, "B")` below KB.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  KB
Output: 1024
```

**Example 2:**

```
Input:  Humanize(1048576)
Output: 1, "MB"
```

**Example 3:**

```
Input:  Humanize(512)
Output: 512, "B"
```

_Explanation:_ Below KB, reported in bytes.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **iota with shift** | `1 << (10*iota)` gives 1,1024,1048576,… when the first slot is skipped. |
| 2 | **Blank identifier in const** | `_ = iota` consumes value 0 so KB lands on iota==1. |
| 3 | **Untyped→typed** | `ByteSize` typing flows through the shift expression. |

## Hint

Only the first line needs the full `1 << (10 * iota)` expression; the rest
repeat it implicitly as iota climbs.

## Validate

```bash
make verify
```
