# Defer Snapshots Slice Header

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _defer_

## Context

Passing `xs` as a deferred argument copies its slice HEADER (ptr,len,cap) at
defer-time, when len is 0. Later appends reassign `xs` to a new header the
snapshot never sees. Read `xs` in the closure body to get the final length.

## Task

Fix [deferslicearg.go](deferslicearg.go) so the reported length is final.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  BuildAndReport(4)
Output: 4
```

**Example 2:**

```
Input:  BuildAndReport(0)
Output: 0
```

**Example 3:**

```
Input:  reported equals final length
Output: true
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Slice header snapshot** | A deferred arg copies len/cap now. |
| 2 | **Reassignment invisibility** | `xs = append(...)` makes a new header. |
| 3 | **Body capture** | Read `xs` at return-time in the closure. |

## Hint

Capture in the body instead of as an argument: `defer func(){ reported = len(xs) }()`.

## Validate

```bash
make verify
```
