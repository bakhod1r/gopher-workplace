# Ring Buffer Wrap

**Level:** staff
**Topic:** 01-language-basics → 03-composite-types

## Context

A ring buffer maps logical index `i` to physical `(head+i) mod len`. The code
omits the modulo, so once `head+i` reaches the end it indexes out of range (and
never wraps).

## Task

Fix the index between the markers in [ringbuffer.go](ringbuffer.go) to wrap.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  buf=[10 20 30], head=1, i=0
Output: 20
```

**Example 2:**

```
Input:  buf=[10 20 30], head=1, i=2
Output: 10
```

_Explanation:_ (1+2)%3=0 wraps around.

**Example 3:**

```
Input:  buf=[10 20 30], head=0, i=4
Output: 20
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Modular index** | `(head+i) % len(buf)`. |
| 2 | **Wraparound** | Logical past-end maps to the front. |
| 3 | **Fixed storage** | Physical size is constant. |

## Hint

`return buf[(head+i)%len(buf)]`.

## Validate

```bash
make verify
```
