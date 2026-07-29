# LIFO via Deferred Push

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _defer_

## Context

The deferred closure always appends `xs[len(xs)-1]` — the same last element —
regardless of `i`. To reverse, each deferred push must record its OWN element,
captured per iteration (as an argument or via the per-iteration `i`).

## Task

Fix [stackpush.go](stackpush.go) so it appends each element once, in reverse.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ReverseInts([1 2 3])
Output: [3 2 1]
```

**Example 2:**

```
Input:  ReverseInts([1])
Output: [1]
```

**Example 3:**

```
Input:  ReverseInts(nil)
Output: []
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Per-iteration value** | Each defer must push its own element. |
| 2 | **LIFO ordering** | Deferred pushes run last-to-first. |
| 3 | **Index vs constant** | `xs[i]`, not `xs[len-1]`. |

## Hint

Push the current element: `defer func(v int){ out = append(out, v) }(xs[i])` (or use `xs[i]` with per-iteration `i`).

## Validate

```bash
make verify
```
