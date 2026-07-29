# Field Order and Struct Size

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _unsafe-pointer_

## Context

Each field aligns to its size; the compiler inserts padding. `bool,int64,bool`
pads to 24; ordering widest-first `int64,bool,bool` packs to 16.

## Task

Fix the field order in [sizeorder.go](sizeorder.go) so `Record` is 16 bytes.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  unsafe.Sizeof(Record{})
Output: 16
```

_Explanation:_ Fields ordered wide-to-narrow pack without padding waste.

**Example 2:**

```
Input:  fields A bool, B int64, C bool
Output: 24 (padded)
```

**Example 3:**

```
Input:  fields B int64, A bool, C bool
Output: 16 (packed)
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Alignment and padding** | Fields align to their size. |
| 2 | **Widest-first ordering** | Put int64 before the bools. |
| 3 | **unsafe.Sizeof** | Reveals the true size. |

## Hint

Order widest-to-narrowest: `B int64; A bool; C bool`.

## Validate

```bash
make verify
```
