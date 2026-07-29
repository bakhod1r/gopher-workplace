# First, Last, OK

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _multiple-return_

## Context

Returning an `ok bool` alongside the data lets the caller distinguish a real
result from the empty case without a sentinel value or panic.

## Task

Implement `FirstLast` in [firstlast.go](firstlast.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  FirstLast([1 2 3])
Output: 1, 3, true
```

**Example 2:**

```
Input:  FirstLast([9])
Output: 9, 9, true
```

**Example 3:**

```
Input:  FirstLast(nil)
Output: 0, 0, false
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Comma-ok style return** | An extra `bool` signals validity. |
| 2 | **Empty guard** | `len(xs) == 0` returns the zero result early. |
| 3 | **Index math** | Last element is `xs[len(xs)-1]`. |

## Hint

Guard `len(xs)==0` first, else return `xs[0], xs[len(xs)-1], true`.

## Validate

```bash
make verify
```
