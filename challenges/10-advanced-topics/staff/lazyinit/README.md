# Build It Once, However Many Ask

**Level:** staff
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A lookup table is built on first use behind a "if index == nil" check. Under a cold-start burst several goroutines see nil at once, three indexes are built, and two of them are thrown away while readers are using them.

## Task

Implement [lazyinit.go](lazyinit.go):

1. Return the index of `k` in the table's pairs, building the index on first use.
2. The index must be built exactly once per table, even under concurrent first use.
3. Every caller must observe a fully built index.

Replace the stub body in [lazyinit.go](lazyinit.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  NewTable([][2]string{{"a","1"}}).Lookup("a")
Output: 0, true
```

**Example 2:**

```
Input:  500 lookups
Output: one build
```

**Example 3:**

```
Input:  32 concurrent first lookups
Output: one build, all succeed
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **sync.Once** | Runs the function once and blocks the others until it has finished. |
| 2 | **The happens-before it provides** | Every `Do` returns after the initialisation's writes are visible. |
| 3 | **Check-then-act is not atomic** | A nil check plus an assignment is a race with a benign-looking symptom. |
| 4 | **Per-instance state** | The `Once` lives in the table, so two tables build independently. |

## Hint

One line before the map read. The build function is already written.

## Validate

```bash
make verify
```
