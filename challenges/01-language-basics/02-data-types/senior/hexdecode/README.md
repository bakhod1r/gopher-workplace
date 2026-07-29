# Hex Nibble Order

**Level:** senior
**Topic:** 01-language-basics → 02-data-types

## Context

A wire-protocol decoder turns hex into bytes, but every byte comes out with its
two nibbles swapped: `"1a"` decodes to `0xa1`. The high nibble is the *first*
character.

## Task

Fix the single line between the markers in [hexdecode.go](hexdecode.go) so the
first hex char is the high nibble.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  "ff"
Output: [0xff], true
```

**Example 2:**

```
Input:  "476F"
Output: "Go", true
```

**Example 3:**

```
Input:  "abc"
Output: nil, false
```

_Explanation:_ odd length

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Nibble position** | First char = high 4 bits. |
| 2 | **Shift + OR** | `hi<<4 |
| 3 | **Even length** | Two hex chars per byte. |

## Hint

`byte(hi<<4 | lo)`.

## Validate

```bash
make verify
```
