# IPv4 Octet Range

**Level:** senior
**Topic:** 01-language-basics → 02-data-types

## Context

A firewall config loader parses `192.168.1.1` into four bytes. It accepts
`256.1.1.1` because the range check is wrong — and `byte(256)` silently wraps to
0, corrupting the rule.

## Task

Fix the single check between the markers in [ipv4parse.go](ipv4parse.go) so each
octet must be 0..255.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  "192.168.1.1"
Output: [192 168 1 1], true
```

**Example 2:**

```
Input:  "256.1.1.1"
Output: zero, false
```

_Explanation:_ octet > 255

**Example 3:**

```
Input:  "1.2.3"
Output: zero, false
```

_Explanation:_ only 3 fields

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Octet range** | Each field is a byte: 0..255. |
| 2 | **Conversion wrap** | `byte(256)` becomes 0. |
| 3 | **Validate before convert** | Reject >255 before `byte(val)`. |

## Hint

`if val > 255`.

## Validate

```bash
make verify
```
