# Flyweight Pattern

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

Font data is expensive and identical for every user of the same font name. A
flyweight factory keeps one shared instance per key in a map and hands out the
same pointer to everyone who asks.

## Task

Implement `Get` on `*FlyweightFactory` in [flyweightpatt.go](flyweightpatt.go):

1. Look `name` up in `f.fonts`.
2. If it is missing, create a `&FontData{}` and store it under that key.
3. Return the stored pointer.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Get("Arial") on an empty factory
Output: a new *FontData, now cached
```

**Example 2:**

```
Input:  Get("Arial") twice
Output: the same pointer both times
```

**Example 3:**

```
Input:  Get("Arial"), Get("Times")
Output: two different pointers
```

_Explanation:_ sharing is per key, not global.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Comma-ok map lookup** | `v, ok := m[k]` distinguishes "absent" from "present but zero". |
| 2 | **Pointer identity** | The test compares pointers with `==` — returning a fresh value each call fails it. |
| 3 | **Maps hold references** | Storing `*FontData` means every caller shares one allocation. |

## Hint

`if v, ok := f.fonts[name]; ok { return v }`, otherwise create, store, return.
Returning a new `&FontData{}` without storing it breaks the identity check.

## Validate

```bash
make verify
```
