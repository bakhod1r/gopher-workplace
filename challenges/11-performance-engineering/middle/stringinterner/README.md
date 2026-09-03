# One Copy Of Every Label

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

Parse a million log lines and you get a million copies of the twenty distinct level names. Interning collapses them: keep a map from value to canonical instance and hand out that instance forever after. Metric labels, HTTP header names and column names in a query engine are all interned for exactly this reason.

## Task

Implement the four methods in [stringinterner.go](stringinterner.go):

1. `Intern` returns the canonical instance — first one in wins, later equal strings resolve to it and share its backing array.
2. `InternBytes` interns the string form of a `[]byte` without retaining the caller's buffer, and must not allocate when the string is already known.
3. `Stats` reports hits and misses; `Len` reports the number of distinct strings.

## Examples

**Example 1:**

```
Input:  Intern(x) twice
Output: the same instance both times, sharing one backing array
```

**Example 2:**

```
Input:  buf := []byte("label"); InternBytes(buf); overwrite buf
Output: the interned string still reads "label"
```

**Example 3:**

```
Input:  Intern("a") three times and Intern("b") once
Output: Stats 2 hits, 2 misses; Len 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Interning trades a map for memory** | One hash lookup per value, one copy per distinct value. |
| 2 | **`m[string(b)]` does not allocate** | The compiler recognises the pattern and skips the conversion for the lookup. |
| 3 | **The map key must be a copy** | Storing a string built from a buffer the caller reuses is a use-after-write bug. |

## Topics used again

Maps, lazy initialisation, `unsafe.StringData` semantics, the comma-ok idiom.

## Hint

Look up with `in.seen[string(b)]` first; only on a miss do you need `string(b)` as a real, allocated key.

## Validate

```bash
make verify
```
