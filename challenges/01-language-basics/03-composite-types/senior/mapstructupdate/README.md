# Update Struct in Map

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`m[key]` returns a **copy** of the struct. Incrementing `s.Hits` changes the copy;
without writing it back, the map keeps the old value. (Go even forbids
`m[key].Hits++` directly.)

## Task

Fix the line between the markers in
[mapstructupdate.go](mapstructupdate.go) to store the updated struct.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Record(m,'a') once
Output: m['a'].Hits==1
```

**Example 2:**

```
Input:  Record(m,'a') twice
Output: m['a'].Hits==2
```

**Example 3:**

```
Input:  Record(m,'b') on empty m
Output: m['b'].Hits==1
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Map value copy** | `m[key]` is not addressable. |
| 2 | **Read-modify-write** | Copy, mutate, assign back. |
| 3 | **No `m[k].F++`** | Compile error; that's why. |

## Hint

`m[key] = s`.

## Validate

```bash
make verify
```
