# iota Sequence Break

**Level:** staff
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

`iota` still increments per ConstSpec line even when you assign an explicit
value — but hardcoding `Redirect = 7` desyncs the *named* value from the counter,
and the following bare lines resume from iota (3, 4), leaving a gap. The fix is
to let `Redirect` follow the run.

## Task

Fix the single line between the markers in [codes.go](codes.go) so the classes
are 0,1,2,3,4 in order.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Info, Success
Output: 0, 1
```

**Example 2:**

```
Input:  ClientError
Output: 3
```

**Example 3:**

```
Input:  ServerError
Output: 4
```

_Explanation:_ The explicit 7 broke the sequence.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **iota per line** | Increments even across explicit assignments. |
| 2 | **Implicit repetition** | A bare name repeats the previous expression. |
| 3 | **Sequence integrity** | An explicit value desyncs the enum. |

## Hint

Make `Redirect` a bare line (drop `= 7`) so it inherits `iota`.

## Validate

```bash
make verify
```
