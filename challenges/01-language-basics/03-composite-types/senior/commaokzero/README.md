# Zero vs Missing

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`GetOr` checks `v != 0` to decide presence, so a key stored with value `0`
wrongly returns the default. Only the comma-ok form distinguishes them.

## Task

Fix the presence check between the markers in
[commaokzero.go](commaokzero.go) to use comma-ok.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  m={a:0}, key=a, def=7
Output: 0
```

_Explanation:_ Key present with value 0 -> return stored 0, not default.

**Example 2:**

```
Input:  m={a:5}, key=b, def=7
Output: 7
```

**Example 3:**

```
Input:  m={}, key=x, def=-1
Output: -1
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Comma-ok** | `v, ok := m[key]`. |
| 2 | **Zero ≠ absent** | A stored 0 is present. |
| 3 | **Presence flag** | Branch on `ok`, not the value. |

## Hint

`if v, ok := m[key]; ok { return v }`.

## Validate

```bash
make verify
```
