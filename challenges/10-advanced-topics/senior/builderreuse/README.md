# One Builder, Many Lines

**Level:** senior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A CSV writer constructs a `strings.Builder` per row. Each one starts empty, grows through several sizes, and is discarded a line later.

## Task

Implement [builderreuse.go](builderreuse.go):

1. Render each row as its values joined by `-`.
2. Use one builder for the whole call, reset between rows.
3. Reserve its capacity once, before the loop.
4. An empty row renders as the empty string.

Replace the stub body in [builderreuse.go](builderreuse.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  RenderLines([][]int{{1,2},{3}})
Output: ["1-2" "3"]
```

**Example 2:**

```
Input:  RenderLines([][]int{{}})
Output: [""]
```

**Example 3:**

```
Input:  RenderLines([][]int{{-1,2,-3}})
Output: ["-1-2--3"]
```

_Explanation:_ The separator and the minus sign both appear.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Builder.Reset** | Empties the builder while keeping the buffer it has grown. |
| 2 | **Grow once** | Reserving the widest row's worth removes every intermediate growth. |
| 3 | **String extraction copies** | `b.String()` is the per-row allocation you cannot avoid. |

## Hint

Construct and `Grow` above the loop; `Reset` inside it.

## Validate

```bash
make verify
```
