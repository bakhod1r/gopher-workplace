# Let The Standard Library Size It

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

`strings.Join` measures the result, allocates once, and copies each segment in. Hand-rolling the same thing with `s += part + "/"` allocates once per segment and copies the growing prefix every time — quadratic work for an operation the standard library already does in linear time.

## Task

Implement both functions in [joinvsloop.go](joinvsloop.go):

1. `JoinPath` joins the segments with `"/"` in a single allocation; no segments yields `""`.
2. `SplitPath` splits on `"/"`.
3. `SplitPath("")` returns an empty, non-nil slice, so the two are inverses.

## Examples

**Example 1:**

```
Input:  JoinPath([a b])
Output: "a/b"
```

**Example 2:**

```
Input:  JoinPath([a "" b])
Output: "a//b"
```

**Example 3:**

```
Input:  SplitPath("")
Output: [] (non-nil)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`strings.Join` is one allocation** | It sums the lengths first, exactly like a hand-written `Grow`. |
| 2 | **`+=` in a loop is quadratic** | Each concatenation copies everything accumulated so far. |
| 3 | **`strings.Split("")` returns `[""]`** | One empty element, not zero — the inverse property needs that special case. |

## Topics used again

`strings.Join`, `strings.Split`, slices.

## Hint

`SplitPath` needs one guard that `JoinPath` does not.

## Validate

```bash
make verify
```
