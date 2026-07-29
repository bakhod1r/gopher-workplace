# Recovered Error Lost to Local

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _defer_

## Context

Inside the deferred recover, assigning to a fresh local `e` does nothing for the
caller; you must assign the NAMED return `err` so it propagates out.

## Task

Fix [recoverlost.go](recoverlost.go) so the recovered panic becomes the returned error.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Safe(func(){ panic("boom") })
Output: non-nil error
```

**Example 2:**

```
Input:  Safe(func(){})
Output: nil
```

**Example 3:**

```
Input:  error message contains boom
Output: true
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Named return from defer** | Only assigning `err` reaches the caller. |
| 2 | **Local shadow** | A new `e` is discarded. |
| 3 | **panic -> error** | Convert `r` with `fmt.Errorf`. |

## Hint

Assign the named return: `err = fmt.Errorf("panic: %v", r)`.

## Validate

```bash
make verify
```
