# Map Get With Default

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

Reading a missing map key returns the zero value — indistinguishable from a real
zero. The comma-ok idiom tells them apart.

## Task

Implement `GetOr(m, key, def)` returning the value if present, else `def`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  m={"a":1,"zero":0}, key="a", def=99
Output: 1
```

**Example 2:**

```
Input:  key="zero", def=99
Output: 0
```

_Explanation:_ Present with value 0 — the comma-ok distinguishes it from missing.

**Example 3:**

```
Input:  key="missing", def=99
Output: 99
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Comma-ok** | `v, ok := m[key]` reports presence. |
| 2 | **Zero vs absent** | A stored 0 differs from a missing key. |
| 3 | **Nil map read** | Reading nil is safe, returns absent. |

## Hint

`if v, ok := m[key]; ok { return v }; return def`.

## Validate

```bash
make verify
```
