# Append to Nil Is Valid

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

`append` to a nil slice allocates and works exactly like appending to an empty
one. Guarding against nil with an early return discards data unnecessarily.

## Task

Fix [nilappend.go](nilappend.go) so a nil starting slice still collects the extras.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Collect(nil, [1 2 3])
Output: [1 2 3]
```

**Example 2:**

```
Input:  Collect([0], [9])
Output: [0 9]
```

**Example 3:**

```
Input:  Collect(nil, nil)
Output: []
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **nil slice append** | `append(nil, v)` is valid and allocates. |
| 2 | **nil vs empty** | A nil slice behaves as empty for len/range/append. |
| 3 | **Unnecessary guard** | The nil check drops valid work. |

## Hint

Remove the `if xs == nil { return nil }` guard; append works on nil.

## Validate

```bash
make verify
```
