# Map Comma-Ok

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

A feature-flag store maps names to counts. A plain `m[key]` returns `0` both
when a key is missing *and* when it is present with value `0` — the code can't
tell them apart. Go's comma-ok read solves this.

## Task

Implement `Lookup` in [lookup.go](lookup.go) so it returns the value for `key`
and a boolean reporting whether the key was actually present. A key mapped to
`0` must report `true`.

Do **not** change the function signature or the tests.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  m={"a":5}, key="a"
Output: 5, true
```

**Example 2:**

```
Input:  m={"a":5}, key="z"
Output: 0, false
```

**Example 3:**

```
Input:  m={"z":0}, key="z"
Output: 0, true
```

_Explanation:_ Present with zero value; comma-ok reports found.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Comma-ok read** | `v, ok := m[key]` returns the value and whether the key exists. |
| 2 | **Zero value ambiguity** | Plain `m[key]` yields the value type's zero for a missing key — indistinguishable from a stored zero. |
| 3 | **Nil map reads** | Reading a `nil` map is safe and always reports "not present". |

## Hint

Use the two-result form of the map index: `v, ok := m[key]` and return both.
Don't test `v != 0` to decide presence — that misreports a stored `0`.

## Validate

```bash
make verify
```
