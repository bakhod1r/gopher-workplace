# Struct Field Padding

**Level:** staff
**Topic:** 01-language-basics → 03-composite-types

## Context

Fields are laid out in declaration order with alignment padding. `bool, int64,
bool` forces the `int64` to an 8-byte boundary and pads the trailing bool, giving
24 bytes. Grouping the wide field first packs it to 16.

## Task

Reorder the fields between the markers in
[structpadding.go](structpadding.go) to minimize size (16 bytes on 64-bit).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  fields A bool, B int64, C bool
Output: size 24 (bloated)
```

**Example 2:**

```
Input:  fields B int64, A bool, C bool
Output: size 16 (minimal)
```

**Example 3:**

```
Input:  unsafe.Sizeof(Record{})
Output: 16
```

_Explanation:_ grouping the two bools after the int64 removes padding.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Alignment** | An int64 aligns to 8 bytes. |
| 2 | **Padding** | Gaps inserted to satisfy alignment. |
| 3 | **Field ordering** | Largest-first reduces padding. |

## Hint

Order `B int64` first, then the two bools: `B; A; C`.

## Validate

```bash
make verify
```
