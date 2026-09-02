# Embedding comparable

**Level:** middle  
**Topic:** 03-generics

## Context

An audit tool tallies IDs that are always ints or strings. Accepting arbitrary comparable types would let structs in, which the storage layer cannot serialise.

## Task

Implement the stub(s) in [comparableembed.go](comparableembed.go):

1. Implement `Tally`, returning the count per distinct key.
2. Study `Key`: it embeds `comparable` and then narrows to a type set.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Tally([]int{1,1,2})
Output: {1:2, 2:1}
```

**Example 2:**

```
Input:  Tally([]string{"a"})
Output: {"a":1}
```

**Example 3:**

```
Input:  Tally([]int{})
Output: {}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Embedding `comparable`** | `comparable` may be embedded in a constraint interface alongside a type set. |
| 2 | **Narrowing on purpose** | The set excludes structs and pointers even though they are comparable. |
| 3 | **Redundant but explicit** | Every member of the set is already comparable — the embed documents the requirement. |

## Hint

The constraint is `comparable` plus a type set — both restrictions apply at once.

## Validate

```bash
make verify
```
