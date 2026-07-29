# Address of a Hoisted Loop Var

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _with-maps-and-slices_

## Context

`v` is declared ONCE outside the loop, so `&v` is the same address every
iteration; all pointers see the final value. Take the address of a fresh
per-iteration variable instead.

## Task

Fix [hoistedaddr.go](hoistedaddr.go) so each pointer holds a distinct value.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Pointers([10 20 30])
Output: pointers to 10, 20, 30
```

**Example 2:**

```
Input:  deref each result
Output: 10, 20, 30 (distinct)
```

**Example 3:**

```
Input:  Pointers([])
Output: empty
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Hoisted variable address** | One `v`, one address, shared. |
| 2 | **Per-iteration copy** | Declare a fresh variable inside the loop. |
| 3 | **Distinct pointers** | Each must own its value. |

## Hint

Take the address of a fresh variable per iteration: `v := xs[i]; out = append(out, &v)` (declared inside the loop).

## Validate

```bash
make verify
```
