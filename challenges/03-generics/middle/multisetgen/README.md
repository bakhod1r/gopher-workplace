# Multiset

**Level:** middle  
**Topic:** 03-generics

## Context

An inventory tracks how many of each item are on hand, and removing the last one should leave no trace of the item.

## Task

Implement the stub(s) in [multisetgen.go](multisetgen.go):

1. Implement `NewBag`, `Add`, `Remove`, `Count`, and `Distinct`.
2. `Remove` reports `false` when there is nothing to remove.
3. Removing the last occurrence must delete the key, not leave a zero count.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Add(a); Add(a); Count(a)
Output: 2
```

**Example 2:**

```
Input:  Add(a); Remove(a); Distinct()
Output: 0
```

**Example 3:**

```
Input:  Remove(missing)
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Zero counts are not entries** | Deleting at zero keeps `Distinct` honest. |
| 2 | **`delete` on maps** | Reused from language basics: removing a key is how you shrink a map. |
| 3 | **Structural invariants** | A data structure is defined by what stays true after every operation. |

## Hint

Delete the key when the count reaches zero — otherwise `Distinct` over-reports.

## Validate

```bash
make verify
```
