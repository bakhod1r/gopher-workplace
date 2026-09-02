# Writer Interface

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A report generator writes bytes into whichever sink is configured, and reports how much it wrote.

## Task

Implement the stub(s) in [writerifc.go](writerifc.go):

1. Implement `Write` on `*Builder` — append the text and return how many bytes were written.
2. Implement `WriteLines`, which writes each line followed by `"\n"` and returns the total byte count.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  b := &Builder{}; b.Write("ab")
Output: 2
```

**Example 2:**

```
Input:  b.String()
Output: "ab"
```

**Example 3:**

```
Input:  WriteLines(&Builder{}, []string{"a", "b"})
Output: 4
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Writer-shaped interface** | Return the byte count, like `io.Writer` does. |
| 2 | **Accumulating a total** | Reused: sum the per-call results. |
| 3 | **String building** | Reused: `+=` on a string field. |

## Hint

Each line contributes `len(line) + 1` bytes because of the newline.

## Validate

```bash
make verify
```
