# Group By

**Level:** middle
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A reporting job buckets records by a key the caller chooses at run time.

## Task

Implement the stub(s) in [groupby.go](groupby.go):

1. Implement `Key` on `ByFirstLetter` and `ByLength`.
2. Implement `Group`, which buckets the input by the key function, preserving input order inside each bucket.
3. Implement `SortedKeys`, which returns a grouping's keys in sorted order.
4. Make one pass over the input — do not scan it once per key.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Group([]string{"apple", "avocado", "beet"}, ByFirstLetter{})
Output: {"a": [apple avocado], "b": [beet]}
```

**Example 2:**

```
Input:  Group([]string{"ab", "cd"}, ByLength{})
Output: {"2": [ab cd]}
```

**Example 3:**

```
Input:  SortedKeys of the first grouping
Output: ["a", "b"]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Key extractor interface** | The grouping policy is a value, so one function serves every report. |
| 2 | **map of slices** | Reused from composite types: appending to a map value works on a nil slice. |
| 3 | **Single pass** | Efficiency: O(n) instead of O(n·k). |

## Hint

`m[k] = append(m[k], v)` works even when `m[k]` does not exist yet.

## Validate

```bash
make verify
```
