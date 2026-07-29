# Caesar Cipher

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

Each letter shifts by `n` within its case, wrapping z→a. Non-letters stay put.
Rune arithmetic plus a modulo does the wrap.

## Task

Implement `Shift(s, n)` (letters only, case preserved, wrap, `n` any int).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Shift("abc", 1)
Output: "bcd"
```

_Explanation:_ each letter +1

**Example 2:**

```
Input:  Shift("xyz", 3)
Output: "abc"
```

_Explanation:_ wraps past z back to a

**Example 3:**

```
Input:  Shift("Hello, World!", 13)
Output: "Uryyb, Jbeyq!"
```

_Explanation:_ ROT13; punctuation/space unchanged

**Example 4:**

```
Input:  Shift("abc", -1)
Output: "zab"
```

_Explanation:_ negative shift normalized modulo 26

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Rune ranges** | `'a'..'z'`, `'A'..'Z'` are contiguous. |
| 2 | **Modular wrap** | `((r-base+n)%26+26)%26` keeps 0..25. |
| 3 | **Building strings** | Accumulate in a `[]rune` or strings.Builder. |

## Hint

For a letter, base = `'a'` or `'A'`; new = `base + ((r-base+n)%26+26)%26`.

## Validate

```bash
make verify
```
