# Zip With

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

Zipping walks two slices in lockstep, stopping at the shorter one, applying a
binary function to each pair.

## Task

Implement `ZipWith` in [zipwith.go](zipwith.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ZipWith([1 2 3], [10 20 30], add)
Output: [11 22 33]
```

**Example 2:**

```
Input:  ZipWith([1 2], [10], add)
Output: [11]
```

**Example 3:**

```
Input:  ZipWith(nil, [1], add)
Output: []
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Lockstep iteration** | Index up to `min(len(a), len(b))`. |
| 2 | **Binary combiner** | `f(a[i], b[i])`. |
| 3 | **Shorter wins** | Extra tail is dropped. |

## Hint

Compute `n := min(len(a), len(b))`; loop `i < n` appending `f(a[i], b[i])`.

## Validate

```bash
make verify
```
