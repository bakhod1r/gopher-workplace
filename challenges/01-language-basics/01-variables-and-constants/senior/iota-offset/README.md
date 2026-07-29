# Off-By-One iota Flags

**Level:** senior
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

A permission block compiles, `Has` looks right, yet every flag is double what it
should be — masks silently overlap. The bug is one token in the iota expression.

## Task

Fix the single line between the `CHANGE CODE` markers in [perms.go](perms.go) so
`Read=1, Write=2, Execute=4`. Do not touch anything else.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Read, Write, Execute
Output: 1, 2, 4
```

**Example 2:**

```
Input:  Has(Read|Execute, Read)
Output: true
```

**Example 3:**

```
Input:  Has(Write, Read)
Output: false
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **iota base** | The first line's expression seeds the whole block. |
| 2 | **`1 << iota`** | Correct shift starts at `1 << 0 == 1`. |
| 3 | **Bit overlap** | Off-by-one shifts make masks collide, breaking `Has`. |

## Hint

`1 << (iota + 1)` starts Read at 2. Drop the `+ 1`.

## Validate

```bash
make verify
```
