# Clone A Map Without Reflection

**Level:** staff
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A codebase has four hand-written map cloners, one per key-value pair it needed, and a fifth written with reflection that boxes every key and value.

## Task

Implement [typedclone.go](typedclone.go):

1. Return a shallow copy of `m` with the same entries.
2. A nil map clones to nil; an empty map clones to an empty, non-nil map.
3. Size the new map up front; nothing may be boxed.

Replace the stub body in [typedclone.go](typedclone.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  CloneMap(map[string]int{"a":1})
Output: a new map with the same entry
```

**Example 2:**

```
Input:  CloneMap(nil)
Output: <nil>
```

_Explanation:_ nil is preserved, not turned into an empty map.

**Example 3:**

```
Input:  got["a"][0] = 99 on a map of slices
Output: the original sees it
```

_Explanation:_ The clone is shallow.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Type parameters over reflection** | The compiler emits code for the concrete types; nothing is boxed. |
| 2 | **comparable as the key constraint** | Exactly the types Go allows as map keys. |
| 3 | **nil versus empty** | They behave differently and a clone must preserve which one it had. |
| 4 | **Shallow by contract** | Reference values inside are shared, and that must be documented. |

## Hint

Preserve nil, size the result, and copy.

## Validate

```bash
make verify
```
