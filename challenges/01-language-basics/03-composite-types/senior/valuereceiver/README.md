# Receiver Doesn't Mutate

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`Credit` has a **value** receiver, so `w.Balance += amount` updates a copy that's
discarded — the wallet never changes.

## Task

Fix the receiver between the markers in
[valuereceiver.go](valuereceiver.go) to a pointer.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  w={0}; Credit(100)
Output: w.Balance==100
```

**Example 2:**

```
Input:  w={50}; Credit(25)
Output: w.Balance==75
```

**Example 3:**

```
Input:  w={10}; Credit(0)
Output: w.Balance==10
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

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
