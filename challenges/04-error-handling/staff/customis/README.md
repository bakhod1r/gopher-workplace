# Custom Is Method

**Level:** staff
**Topic:** 04-error-handling

## Context

Status errors must match by class as well as by exact code, so a handler can ask for "any 5xx" without listing every code.

## Task

Implement `StatusError` in [customis.go](customis.go):

1. Give `*StatusError` an `Error() string` of the form `"status <Code>"`.
2. Give it an `Is(target error) bool` matching another `*StatusError` with the same code, or with a `Code` that is the class marker (500 matches any 5xx).
3. Leave non-`*StatusError` targets unmatched.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  errors.Is(&StatusError{503}, &StatusError{503})
Output: true
```

**Example 2:**

```
Input:  errors.Is(&StatusError{503}, &StatusError{500})
Output: true
```

**Example 3:**

```
Input:  errors.Is(&StatusError{404}, &StatusError{500})
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **The Is method** | `errors.Is` calls it before comparing values. |
| 2 | **Asymmetry** | The receiver decides what it matches; the target does not. |
| 3 | **Class matching** | Integer division by 100 gives the status class. |

## Hint

`errors.Is(err, target)` calls `err.Is(target)` — the receiver is the error being tested, the argument is what the caller asked for.

## Validate

```bash
make verify
```
