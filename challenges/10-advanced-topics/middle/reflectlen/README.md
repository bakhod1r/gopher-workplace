# Length Of Whatever You Are Handed

**Level:** middle
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A validation layer checks that a request field is not empty. The field can be a string, a slice or a map depending on the endpoint, and each one got its own copy of the same check.

## Task

Implement [reflectlen.go](reflectlen.go):

1. Return `v`'s length when it has one, and whether it did.
2. Strings, slices, arrays, maps and channels have lengths; nothing else does.
3. Never panic — a nil interface reports 0, false.

Replace the stub body in [reflectlen.go](reflectlen.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Length([]int{1,2})
Output: 2, true
```

**Example 2:**

```
Input:  Length("日本")
Output: 6, true
```

_Explanation:_ A string's length is in bytes.

**Example 3:**

```
Input:  Length(3)
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Value.Len** | Defined only for the five kinds that have a length; it panics on the rest. |
| 2 | **Kind switching as a guard** | The switch is both the dispatch and the safety check. |
| 3 | **Nil containers have length 0** | A nil slice or map is still a valid Value of that kind. |

## Hint

A kind switch with five cases in it and a default that says no.

## Validate

```bash
make verify
```
