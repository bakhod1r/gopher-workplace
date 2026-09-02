# Slice Retention Leak

**Level:** staff
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A parser kept only the first 8 bytes of each 1MB message, and memory still grew without bound. The small slice was pinning the whole backing array.

## Task

Implement the stub(s) in [sliceleak.go](sliceleak.go):

1. Implement `Prefix`, returning a sub-slice that still aliases the source (the leak).
2. Implement `PrefixCopy`, returning an independent copy that releases the source.
3. Implement `RetainedBytes`, reporting how much memory a result still pins (its capacity).
4. Constraint: `PrefixCopy` must retain only what it keeps; the test compares the retained capacities.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Prefix(1MB buffer, 8)
Output: length 8, capacity ~1MB
```

**Example 2:**

```
Input:  PrefixCopy(the same buffer, 8)
Output: length 8, capacity 8
```

**Example 3:**

```
Input:  RetainedBytes on each
Output: the copy retains far less
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Slice retention** | A sub-slice keeps the entire backing array reachable, whatever its length. |
| 2 | **cap versus len** | Capacity is what the GC sees; length is only what you can index. |
| 3 | **Copy to release** | The standard fix: copy the bytes you keep and drop the reference. |

## Hint

`b[:n]` shares the array. `append([]byte(nil), b[:n]...)` allocates exactly what it keeps.

## Validate

```bash
make verify
```
