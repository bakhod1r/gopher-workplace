# Reject A Bad Schema Before It Ships

**Level:** senior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A tag typo maps two struct fields to one database column. The mistake is invisible in review and shows up as data loss in production.

## Task

Implement [tagvalidate.go](tagvalidate.go):

1. Report every problem with `v`'s `col` tags, in field order.
2. A missing or empty tag, a comma in the tag, and a duplicate tag are each a problem.
3. Skip unexported fields; report `"not a struct"` for anything that is not a struct.
4. One problem per field: report the first one found and move on.

Replace the stub body in [tagvalidate.go](tagvalidate.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Validate(good{})
Output: []
```

_Explanation:_ No problems.

**Example 2:**

```
Input:  Validate(missing{})
Output: [B: missing col tag]
```

**Example 3:**

```
Input:  Validate(dup{})
Output: [B: duplicate tag of A]
```

_Explanation:_ The later field is the one reported.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Schema validation at run time** | The struct is the schema, so it can check itself. |
| 2 | **Tag.Lookup** | Distinguishes an absent tag from an empty one — both are problems here. |
| 3 | **Deterministic reporting** | Field order makes the output stable and diffable. |
| 4 | **First problem per field** | Continuing after a report keeps each field to one line. |

## Hint

One pass, a `seen` map from tag to the field that claimed it.

## Validate

```bash
make verify
```
