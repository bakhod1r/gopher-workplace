# Do These Two Strings Share Their Bytes

**Level:** middle
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A cache is supposed to hand out the interned copy of a key. Proving it does — rather than returning an equal but separate string — needs identity, which `==` cannot express.

## Task

Implement [samestring.go](samestring.go):

1. Report whether `a` and `b` have the same length and the same start address.
2. Two equal but separately allocated strings report false.
3. Empty strings report false — there is no storage to share.

Replace the stub body in [samestring.go](samestring.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  s := "abc"; SameBytes(s, s)
Output: true
```

**Example 2:**

```
Input:  SameBytes(s[1:], s[:7])
Output: false
```

_Explanation:_ Different starts.

**Example 3:**

```
Input:  a and string([]byte(a))
Output: false
```

_Explanation:_ Equal, not identical.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **unsafe.StringData** | The address of a string's first byte. |
| 2 | **Equality versus identity** | `==` compares bytes; this compares storage. |
| 3 | **Interning is about identity** | The whole point is that repeats share one allocation. |

## Hint

Length first, then the data pointers.

## Validate

```bash
make verify
```
