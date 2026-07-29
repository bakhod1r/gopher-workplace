# Append to Nil

**Level:** staff
**Topic:** 01-language-basics → 03-composite-types

## Context

A defensive nil-check returns early, but appending to a **nil** slice is perfectly
valid — it allocates and returns a one-element slice. The guard drops the first
element.

## Task

Remove the unnecessary guard between the markers in
[appendnilreturn.go](appendnilreturn.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  s=nil, x=5
Output: [5]
```

_Explanation:_ append to a nil slice allocates a new one-element slice.

**Example 2:**

```
Input:  s=[1 2], x=3
Output: [1 2 3]
```

**Example 3:**

```
Input:  s=[], x=7
Output: [7]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Nil is appendable** | `append(nil, x)` works. |
| 2 | **Zero value slice** | nil behaves like empty for append/range/len. |
| 3 | **Over-defensive code** | The check is harmful here. |

## Hint

Delete the `if s == nil { return nil }` block.

## Validate

```bash
make verify
```
